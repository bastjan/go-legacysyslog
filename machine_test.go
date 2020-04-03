package legacysyslog

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/go-syslog/v3"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	description string
	line        string
	expected    *SyslogMessage
}

func stringAddr(s string) *string {
	return &s
}

func uint8Addr(i uint8) *uint8 {
	return &i
}

func timeAddr(f, s string) *time.Time {
	time, err := time.Parse(f, s)
	if err != nil {
		panic("time invalid")
	}
	return &time
}

var tests = []testCase{
	testCase{
		description: "RFC3164",
		line:        `<13>Dec  1 09:15:22 0304eebf3c1e root[3037]: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
				ProcID:    stringAddr("3037"),
			},
			Tag:     stringAddr("root"),
			Content: stringAddr("3037"),
		},
	},
	testCase{
		description: "RFC3164 no space between tag and message (with content)",
		line:        `<13>Dec  1 09:15:22 0304eebf3c1e root[3037]:yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
				ProcID:    stringAddr("3037"),
			},
			Tag:     stringAddr("root"),
			Content: stringAddr("3037"),
		},
	},
	testCase{
		description: "RFC3164 no space between tag and message",
		line:        `<13>Dec  1 09:15:22 0304eebf3c1e root:yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 all caps month",
		line:        `<13>JUL  1 09:15:22 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Message:   stringAddr("yolo"),
				Timestamp: timeAddr(time.Stamp, "Jul  1 09:15:22"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},
	testCase{
		description: "RFC3164 no host",
		line:        `<13>Dec  1 09:15:22 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 IPv4 host",
		line:        `<13>Dec  1 09:15:22 10.77.0.1 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("10.77.0.1"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 IPv6 host",
		line:        `<13>Dec  1 09:15:22 2001:db8::8a2e:370:7334 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("2001:db8::8a2e:370:7334"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},
	testCase{
		description: "RFC3164 IPv6 host",
		line:        `<13>Dec  1 09:15:22 2001:0db8:0000:0000:0000:8a2e:0370:7334 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("2001:0db8:0000:0000:0000:8a2e:0370:7334"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},
	testCase{
		description: "RFC3164 IPv6 host",
		line:        `<13>Dec  1 09:15:22 ::1 root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("::1"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 space between priority and date",
		line:        `<13> Dec  1 09:15:22 0304eebf3c1e root[3037]: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
				ProcID:    stringAddr("3037"),
			},
			Tag:     stringAddr("root"),
			Content: stringAddr("3037"),
		},
	},
	testCase{
		description: "RFC3164 without content (procid)",
		line:        `<13>Dec  1 09:15:22 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 Cisco Year",
		line:        `<13>Dec  1 1999 09:15:22 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.RFC3339, "1999-12-01T09:15:22Z"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 Linksys Year",
		line:        `<13>Dec  1 09:15:22 1999 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.RFC3339, "1999-12-01T09:15:22Z"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "RFC3164 ISO Timestamp",
		line:        `<13>2019-12-04T09:55:56+01:00 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.RFC3339, "2019-12-04T09:55:56+01:00"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},
	testCase{
		description: "RFC3164 ISO Timestamp",
		line:        `<13>2019-12-04T09:55:56Z 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.RFC3339, "2019-12-04T09:55:56Z"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},
	testCase{
		description: "RFC3164 ISO Timestamp",
		line:        `<13>2019-12-04T09:55:56.234Z 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.RFC3339Nano, "2019-12-04T09:55:56.234Z"),
				Hostname:  stringAddr("0304eebf3c1e"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			Tag: stringAddr("root"),
		},
	},

	testCase{
		description: "Cisco Sequence ID",
		line:        `000123: *Dec  1 09:15:22: 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Hostname:  stringAddr("0304eebf3c1e"),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Message:   stringAddr("yolo"),
				Appname:   stringAddr("root"),
			},
			CiscoSequenceID:   stringAddr("000123"),
			CiscoTimestampExt: CiscoTimeClockModeUnsynced,
			Tag:               stringAddr("root"),
		},
	},

	testCase{
		description: "PulseSecure",
		line:        `<142> 2017-12-05T16:03:37+01:00  PulseSecure: 2017-12-05 16:03:37 - sslvpn01 - [87.245.121.26] xxx(VER-SMS)[] - Secondary authentication`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(142),
				Facility:  uint8Addr(17),
				Severity:  uint8Addr(6),
				Timestamp: timeAddr(time.RFC3339Nano, "2017-12-05T16:03:37+01:00"),
				Message:   stringAddr("2017-12-05 16:03:37 - sslvpn01 - [87.245.121.26] xxx(VER-SMS)[] - Secondary authentication"),
				Appname:   stringAddr("PulseSecure"),
			},
			Tag: stringAddr("PulseSecure"),
		},
	},

	testCase{
		description: "Seen from Symantec BlueCoat webproxy but configurable - no header at all",
		line:        `ProxyTest: Received:5 Sent:5`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Message: stringAddr("Received:5 Sent:5"),
				Appname: stringAddr("ProxyTest"),
			},
			Tag: stringAddr("ProxyTest"),
		},
	},
	testCase{
		description: "Message only (syslog-ng uses the first word as tag)",
		line:        `I like trains.`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Message: stringAddr("like trains."),
				Appname: stringAddr("I"),
			},
			Tag: stringAddr("I"),
		},
	},
	testCase{
		description: "Message only (syslog-ng uses the first word as tag)",
		line:        `I like: trains.`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Message: stringAddr("like: trains."),
				Appname: stringAddr("I"),
			},
			Tag: stringAddr("I"),
		},
	},
	testCase{
		description: "One word (syslog-ng uses the first word as tag)",
		line:        `test`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Appname: stringAddr("test"),
			},
			Tag: stringAddr("test"),
		},
	},
	testCase{
		description: "One word with '[' content mark (syslog-ng uses the first word as tag)",
		line:        `test[`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Appname: stringAddr("test"),
			},
			Tag: stringAddr("test"),
		},
	},
	testCase{
		description: "One word with '[234' content mark (syslog-ng uses the first word as tag)",
		line:        `test[234`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Appname: stringAddr("test"),
			},
			Tag: stringAddr("test"),
		},
	},
	testCase{
		description: "Empty",
		line:        ``,
		expected:    &SyslogMessage{},
	},
	testCase{
		description: "Hostname length limit 255",
		line:        `<13>Dec  1 09:15:22 ` + strings.Repeat("a", 500) + ` test`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority:  uint8Addr(13),
				Facility:  uint8Addr(1),
				Severity:  uint8Addr(5),
				Timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
				Message:   stringAddr("test"),
				Appname:   stringAddr(strings.Repeat("a", 500)),
			},
			Tag: stringAddr(strings.Repeat("a", 500)),
		},
	},

	testCase{
		description: "Citrix Netscaler (does not work in syslog-ng) tested here to ensure parsing is the same as in syslog-ng",
		line:        `<134> 11/28/2019:15:31:21 GMT netscaler1 0-PPE-0 : default TCP CONN_TERMINATE 17000000 0 :  Source 127.0.0.1:80 - Destination 127.0.0.1:25963 - Start Time 11/28/2019:15:30:06 GMT - End Time 11/28/2019:15:31:21 GMT - Total_bytes_send 1 - Total_bytes_recv 1`,
		expected: &SyslogMessage{
			Base: syslog.Base{
				Priority: uint8Addr(134),
				Facility: uint8Addr(16),
				Severity: uint8Addr(6),
				Message:  stringAddr("15:31:21 GMT netscaler1 0-PPE-0 : default TCP CONN_TERMINATE 17000000 0 :  Source 127.0.0.1:80 - Destination 127.0.0.1:25963 - Start Time 11/28/2019:15:30:06 GMT - End Time 11/28/2019:15:31:21 GMT - Total_bytes_send 1 - Total_bytes_recv 1"),
				Appname:  stringAddr("11/28/2019"),
			},
			Tag: stringAddr("11/28/2019"),
		},
	},
}

func TestMachineParse(t *testing.T) {
	m := NewMachine()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			msg, err := m.Parse([]byte(test.line))
			if err != nil {
				t.Error("Test failed:", err)
			}
			require.Equal(t, test.expected, msg)
		})
	}
}

func TestSetLocation(t *testing.T) {
	line := `<13>Dec  1 09:55:56 0304eebf3c1e root: yolo`

	locations := []*time.Location{
		time.UTC,
		time.Local,
		time.FixedZone("testzone", int(8*(time.Hour/time.Second))),
	}

	m := NewMachine()
	for _, loc := range locations {
		m.SetLocation(loc)

		expected, _ := time.ParseInLocation(time.Stamp, "Dec  1 09:55:56", loc)

		msg, err := m.Parse([]byte(line))
		require.NoError(t, err)
		require.Equal(t, expected, *msg.(*SyslogMessage).Base.Timestamp)
	}
}

func TestWriteFuzzerCorpus(t *testing.T) {
	err := os.MkdirAll("corpus", 0777)
	require.NoError(t, err)

	for _, test := range tests {
		name := url.QueryEscape(test.description) + ".from_test"
		path := filepath.Join("corpus", name)
		err := ioutil.WriteFile(path, []byte(test.line), 0666)
		require.NoError(t, err)
	}
}

var benchres syslog.Message

func BenchmarkParseStandardMessage(b *testing.B) {
	m := NewMachine()

	line := []byte(`<13>Dec  1 09:15:22 0304eebf3c1e root[3037]: yolo`)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		benchres, _ = m.Parse(line)
	}
}

func BenchmarkParseLinksysYearMessage(b *testing.B) {
	m := NewMachine()

	line := []byte(`<13>Dec  1 09:15:22 2017 0304eebf3c1e root[3037]: yolo`)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		benchres, _ = m.Parse(line)
	}
}
