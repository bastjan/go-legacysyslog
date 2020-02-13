// Original work Copyright (c) 2018, InfluxData Inc.

package legacysyslog

import (
	"time"

	"github.com/influxdata/go-syslog/v3"
)

type CiscoTimeClockMode int8

const (
	CiscoTimeClockModeUnset CiscoTimeClockMode = iota
	CiscoTimeClockModeSynced
	CiscoTimeClockModeUnsynced
)

type syslogMessage struct {
	prioritySet        bool
	ciscoSequenceIDSet bool
	timestampSet       bool
	hostnameSet        bool
	tagSet             bool
	contentSet         bool
	messageSet         bool

	ciscoTimestampExt CiscoTimeClockMode
	priority          uint8

	ciscoSequenceID string
	timestamp       time.Time
	hostname        string
	tag             string
	content         string
	message         string
}

type SyslogMessage struct {
	priority *uint8
	facility *uint8
	severity *uint8

	ciscoSequenceID   *string
	ciscoTimestampExt CiscoTimeClockMode
	timestamp         *time.Time
	hostname          *string
	tag               *string
	content           *string
	message           *string
}

var _ syslog.Message = &SyslogMessage{}

func (sm *syslogMessage) export() *SyslogMessage {
	out := &SyslogMessage{
		ciscoTimestampExt: sm.ciscoTimestampExt,
	}

	if sm.prioritySet {
		out.setPriority(sm.priority)
	}
	if sm.ciscoSequenceIDSet {
		out.ciscoSequenceID = &sm.ciscoSequenceID
	}

	if sm.timestampSet {
		out.timestamp = &sm.timestamp
	}
	if sm.hostnameSet {
		out.hostname = &sm.hostname
	}
	if sm.tagSet {
		out.tag = &sm.tag
	}
	if sm.contentSet {
		out.content = &sm.content
	}
	if sm.messageSet {
		out.message = &sm.message
	}

	return out
}

// Valid tells whether the receiving SyslogMessage is well-formed or not.
// TODO: only valid if priority or default priority
func (sm *SyslogMessage) Valid() bool {
	return true
}

// Priority returns the syslog priority or nil when not set
func (sm *SyslogMessage) Priority() *uint8 {
	return sm.priority
}

// Version returns the syslog version or nil when not set
func (sm *SyslogMessage) Version() uint16 {
	return 0
}

func (sm *SyslogMessage) setPriority(value uint8) {
	sm.priority = &value
	facility := uint8(value / 8)
	severity := uint8(value % 8)
	sm.facility = &facility
	sm.severity = &severity
}

// Facility returns the facility code.
func (sm *SyslogMessage) Facility() *uint8 {
	return sm.facility
}

// Severity returns the severity code.
func (sm *SyslogMessage) Severity() *uint8 {
	return sm.severity
}

// FacilityMessage returns the text message for the current facility value.
func (sm *SyslogMessage) FacilityMessage() *string {
	if sm.facility != nil {
		msg := facilities[*sm.facility]
		return &msg
	}

	return nil
}

// FacilityLevel returns the
func (sm *SyslogMessage) FacilityLevel() *string {
	if sm.facility != nil {
		if msg, ok := facilityKeywords[*sm.facility]; ok {
			return &msg
		}

		// Fallback to facility message
		msg := facilities[*sm.facility]
		return &msg
	}

	return nil
}

// SeverityMessage returns the text message for the current severity value.
func (sm *SyslogMessage) SeverityMessage() *string {
	if sm.severity != nil {
		msg := severityMessages[*sm.severity]
		return &msg
	}

	return nil
}

// SeverityLevel returns the text level for the current severity value.
func (sm *SyslogMessage) SeverityLevel() *string {
	if sm.severity != nil {
		msg := severityLevels[*sm.severity]
		return &msg
	}

	return nil
}

// SeverityShortLevel returns the short text level for the current severity value.
func (sm *SyslogMessage) SeverityShortLevel() *string {
	if sm.severity != nil {
		msg := severityLevelsShort[*sm.severity]
		return &msg
	}

	return nil
}

var severityMessages = map[uint8]string{
	0: "system is unusable",
	1: "action must be taken immediately",
	2: "critical conditions",
	3: "error conditions",
	4: "warning conditions",
	5: "normal but significant condition",
	6: "informational messages",
	7: "debug-level messages",
}

var severityLevels = map[uint8]string{
	0: "emergency",
	1: "alert",
	2: "critical",
	3: "error",
	4: "warning",
	5: "notice",
	6: "informational",
	7: "debug",
}

// As per https://github.com/torvalds/linux/blob/master/tools/include/linux/kern_levels.h and syslog(3)
var severityLevelsShort = map[uint8]string{
	0: "emerg",
	1: "alert",
	2: "crit",
	3: "err",
	4: "warning",
	5: "notice",
	6: "info",
	7: "debug",
}

var facilities = map[uint8]string{
	0:  "kernel messages",
	1:  "user-level messages",
	2:  "mail system",
	3:  "system daemons",
	4:  "security/authorization messages",
	5:  "messages generated internally by syslogd",
	6:  "line printer subsystem",
	7:  "network news subsystem",
	8:  "UUCP subsystem",
	9:  "clock daemon",
	10: "security/authorization messages",
	11: "FTP daemon",
	12: "NTP subsystem",
	13: "log audit",
	14: "log alert",
	15: "clock daemon (note 2)", // (todo) > some sources reporting "scheduling daemon"
	16: "local use 0 (local0)",
	17: "local use 1 (local1)",
	18: "local use 2 (local2)",
	19: "local use 3 (local3)",
	20: "local use 4 (local4)",
	21: "local use 5 (local5)",
	22: "local use 6 (local6)",
	23: "local use 7 (local7)",
}

// As per syslog(3)
var facilityKeywords = map[uint8]string{
	0:  "kern",
	1:  "user",
	2:  "mail",
	3:  "daemon",
	4:  "auth",
	5:  "syslog",
	6:  "lpr",
	7:  "news",
	8:  "uucp",
	10: "authpriv",
	11: "ftp",
	15: "cron",
	16: "local0",
	17: "local1",
	18: "local2",
	19: "local3",
	20: "local4",
	21: "local5",
	22: "local6",
	23: "local7",
}

// Timestamp returns the syslog timestamp or nil when not set
func (sm *SyslogMessage) Timestamp() *time.Time {
	return sm.timestamp
}

// Hostname returns the syslog hostname or nil when not set
func (sm *SyslogMessage) Hostname() *string {
	return sm.hostname
}

// Content returns the syslog proc ID or nil when not set
func (sm *SyslogMessage) Content() *string {
	return sm.content
}

// ProcID returns the syslog proc ID or nil when not set
func (sm *SyslogMessage) ProcID() *string {
	return sm.Content()
}

// Tag returns the syslog tag or nil when not set
func (sm *SyslogMessage) Tag() *string {
	return sm.tag
}

// Appname returns the syslog appname or nil when not set
func (sm *SyslogMessage) Appname() *string {
	return sm.Tag()
}

// CiscoSequenceID returns the syslog msg ID or nil when not set
func (sm *SyslogMessage) CiscoSequenceID() *string {
	return sm.ciscoSequenceID
}

// MsgID returns the syslog msg ID or nil when not set
func (sm *SyslogMessage) MsgID() *string {
	return sm.CiscoSequenceID()
}

// Message returns the syslog message or nil when not set
func (sm *SyslogMessage) Message() *string {
	return sm.message
}

// StructuredData returns the syslog structured data or nil when not set
func (sm *SyslogMessage) StructuredData() *map[string]map[string]string {
	return nil
}

// ComputeFromPriority set the priority values and computes facility and severity from it.
//
// It does NOT check the input value validity.
func (sm *SyslogMessage) ComputeFromPriority(value uint8) {
	sm.priority = &value
	facility := uint8(value / 8)
	severity := uint8(value % 8)
	sm.facility = &facility
	sm.severity = &severity
}
