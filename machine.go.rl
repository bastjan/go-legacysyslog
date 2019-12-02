package legacysyslog

import (
	"fmt"
)

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

action print_err {
	// cs:1 p:26 pe:55 eof:55 pb:8
	fmt.Println("-- ERROR --")
	fmt.Println("POS: ", m.p)
	fmt.Println("Last mark: ", m.pb)
	fmt.Printf("Str: '%s'\n", string(m.text()))
	fmt.Printf("Before: '%s'\n", string(m.data[:m.p]))
	fmt.Printf("After: '%s'\n", string(m.data[m.p:]))

	return fmt.Errorf("failed at %d", m.p)
}

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L51
priority = '<' digit* >mark %print '>';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L185
cisco_sequence_id = digit* >mark %print ': ';

# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L214
cisco_timestamp_attributes = [*.] >mark %print;

# TODO: Case insensitive 
month_abrev = ('Jan' | 'Feb' | 'Mar' | 'Apr' | 'May' | 'Jun' | 'Jul' | 'Aug' | 'Sep' | 'Oct' | 'Nov' | 'Dec');
bsd_day = ([0 ] . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');
hhmmss = digit{2} ':' digit{2} ':' digit{2};

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L220
cisco_timestamp = (month_abrev ' ' bsd_day ' ' digit{4} ' ' hhmmss);

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L261
linksys_timestamp = (month_abrev ' ' bsd_day ' ' hhmmss ' ' digit{4} ' ');

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L145
bsd_stamp = (month_abrev ' ' bsd_day ' ' hhmmss);

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L391
rfc3164_stamp = (cisco_timestamp | linksys_timestamp | bsd_stamp) >mark %print;

iso_stamp = 'TODO';

# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L426
timestamp = 
	iso_stamp | rfc3164_stamp
	# https://github.com/syslog-ng/syslog-ng/blob/eedebbfd3fc9d14389abf53c7efead1ecfea8d12/lib/timeutils/scan-timestamp.c#L447
	':'?
	;

# TODO
host = graph+ >mark %print;  

main := 
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L757
	priority?
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L762
	cisco_sequence_id?
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L763
	' '*
	# https://github.com/syslog-ng/syslog-ng/blob/3a1bda0d9a9e42b5cd7e5a02ca05f5f896ef82b6/modules/syslogformat/syslog-format.c#L764
	cisco_timestamp_attributes?

	# TODO from here
	timestamp
	' '*
	host

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
func (m *machine) Parse(input []byte) error {
	m.data = input
	m.p = 0
	m.pb = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	%% write init;
	%% write exec;

	return nil
}
