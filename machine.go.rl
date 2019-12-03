package legacysyslog

import (
	"fmt"
	"time"
	"bytes"
	"strconv"

	"github.com/influxdata/go-syslog/v2/common"
)

var _ = fmt.Print

%%{
machine legacysyslog;

# unsigned alphabet
alphtype uint8;

action mark {
	m.pb = m.p
}

action print {
	fmt.Println("Val: ", string(m.text()))
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

action set_bsd_timestamp {
	{
		output.timestampSet = true
		t, err := time.Parse(time.Stamp, string(m.text()))
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		output.timestamp = t
	}
}

action set_cisco_timestamp {
	{
		output.timestampSet = true
		t, err := time.Parse("Jan _2 2006 15:04:05", string(m.text()))
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		output.timestamp = t
	}
}

action append_linksys_year {
	{
		yearRaw := bytes.Trim(m.text(), " ")
		year, err := strconv.Atoi(string(yearRaw))
		if err != nil {
			m.err = err
			return output.export(), m.err
		}
		// Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
		t := output.timestamp.UTC()
		output.timestamp = time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	}
}

action set_hostname {
	output.hostnameSet = true
	output.hostname = string(m.text())
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

# TODO: Case insensitive 
month_abrev = ('Jan' | 'Feb' | 'Mar' | 'Apr' | 'May' | 'Jun' | 'Jul' | 'Aug' | 'Sep' | 'Oct' | 'Nov' | 'Dec');
bsd_day = ([0 ] . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');
# TODO: syslog-ng actually allows space instead of 0
hhmmss = digit{2} ':' digit{2} ':' digit{2};

cisco_timestamp =
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L220
	(month_abrev ' ' bsd_day ' ' digit{4} ' ' hhmmss) >mark %set_cisco_timestamp
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L211
	[: ] >{ fhold; }
	;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L145
bsd_stamp = (month_abrev ' ' bsd_day ' ' hhmmss) >mark %set_bsd_timestamp;

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L261
linksys_timestamp =
	# we have 2 final states here - the bsd stamp begins with the same characters
	# so set_timestamp and append_linksys_year is called
	(bsd_stamp ' ' digit{4}) >mark %append_linksys_year
	' ' >{ fhold; };

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L391
rfc3164_stamp = (cisco_timestamp | linksys_timestamp | bsd_stamp );

iso_stamp = 'TODO';

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L426
timestamp = 
	iso_stamp | rfc3164_stamp
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L447
	':'?
	;

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L456
# TODO: ipv6 heuristics
# TODO: check hostname and badhostname options
hostname = (any - [\[ ])* >mark %set_hostname ' ';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L315
tag = (any - [\[: ])* >mark %set_tag;

content = 
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L330
	'['
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L333
	(any - [\]: ])* >mark %set_content
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L342
	']'? ;

tag_optional_content = tag content? [: ] ' '?;

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
	# TODO: no timestamp case
	timestamp
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L775
	' '*
	# TODO: last message repeated / forwarded from
	hostname
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L804
	' '*
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L808
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
	pb           int
	err          error
	bestEffort   bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{}

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
func (m *machine) Parse(input []byte) (*SyslogMessage, error) {
	m.data = input
	m.p = 0
	m.pb = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	output := &syslogMessage{}

	%% write init;
	%% write exec;

	return output.export(), m.err
}
