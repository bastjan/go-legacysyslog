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
	syslog.Base

	CiscoSequenceID   *string
	CiscoTimestampExt CiscoTimeClockMode
	Tag               *string
	Content           *string
}

var _ syslog.Message = &SyslogMessage{}

func (sm *syslogMessage) export() *SyslogMessage {
	out := &SyslogMessage{
		CiscoTimestampExt: sm.ciscoTimestampExt,
	}

	if sm.prioritySet {
		out.ComputeFromPriority(sm.priority)
	}
	if sm.ciscoSequenceIDSet {
		out.CiscoSequenceID = &sm.ciscoSequenceID
	}

	if sm.timestampSet {
		out.Timestamp = &sm.timestamp
	}
	if sm.hostnameSet {
		out.Hostname = &sm.hostname
	}
	if sm.tagSet {
		out.Tag = &sm.tag
		out.Appname = out.Tag
	}
	if sm.contentSet {
		out.Content = &sm.content
		out.ProcID = out.Content
	}
	if sm.messageSet {
		out.Message = &sm.message
	}

	return out
}
