//line machine.go.rl:1
package legacysyslog

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/influxdata/go-syslog/v2/common"
)

var _ = fmt.Print

//line machine.go.rl:220

//line machine.go:22
const start int = 1
const first_final int = 109

const en_main int = 1

//line machine.go.rl:223

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	err        error
	bestEffort bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{}

//line machine.go.rl:238

//line machine.go.rl:239

//line machine.go.rl:240

//line machine.go.rl:241

//line machine.go.rl:242

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

//line machine.go:80
	{
		m.cs = start
	}

//line machine.go.rl:268

//line machine.go:87
	{
		if (m.p) == (m.pe) {
			goto _test_eof
		}
		switch m.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		}
		goto st_out
	st_case_1:
		switch (m.data)[(m.p)] {
		case 32:
			goto st2
		case 42:
			goto st3
		case 46:
			goto st3
		case 58:
			goto tr4
		case 60:
			goto st106
		case 65:
			goto tr6
		case 68:
			goto tr7
		case 70:
			goto tr8
		case 74:
			goto tr9
		case 77:
			goto tr10
		case 78:
			goto tr11
		case 79:
			goto tr12
		case 83:
			goto tr13
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch (m.data)[(m.p)] {
		case 32:
			goto st2
		case 42:
			goto st3
		case 46:
			goto st3
		case 65:
			goto tr6
		case 68:
			goto tr7
		case 70:
			goto tr8
		case 74:
			goto tr9
		case 77:
			goto tr10
		case 78:
			goto tr11
		case 79:
			goto tr12
		case 83:
			goto tr13
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr14
		}
		goto st0
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch (m.data)[(m.p)] {
		case 65:
			goto tr16
		case 68:
			goto tr17
		case 70:
			goto tr18
		case 74:
			goto tr19
		case 77:
			goto tr20
		case 78:
			goto tr21
		case 79:
			goto tr22
		case 83:
			goto tr23
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr15
		}
		goto st0
	tr14:
//line machine.go.rl:21

		m.pb = m.p

		goto st4
	tr15:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:466
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st6
		}
		goto st0
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st7
		}
		goto st0
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof7
		}
	st_case_7:
		if (m.data)[(m.p)] == 45 {
			goto st8
		}
		goto st0
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch (m.data)[(m.p)] {
		case 48:
			goto st9
		case 49:
			goto st48
		}
		goto st0
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof9
		}
	st_case_9:
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st10
		}
		goto st0
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		if (m.data)[(m.p)] == 45 {
			goto st11
		}
		goto st0
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch (m.data)[(m.p)] {
		case 48:
			goto st12
		case 51:
			goto st47
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st46
		}
		goto st0
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof12
		}
	st_case_12:
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof13
		}
	st_case_13:
		if (m.data)[(m.p)] == 84 {
			goto st14
		}
		goto st0
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof14
		}
	st_case_14:
		if (m.data)[(m.p)] == 50 {
			goto st45
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st15
		}
		goto st0
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st16
		}
		goto st0
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof16
		}
	st_case_16:
		if (m.data)[(m.p)] == 58 {
			goto st17
		}
		goto st0
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st18
		}
		goto st0
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st19
		}
		goto st0
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof19
		}
	st_case_19:
		if (m.data)[(m.p)] == 58 {
			goto st20
		}
		goto st0
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st21
		}
		goto st0
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st22
		}
		goto st0
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 46:
			goto st38
		case 90:
			goto st28
		}
		goto st0
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
		if (m.data)[(m.p)] == 50 {
			goto st37
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st24
		}
		goto st0
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st25
		}
		goto st0
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
		if (m.data)[(m.p)] == 58 {
			goto st26
		}
		goto st0
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof26
		}
	st_case_26:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st27
		}
		goto st0
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof27
		}
	st_case_27:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st28
		}
		goto st0
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr55
		case 91:
			goto st0
		}
		goto tr54
	tr108:
//line machine.go.rl:21

		m.pb = m.p

		goto st29
	tr54:
//line machine.go.rl:48

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st29
	tr116:
//line machine.go.rl:60

		{
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st29
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
//line machine.go:759
		switch (m.data)[(m.p)] {
		case 32:
			goto tr57
		case 91:
			goto st0
		}
		goto st29
	tr57:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st30
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line machine.go:779
		switch (m.data)[(m.p)] {
		case 32:
			goto tr59
		case 58:
			goto tr60
		case 91:
			goto tr61
		}
		goto tr58
	tr58:
//line machine.go.rl:21

		m.pb = m.p

		goto st31
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line machine.go:800
		switch (m.data)[(m.p)] {
		case 32:
			goto tr63
		case 58:
			goto tr63
		case 91:
			goto tr64
		}
		goto st31
	tr60:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st109
	tr63:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st109
	tr66:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st109
	tr69:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st109
	st109:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof109
		}
	st_case_109:
//line machine.go:851
		if (m.data)[(m.p)] == 32 {
			goto tr142
		}
		goto tr141
	tr141:
//line machine.go.rl:21

		m.pb = m.p

		goto st110
	st110:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof110
		}
	st_case_110:
//line machine.go:867
		goto st110
	tr142:
//line machine.go.rl:21

		m.pb = m.p

		goto st111
	st111:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof111
		}
	st_case_111:
//line machine.go:880
		goto tr141
	tr61:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st32
	tr64:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st32
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:905
		switch (m.data)[(m.p)] {
		case 32:
			goto tr66
		case 58:
			goto tr66
		case 93:
			goto tr67
		}
		goto tr65
	tr65:
//line machine.go.rl:21

		m.pb = m.p

		goto st33
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
//line machine.go:926
		switch (m.data)[(m.p)] {
		case 32:
			goto tr69
		case 58:
			goto tr69
		case 93:
			goto tr70
		}
		goto st33
	tr67:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st34
	tr70:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st34
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
//line machine.go:959
		switch (m.data)[(m.p)] {
		case 32:
			goto st109
		case 58:
			goto st109
		}
		goto st0
	tr59:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st112
	tr76:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st112
	tr153:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:21

		m.pb = m.p

		goto st112
	st112:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof112
		}
	st_case_112:
//line machine.go:1006
		switch (m.data)[(m.p)] {
		case 32:
			goto tr59
		case 58:
			goto tr60
		case 91:
			goto tr145
		}
		goto tr144
	tr144:
//line machine.go.rl:21

		m.pb = m.p

		goto st113
	st113:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof113
		}
	st_case_113:
//line machine.go:1027
		switch (m.data)[(m.p)] {
		case 32:
			goto tr63
		case 58:
			goto tr63
		case 91:
			goto tr147
		}
		goto st113
	tr145:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st114
	tr147:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st114
	st114:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof114
		}
	st_case_114:
//line machine.go:1060
		switch (m.data)[(m.p)] {
		case 32:
			goto tr66
		case 58:
			goto tr66
		case 93:
			goto tr149
		}
		goto tr148
	tr148:
//line machine.go.rl:21

		m.pb = m.p

		goto st115
	st115:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof115
		}
	st_case_115:
//line machine.go:1081
		switch (m.data)[(m.p)] {
		case 32:
			goto tr69
		case 58:
			goto tr69
		case 93:
			goto tr151
		}
		goto st115
	tr149:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st116
	tr151:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st116
	st116:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof116
		}
	st_case_116:
//line machine.go:1114
		switch (m.data)[(m.p)] {
		case 32:
			goto st109
		case 58:
			goto st109
		}
		goto st110
	tr55:
//line machine.go.rl:48

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st35
	tr109:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st35
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line machine.go:1161
		switch (m.data)[(m.p)] {
		case 32:
			goto tr73
		case 58:
			goto tr74
		case 91:
			goto tr61
		}
		goto tr72
	tr72:
//line machine.go.rl:21

		m.pb = m.p

		goto st36
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
//line machine.go:1182
		switch (m.data)[(m.p)] {
		case 32:
			goto tr76
		case 58:
			goto tr77
		case 91:
			goto tr64
		}
		goto st36
	tr74:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st117
	tr77:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st117
	st117:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof117
		}
	st_case_117:
//line machine.go:1215
		switch (m.data)[(m.p)] {
		case 32:
			goto tr153
		case 91:
			goto tr141
		}
		goto tr152
	tr152:
//line machine.go.rl:21

		m.pb = m.p

		goto st118
	st118:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof118
		}
	st_case_118:
//line machine.go:1234
		switch (m.data)[(m.p)] {
		case 32:
			goto tr155
		case 91:
			goto st110
		}
		goto st118
	tr155:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st119
	st119:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof119
		}
	st_case_119:
//line machine.go:1254
		switch (m.data)[(m.p)] {
		case 32:
			goto tr59
		case 58:
			goto tr60
		case 91:
			goto tr145
		}
		goto tr144
	tr73:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st120
	tr159:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st120
	st120:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof120
		}
	st_case_120:
//line machine.go:1296
		switch (m.data)[(m.p)] {
		case 32:
			goto tr73
		case 58:
			goto tr74
		case 91:
			goto tr145
		}
		goto tr156
	tr156:
//line machine.go.rl:21

		m.pb = m.p

		goto st121
	st121:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof121
		}
	st_case_121:
//line machine.go:1317
		switch (m.data)[(m.p)] {
		case 32:
			goto tr76
		case 58:
			goto tr77
		case 91:
			goto tr147
		}
		goto st121
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof37
		}
	st_case_37:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st25
		}
		goto st0
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof38
		}
	st_case_38:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st39
		}
		goto st0
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st40
		}
		goto st0
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st41
		}
		goto st0
	st41:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st42
		}
		goto st0
	st42:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st43
		}
		goto st0
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st44
		}
		goto st0
	st44:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch (m.data)[(m.p)] {
		case 43:
			goto st23
		case 45:
			goto st23
		case 90:
			goto st28
		}
		goto st0
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof45
		}
	st_case_45:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st16
		}
		goto st0
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof46
		}
	st_case_46:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st13
		}
		goto st0
	st47:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof47
		}
	st_case_47:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st13
		}
		goto st0
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof48
		}
	st_case_48:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st10
		}
		goto st0
	tr6:
//line machine.go.rl:21

		m.pb = m.p

		goto st49
	tr16:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st49
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof49
		}
	st_case_49:
//line machine.go:1506
		switch (m.data)[(m.p)] {
		case 112:
			goto st50
		case 117:
			goto st84
		}
		goto st0
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof50
		}
	st_case_50:
		if (m.data)[(m.p)] == 114 {
			goto st51
		}
		goto st0
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof51
		}
	st_case_51:
		if (m.data)[(m.p)] == 32 {
			goto st52
		}
		goto st0
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch (m.data)[(m.p)] {
		case 32:
			goto st53
		case 48:
			goto st53
		case 51:
			goto st83
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st82
		}
		goto st0
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof53
		}
	st_case_53:
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st54
		}
		goto st0
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof54
		}
	st_case_54:
		if (m.data)[(m.p)] == 32 {
			goto st55
		}
		goto st0
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof55
		}
	st_case_55:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st56
		}
		goto st0
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof56
		}
	st_case_56:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st57
		}
		goto st0
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof57
		}
	st_case_57:
		if (m.data)[(m.p)] == 58 {
			goto st71
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st58
		}
		goto st0
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof58
		}
	st_case_58:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st59
		}
		goto st0
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof59
		}
	st_case_59:
		if (m.data)[(m.p)] == 32 {
			goto st60
		}
		goto st0
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof60
		}
	st_case_60:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st61
		}
		goto st0
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof61
		}
	st_case_61:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st62
		}
		goto st0
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof62
		}
	st_case_62:
		if (m.data)[(m.p)] == 58 {
			goto st63
		}
		goto st0
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof63
		}
	st_case_63:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st64
		}
		goto st0
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof64
		}
	st_case_64:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st65
		}
		goto st0
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof65
		}
	st_case_65:
		if (m.data)[(m.p)] == 58 {
			goto st66
		}
		goto st0
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof66
		}
	st_case_66:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st67
		}
		goto st0
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof67
		}
	st_case_67:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st68
		}
		goto st0
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr107
		case 58:
			goto tr107
		}
		goto st0
	tr107:
//line machine.go.rl:72

		{
			t, err := time.Parse("Jan _2 2006 15:04:05", string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:149
		(m.p)--

		goto st69
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof69
		}
	st_case_69:
//line machine.go:1721
		switch (m.data)[(m.p)] {
		case 32:
			goto tr109
		case 58:
			goto tr110
		case 91:
			goto st0
		}
		goto tr108
	tr110:
//line machine.go.rl:21

		m.pb = m.p

		goto st70
	tr118:
//line machine.go.rl:60

		{
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st70
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof70
		}
	st_case_70:
//line machine.go:1760
		switch (m.data)[(m.p)] {
		case 32:
			goto tr109
		case 91:
			goto st0
		}
		goto tr108
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof71
		}
	st_case_71:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st72
		}
		goto st0
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof72
		}
	st_case_72:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st73
		}
		goto st0
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof73
		}
	st_case_73:
		if (m.data)[(m.p)] == 58 {
			goto st74
		}
		goto st0
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof74
		}
	st_case_74:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st75
		}
		goto st0
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof75
		}
	st_case_75:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st76
		}
		goto st0
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr117
		case 58:
			goto tr118
		case 91:
			goto st0
		}
		goto tr116
	tr117:
//line machine.go.rl:60

		{
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st77
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof77
		}
	st_case_77:
//line machine.go:1855
		switch (m.data)[(m.p)] {
		case 32:
			goto tr73
		case 58:
			goto tr74
		case 91:
			goto tr61
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr119
		}
		goto tr72
	tr119:
//line machine.go.rl:21

		m.pb = m.p

		goto st78
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof78
		}
	st_case_78:
//line machine.go:1879
		switch (m.data)[(m.p)] {
		case 32:
			goto tr76
		case 58:
			goto tr77
		case 91:
			goto tr64
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st79
		}
		goto st36
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr76
		case 58:
			goto tr77
		case 91:
			goto tr64
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st80
		}
		goto st36
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr76
		case 58:
			goto tr77
		case 91:
			goto tr64
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st81
		}
		goto st36
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr123
		case 58:
			goto tr77
		case 91:
			goto tr64
		}
		goto st36
	tr123:
//line machine.go.rl:84

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

//line machine.go.rl:160
		(m.p)--

//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st122
	st122:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof122
		}
	st_case_122:
//line machine.go:1974
		switch (m.data)[(m.p)] {
		case 32:
			goto tr73
		case 58:
			goto tr158
		case 91:
			goto tr145
		}
		goto tr156
	tr158:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st123
	st123:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof123
		}
	st_case_123:
//line machine.go:2000
		switch (m.data)[(m.p)] {
		case 32:
			goto tr159
		case 91:
			goto tr141
		}
		goto tr152
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof82
		}
	st_case_82:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st54
		}
		goto st0
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof83
		}
	st_case_83:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st54
		}
		goto st0
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof84
		}
	st_case_84:
		if (m.data)[(m.p)] == 103 {
			goto st51
		}
		goto st0
	tr7:
//line machine.go.rl:21

		m.pb = m.p

		goto st85
	tr17:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st85
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof85
		}
	st_case_85:
//line machine.go:2061
		if (m.data)[(m.p)] == 101 {
			goto st86
		}
		goto st0
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof86
		}
	st_case_86:
		if (m.data)[(m.p)] == 99 {
			goto st51
		}
		goto st0
	tr8:
//line machine.go.rl:21

		m.pb = m.p

		goto st87
	tr18:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st87
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof87
		}
	st_case_87:
//line machine.go:2101
		if (m.data)[(m.p)] == 101 {
			goto st88
		}
		goto st0
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof88
		}
	st_case_88:
		if (m.data)[(m.p)] == 98 {
			goto st51
		}
		goto st0
	tr9:
//line machine.go.rl:21

		m.pb = m.p

		goto st89
	tr19:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st89
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof89
		}
	st_case_89:
//line machine.go:2141
		switch (m.data)[(m.p)] {
		case 97:
			goto st90
		case 117:
			goto st91
		}
		goto st0
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof90
		}
	st_case_90:
		if (m.data)[(m.p)] == 110 {
			goto st51
		}
		goto st0
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch (m.data)[(m.p)] {
		case 108:
			goto st51
		case 110:
			goto st51
		}
		goto st0
	tr10:
//line machine.go.rl:21

		m.pb = m.p

		goto st92
	tr20:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st92
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof92
		}
	st_case_92:
//line machine.go:2196
		if (m.data)[(m.p)] == 97 {
			goto st93
		}
		goto st0
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch (m.data)[(m.p)] {
		case 114:
			goto st51
		case 121:
			goto st51
		}
		goto st0
	tr11:
//line machine.go.rl:21

		m.pb = m.p

		goto st94
	tr21:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st94
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof94
		}
	st_case_94:
//line machine.go:2239
		if (m.data)[(m.p)] == 111 {
			goto st95
		}
		goto st0
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof95
		}
	st_case_95:
		if (m.data)[(m.p)] == 118 {
			goto st51
		}
		goto st0
	tr12:
//line machine.go.rl:21

		m.pb = m.p

		goto st96
	tr22:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st96
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof96
		}
	st_case_96:
//line machine.go:2279
		if (m.data)[(m.p)] == 99 {
			goto st97
		}
		goto st0
	st97:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof97
		}
	st_case_97:
		if (m.data)[(m.p)] == 116 {
			goto st51
		}
		goto st0
	tr13:
//line machine.go.rl:21

		m.pb = m.p

		goto st98
	tr23:
//line machine.go.rl:39

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st98
	st98:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof98
		}
	st_case_98:
//line machine.go:2319
		if (m.data)[(m.p)] == 101 {
			goto st99
		}
		goto st0
	st99:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof99
		}
	st_case_99:
		if (m.data)[(m.p)] == 112 {
			goto st51
		}
		goto st0
	tr3:
//line machine.go.rl:21

		m.pb = m.p

		goto st100
	st100:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof100
		}
	st_case_100:
//line machine.go:2344
		if (m.data)[(m.p)] == 58 {
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st101
		}
		goto st0
	st101:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof101
		}
	st_case_101:
		if (m.data)[(m.p)] == 58 {
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st102
		}
		goto st0
	st102:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof102
		}
	st_case_102:
		if (m.data)[(m.p)] == 58 {
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st103
		}
		goto st0
	st103:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch (m.data)[(m.p)] {
		case 45:
			goto st8
		case 58:
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st104
		}
		goto st0
	st104:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof104
		}
	st_case_104:
		if (m.data)[(m.p)] == 58 {
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st104
		}
		goto st0
	tr4:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:34

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st105
	tr133:
//line machine.go.rl:34

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st105
	st105:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof105
		}
	st_case_105:
//line machine.go:2426
		if (m.data)[(m.p)] == 32 {
			goto st2
		}
		goto st0
	st106:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof106
		}
	st_case_106:
		if (m.data)[(m.p)] == 62 {
			goto tr138
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr137
		}
		goto st0
	tr137:
//line machine.go.rl:21

		m.pb = m.p

		goto st107
	st107:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof107
		}
	st_case_107:
//line machine.go:2454
		if (m.data)[(m.p)] == 62 {
			goto tr140
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st107
		}
		goto st0
	tr138:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:29

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st108
	tr140:
//line machine.go.rl:29

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st108
	st108:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof108
		}
	st_case_108:
//line machine.go:2485
		switch (m.data)[(m.p)] {
		case 32:
			goto st2
		case 42:
			goto st3
		case 46:
			goto st3
		case 58:
			goto tr4
		case 65:
			goto tr6
		case 68:
			goto tr7
		case 70:
			goto tr8
		case 74:
			goto tr9
		case 77:
			goto tr10
		case 78:
			goto tr11
		case 79:
			goto tr12
		case 83:
			goto tr13
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto st0
	st_out:
	_test_eof2:
		m.cs = 2
		goto _test_eof
	_test_eof3:
		m.cs = 3
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof5:
		m.cs = 5
		goto _test_eof
	_test_eof6:
		m.cs = 6
		goto _test_eof
	_test_eof7:
		m.cs = 7
		goto _test_eof
	_test_eof8:
		m.cs = 8
		goto _test_eof
	_test_eof9:
		m.cs = 9
		goto _test_eof
	_test_eof10:
		m.cs = 10
		goto _test_eof
	_test_eof11:
		m.cs = 11
		goto _test_eof
	_test_eof12:
		m.cs = 12
		goto _test_eof
	_test_eof13:
		m.cs = 13
		goto _test_eof
	_test_eof14:
		m.cs = 14
		goto _test_eof
	_test_eof15:
		m.cs = 15
		goto _test_eof
	_test_eof16:
		m.cs = 16
		goto _test_eof
	_test_eof17:
		m.cs = 17
		goto _test_eof
	_test_eof18:
		m.cs = 18
		goto _test_eof
	_test_eof19:
		m.cs = 19
		goto _test_eof
	_test_eof20:
		m.cs = 20
		goto _test_eof
	_test_eof21:
		m.cs = 21
		goto _test_eof
	_test_eof22:
		m.cs = 22
		goto _test_eof
	_test_eof23:
		m.cs = 23
		goto _test_eof
	_test_eof24:
		m.cs = 24
		goto _test_eof
	_test_eof25:
		m.cs = 25
		goto _test_eof
	_test_eof26:
		m.cs = 26
		goto _test_eof
	_test_eof27:
		m.cs = 27
		goto _test_eof
	_test_eof28:
		m.cs = 28
		goto _test_eof
	_test_eof29:
		m.cs = 29
		goto _test_eof
	_test_eof30:
		m.cs = 30
		goto _test_eof
	_test_eof31:
		m.cs = 31
		goto _test_eof
	_test_eof109:
		m.cs = 109
		goto _test_eof
	_test_eof110:
		m.cs = 110
		goto _test_eof
	_test_eof111:
		m.cs = 111
		goto _test_eof
	_test_eof32:
		m.cs = 32
		goto _test_eof
	_test_eof33:
		m.cs = 33
		goto _test_eof
	_test_eof34:
		m.cs = 34
		goto _test_eof
	_test_eof112:
		m.cs = 112
		goto _test_eof
	_test_eof113:
		m.cs = 113
		goto _test_eof
	_test_eof114:
		m.cs = 114
		goto _test_eof
	_test_eof115:
		m.cs = 115
		goto _test_eof
	_test_eof116:
		m.cs = 116
		goto _test_eof
	_test_eof35:
		m.cs = 35
		goto _test_eof
	_test_eof36:
		m.cs = 36
		goto _test_eof
	_test_eof117:
		m.cs = 117
		goto _test_eof
	_test_eof118:
		m.cs = 118
		goto _test_eof
	_test_eof119:
		m.cs = 119
		goto _test_eof
	_test_eof120:
		m.cs = 120
		goto _test_eof
	_test_eof121:
		m.cs = 121
		goto _test_eof
	_test_eof37:
		m.cs = 37
		goto _test_eof
	_test_eof38:
		m.cs = 38
		goto _test_eof
	_test_eof39:
		m.cs = 39
		goto _test_eof
	_test_eof40:
		m.cs = 40
		goto _test_eof
	_test_eof41:
		m.cs = 41
		goto _test_eof
	_test_eof42:
		m.cs = 42
		goto _test_eof
	_test_eof43:
		m.cs = 43
		goto _test_eof
	_test_eof44:
		m.cs = 44
		goto _test_eof
	_test_eof45:
		m.cs = 45
		goto _test_eof
	_test_eof46:
		m.cs = 46
		goto _test_eof
	_test_eof47:
		m.cs = 47
		goto _test_eof
	_test_eof48:
		m.cs = 48
		goto _test_eof
	_test_eof49:
		m.cs = 49
		goto _test_eof
	_test_eof50:
		m.cs = 50
		goto _test_eof
	_test_eof51:
		m.cs = 51
		goto _test_eof
	_test_eof52:
		m.cs = 52
		goto _test_eof
	_test_eof53:
		m.cs = 53
		goto _test_eof
	_test_eof54:
		m.cs = 54
		goto _test_eof
	_test_eof55:
		m.cs = 55
		goto _test_eof
	_test_eof56:
		m.cs = 56
		goto _test_eof
	_test_eof57:
		m.cs = 57
		goto _test_eof
	_test_eof58:
		m.cs = 58
		goto _test_eof
	_test_eof59:
		m.cs = 59
		goto _test_eof
	_test_eof60:
		m.cs = 60
		goto _test_eof
	_test_eof61:
		m.cs = 61
		goto _test_eof
	_test_eof62:
		m.cs = 62
		goto _test_eof
	_test_eof63:
		m.cs = 63
		goto _test_eof
	_test_eof64:
		m.cs = 64
		goto _test_eof
	_test_eof65:
		m.cs = 65
		goto _test_eof
	_test_eof66:
		m.cs = 66
		goto _test_eof
	_test_eof67:
		m.cs = 67
		goto _test_eof
	_test_eof68:
		m.cs = 68
		goto _test_eof
	_test_eof69:
		m.cs = 69
		goto _test_eof
	_test_eof70:
		m.cs = 70
		goto _test_eof
	_test_eof71:
		m.cs = 71
		goto _test_eof
	_test_eof72:
		m.cs = 72
		goto _test_eof
	_test_eof73:
		m.cs = 73
		goto _test_eof
	_test_eof74:
		m.cs = 74
		goto _test_eof
	_test_eof75:
		m.cs = 75
		goto _test_eof
	_test_eof76:
		m.cs = 76
		goto _test_eof
	_test_eof77:
		m.cs = 77
		goto _test_eof
	_test_eof78:
		m.cs = 78
		goto _test_eof
	_test_eof79:
		m.cs = 79
		goto _test_eof
	_test_eof80:
		m.cs = 80
		goto _test_eof
	_test_eof81:
		m.cs = 81
		goto _test_eof
	_test_eof122:
		m.cs = 122
		goto _test_eof
	_test_eof123:
		m.cs = 123
		goto _test_eof
	_test_eof82:
		m.cs = 82
		goto _test_eof
	_test_eof83:
		m.cs = 83
		goto _test_eof
	_test_eof84:
		m.cs = 84
		goto _test_eof
	_test_eof85:
		m.cs = 85
		goto _test_eof
	_test_eof86:
		m.cs = 86
		goto _test_eof
	_test_eof87:
		m.cs = 87
		goto _test_eof
	_test_eof88:
		m.cs = 88
		goto _test_eof
	_test_eof89:
		m.cs = 89
		goto _test_eof
	_test_eof90:
		m.cs = 90
		goto _test_eof
	_test_eof91:
		m.cs = 91
		goto _test_eof
	_test_eof92:
		m.cs = 92
		goto _test_eof
	_test_eof93:
		m.cs = 93
		goto _test_eof
	_test_eof94:
		m.cs = 94
		goto _test_eof
	_test_eof95:
		m.cs = 95
		goto _test_eof
	_test_eof96:
		m.cs = 96
		goto _test_eof
	_test_eof97:
		m.cs = 97
		goto _test_eof
	_test_eof98:
		m.cs = 98
		goto _test_eof
	_test_eof99:
		m.cs = 99
		goto _test_eof
	_test_eof100:
		m.cs = 100
		goto _test_eof
	_test_eof101:
		m.cs = 101
		goto _test_eof
	_test_eof102:
		m.cs = 102
		goto _test_eof
	_test_eof103:
		m.cs = 103
		goto _test_eof
	_test_eof104:
		m.cs = 104
		goto _test_eof
	_test_eof105:
		m.cs = 105
		goto _test_eof
	_test_eof106:
		m.cs = 106
		goto _test_eof
	_test_eof107:
		m.cs = 107
		goto _test_eof
	_test_eof108:
		m.cs = 108
		goto _test_eof

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 110, 113, 114, 115, 116, 118, 119, 121:
//line machine.go.rl:113

				output.messageSet = true
				output.message = string(m.text())

			case 109, 111, 112, 117, 120, 122, 123:
//line machine.go.rl:21

				m.pb = m.p

//line machine.go.rl:113

				output.messageSet = true
				output.message = string(m.text())

//line machine.go:2659
			}
		}

	_out:
		{
		}
	}

//line machine.go.rl:269

	return output.export(), m.err
}
