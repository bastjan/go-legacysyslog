package protocolchooser_test

import (
	"testing"

	"github.com/bastjan/go-legacysyslog/utils/protocolchooser"
)

type testCase struct {
	line     string
	protocol protocolchooser.Protocol
	error    error
}

var testCases = []testCase{
	testCase{
		line:     `<`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<1`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<12`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<123`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<123>`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<123>1`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<123>12`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `<123>123`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `123 `,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},
	testCase{
		line:     `123 <123>`,
		protocol: protocolchooser.Unknown,
		error:    protocolchooser.ErrShortBuf,
	},

	testCase{
		line:     `<>`,
		protocol: protocolchooser.Unknown,
		error:    nil,
	},
	testCase{
		line:     `testblub`,
		protocol: protocolchooser.Unknown,
		error:    nil,
	},
	testCase{
		line:     `1234: testblub`,
		protocol: protocolchooser.Unknown,
		error:    nil,
	},

	testCase{
		line:     `<123> test`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},
	testCase{
		line:     `<123>123: xxx`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},
	testCase{
		line:     `<123>0`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},
	testCase{
		line:     `<123>00`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},
	testCase{
		line:     `<123>000`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},
	testCase{
		line:     `<123>0000`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},

	testCase{
		line:     `1234 <123> test`,
		protocol: protocolchooser.OctetCountingBSD,
		error:    nil,
	},

	testCase{
		line:     `<123>1 test`,
		protocol: protocolchooser.IETF,
		error:    nil,
	},
	testCase{
		line:     `<123>12 test`,
		protocol: protocolchooser.IETF,
		error:    nil,
	},
	testCase{
		line:     `<123>123 test`,
		protocol: protocolchooser.IETF,
		error:    nil,
	},
	testCase{
		line:     `<123>1234 test`,
		protocol: protocolchooser.BSD,
		error:    nil,
	},

	testCase{
		line:     `1234 <123>1 test`,
		protocol: protocolchooser.OctetCountingIETF,
		error:    nil,
	},
}

func TestChoose(t *testing.T) {
	for _, testCase := range testCases {
		pr, err := protocolchooser.Choose([]byte(testCase.line))
		if err != testCase.error {
			t.Errorf("subject '%s': expected error '%s' got '%s'", testCase.line, testCase.error, err)
		}
		if pr != testCase.protocol {
			t.Errorf("subject '%s': expected protocol '%s' got '%s'", testCase.line, testCase.protocol, pr)
		}

	}
}
