package legacysyslog

import (
	"fmt"
	"testing"
	"time"

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
			priority:  uint8Addr(13),
			facility:  uint8Addr(1),
			severity:  uint8Addr(5),
			timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
			hostname:  stringAddr("0304eebf3c1e"),
			tag:       stringAddr("root"),
			content:   stringAddr("3037"),
			message:   stringAddr("yolo"),
		},
	},
	testCase{
		description: "RFC3164 space between priority and date",
		line:        `<13> Dec  1 09:15:22 0304eebf3c1e root[3037]: yolo`,
		expected: &SyslogMessage{
			priority:  uint8Addr(13),
			facility:  uint8Addr(1),
			severity:  uint8Addr(5),
			timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
			hostname:  stringAddr("0304eebf3c1e"),
			tag:       stringAddr("root"),
			content:   stringAddr("3037"),
			message:   stringAddr("yolo"),
		},
	},
	testCase{
		description: "RFC3164 without content (procid)",
		line:        `<13>Dec  1 09:15:22 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			priority:  uint8Addr(13),
			facility:  uint8Addr(1),
			severity:  uint8Addr(5),
			timestamp: timeAddr(time.Stamp, "Dec  1 09:15:22"),
			hostname:  stringAddr("0304eebf3c1e"),
			tag:       stringAddr("root"),
			message:   stringAddr("yolo"),
		},
	},

	testCase{
		description: "RFC3164 Cisco Year",
		line:        `<13>Dec  1 1999 09:15:22 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			priority:  uint8Addr(13),
			facility:  uint8Addr(1),
			severity:  uint8Addr(5),
			timestamp: timeAddr(time.RFC3339, "1999-12-01T09:15:22Z"),
			hostname:  stringAddr("0304eebf3c1e"),
			tag:       stringAddr("root"),
			message:   stringAddr("yolo"),
		},
	},

	testCase{
		description: "RFC3164 Linksys Year",
		line:        `<13>Dec  1 09:15:22 1999 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			priority:  uint8Addr(13),
			facility:  uint8Addr(1),
			severity:  uint8Addr(5),
			timestamp: timeAddr(time.RFC3339, "1999-12-01T09:15:22Z"),
			hostname:  stringAddr("0304eebf3c1e"),
			tag:       stringAddr("root"),
			message:   stringAddr("yolo"),
		},
	},

	testCase{
		description: "Cisco Sequence ID",
		line:        `000123: *Dec  1 09:15:22: 0304eebf3c1e root: yolo`,
		expected: &SyslogMessage{
			hostname:          stringAddr("0304eebf3c1e"),
			ciscoSequenceID:   stringAddr("000123"),
			timestamp:         timeAddr(time.Stamp, "Dec  1 09:15:22"),
			ciscoTimestampExt: CiscoTimeClockModeUnsynced,
			tag:               stringAddr("root"),
			message:           stringAddr("yolo"),
		},
	},
}

func TestMachineParse(t *testing.T) {
	m := NewMachine()

	for _, test := range tests {
		fmt.Println("TestCase:", test.description)
		t.Log("TestCase:", test.description)
		msg, err := m.Parse([]byte(test.line))
		t.Logf("Output: %+v", msg)
		if err != nil {
			t.Error("Test failed:", err)
		}
		require.Equal(t, test.expected, msg)
	}
}
