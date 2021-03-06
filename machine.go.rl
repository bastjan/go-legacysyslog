package legacysyslog

import (
	"fmt"
	"time"

	"github.com/influxdata/go-syslog/v3/common"
	"github.com/influxdata/go-syslog/v3"
)

var _ = fmt.Print

%%{
machine legacysyslog;
include iso_timestamp "iso_timestamp.rl";

# unsigned alphabet
alphtype uint8;

action mark {
	m.pb = m.p
}

action mark_hostname {
	m.pHostname = m.p
}

action mark_timestamp {
	m.pTimestamp = m.p
}

action set_priority {
	output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
	output.prioritySet = true
}

action set_cisco_sequence_id {
	output.ciscoSequenceIDSet = true
	output.ciscoSequenceID = string(m.text())
}

action set_cisco_timestamp_ext {
	switch m.data[m.p-1] {
	case '.':
		output.ciscoTimestampExt = CiscoTimeClockModeSynced
	case '*':
		output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
	}
}

action set_iso_timestamp {
	{
		t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		output.timestampSet = true
		output.timestamp = t
	}
}

action set_bsd_timestamp {
	{
		t, err := time.ParseInLocation(time.Stamp, string(m.data[m.pTimestamp:m.p]), m.location)
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		output.timestampSet = true
		output.timestamp = t
	}
}

action set_cisco_timestamp {
	{
		t, err := time.ParseInLocation("Jan _2 2006 15:04:05", string(m.data[m.pTimestamp:m.p]), m.location)
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		output.timestampSet = true
		output.timestamp = t
	}
}

action append_linksys_year {
	{
		year := common.UnsafeUTF8DecimalCodePointsToInt(m.text())
		output.timestamp = output.timestamp.AddDate(year, 0, 0)
	}
}

action set_hostname {
	output.hostnameSet = true
	output.hostname = string(m.data[m.pHostname:m.p])
}

action set_tag {
	output.tagSet = true
	output.tag = string(m.text())
}

action set_content {
	output.contentSet = true
	output.content = string(m.text())
}

action set_message {
	output.messageSet = true
	output.message = string(m.text())
}

action print_err {
	fmt.Println("-- ERROR --")
	fmt.Println("POS: ", m.p)
	fmt.Println("Last mark: ", m.pb)
	fmt.Printf("Str: '%s'\n", string(m.text()))
	fmt.Printf("Before: '%s'\n", string(m.data[:m.p]))
	fmt.Printf("After: '%s'\n", string(m.data[m.p:]))
	fmt.Print("-- ERROR --\n\n")

	return fmt.Errorf("failed at %d", m.p)
}

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L51
priority = '<' digit* >mark %set_priority '>';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L185
cisco_sequence_id = digit* >mark %set_cisco_sequence_id ': ';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L214
cisco_timestamp_attributes = [*.] %set_cisco_timestamp_ext;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L76
month_abrev = ('J' /an/i | 'F' /eb/i | 'M' /ar/i | 'A' /pr/i | 'M' /ay/i | 'J' /un/i | 'J' /ul/i | 'A' /ug/i | 'S' /ep/i | 'O' /ct/i | 'N' /ov/i | 'D' /ec/i);
bsd_day = ([0 ] . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');
# TODO: syslog-ng actually allows space instead of 0
hhmmss = digit{2} ':' digit{2} ':' digit{2};

cisco_timestamp =
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L220
	(month_abrev ' ' bsd_day ' ' digit{4} ' ' hhmmss) >mark_timestamp %set_cisco_timestamp
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L211
	[: ] >{ fhold; }
	;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L145
bsd_stamp = (month_abrev ' ' bsd_day ' ' hhmmss) >mark_timestamp %set_bsd_timestamp;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L261
linksys_timestamp =
	# we have 2 final states here - the bsd stamp begins with the same characters
	# so set_timestamp and append_linksys_year is called
	(bsd_stamp ' ' digit{4}) >mark %append_linksys_year
	' ' >{ fhold; };

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L391
rfc3164_stamp = (cisco_timestamp | linksys_timestamp | bsd_stamp );

iso_stamp = iso_timestamp >mark_timestamp %set_iso_timestamp;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L426
timestamp = 
	iso_stamp | rfc3164_stamp
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L447
	':'?
	;

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L401
ipv6heuristic = (xdigit{0,4} ':'){1,7} xdigit{1,4};

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L456
# TODO: check hostname and badhostname options
hostname = ( (any - [\[ :]){1,255} | ipv6heuristic) >mark_hostname %set_hostname ' ';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L315
tag = (any - [\[: ])* >mark %set_tag;

content = 
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L330
	'['
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L333
	(any - [\]: ])* >mark %set_content
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L342
	']'? ;

tag_optional_content = tag %/set_tag content? [: ] ' '?;

main := 
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L757
	priority?
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L762
	cisco_sequence_id?
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L763
	' '*
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L764
	cisco_timestamp_attributes?

	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L767
	(	timestamp
		# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L775
		' '*
		# TODO: last message repeated / forwarded from
		hostname?
		# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L804
		' '*
		# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L808
	)?

	tag_optional_content

	# TODO validate utf8 option
	any* >mark %set_message

	$!print_err
	;

}%%

%% write data noerror noprefix;

type machine struct {
	data         []byte
	cs           int
	p, pe, eof   int
	// general mark
	pb           int
	// hostname machine can run in parallel to the tag machine since it is not clear from the beginning
	// if it is a hostname or tag
	pHostname int
	// same for the timestamp
	pTimestamp int

	err          error
	bestEffort   bool

	location *time.Location
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{location: time.UTC}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;

	return m
}

// Err returns the error that occurred on the last call to Parse.
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

// Parse parses the input byte array as a RFC3164 syslog message.
func (m *machine) Parse(input []byte) (syslog.Message, error) {
	if len(input) == 0 {
		return (&syslogMessage{}).export(), nil
	}

	m.data = input
	m.p = 0
	m.pb = 0
	m.pHostname = 0
	m.pTimestamp = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	output := &syslogMessage{}

	%% write init;
	%% write exec;

	return output.export(), m.err
}

// WithBestEffort enables best effort mode. Has no effect since
// legacysyslog parsing is always best effort.
func (m *machine) WithBestEffort() {
}

// SetLocation sets the currently used location if a timestamp does not include timezone information.
// Default is UTC.
func (m *machine) SetLocation(l *time.Location) {
	m.location = l
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
// Always returns true since legacysyslog parsing is always best effort.
func (m *machine) HasBestEffort() bool {
	return true
}
