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

//line machine.go.rl:222

//line machine.go:22
const start int = 1
const first_final int = 103

const en_main int = 1

//line machine.go.rl:225

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	// general mark
	pb int
	// hostname machine can run in parallel to the tag machine since it is not clear from the beginning
	// if it is a hostname or tag
	pHostname int

	err        error
	bestEffort bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{}

//line machine.go.rl:245

//line machine.go.rl:246

//line machine.go.rl:247

//line machine.go.rl:248

//line machine.go.rl:249

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
	m.pHostname = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	output := &syslogMessage{}

//line machine.go:86
	{
		m.cs = start
	}

//line machine.go.rl:276

//line machine.go:93
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
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 111:
			goto st_case_111
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
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
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
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
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
			goto st100
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
//line machine.go:542
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
		case 58:
			goto tr57
		case 91:
			goto tr58
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr56
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr56
			}
		default:
			goto tr56
		}
		goto tr54
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st29
	tr104:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st29
	tr114:
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st29
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
//line machine.go:861
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr61
		case 91:
			goto tr62
		}
		goto st29
	tr60:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st103
	tr137:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st103
	st103:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof103
		}
	st_case_103:
//line machine.go:899
		switch (m.data)[(m.p)] {
		case 32:
			goto tr137
		case 58:
			goto tr138
		case 91:
			goto tr139
		}
		goto tr136
	tr136:
//line machine.go.rl:21

		m.pb = m.p

		goto st104
	st104:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof104
		}
	st_case_104:
//line machine.go:920
		switch (m.data)[(m.p)] {
		case 32:
			goto tr61
		case 58:
			goto tr61
		case 91:
			goto tr141
		}
		goto st104
	tr61:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st105
	tr64:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st105
	tr67:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st105
	tr138:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st105
	st105:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof105
		}
	st_case_105:
//line machine.go:971
		if (m.data)[(m.p)] == 32 {
			goto tr143
		}
		goto tr142
	tr142:
//line machine.go.rl:21

		m.pb = m.p

		goto st106
	st106:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof106
		}
	st_case_106:
//line machine.go:987
		goto st106
	tr143:
//line machine.go.rl:21

		m.pb = m.p

		goto st107
	st107:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof107
		}
	st_case_107:
//line machine.go:1000
		goto tr142
	tr141:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st108
	tr139:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st108
	st108:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof108
		}
	st_case_108:
//line machine.go:1025
		switch (m.data)[(m.p)] {
		case 32:
			goto tr64
		case 58:
			goto tr64
		case 93:
			goto tr146
		}
		goto tr145
	tr145:
//line machine.go.rl:21

		m.pb = m.p

		goto st109
	st109:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof109
		}
	st_case_109:
//line machine.go:1046
		switch (m.data)[(m.p)] {
		case 32:
			goto tr67
		case 58:
			goto tr67
		case 93:
			goto tr148
		}
		goto st109
	tr146:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st110
	tr148:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st110
	st110:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof110
		}
	st_case_110:
//line machine.go:1079
		switch (m.data)[(m.p)] {
		case 32:
			goto st105
		case 58:
			goto st105
		}
		goto st106
	tr58:
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

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st30
	tr62:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st30
	tr108:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st30
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

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st30
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line machine.go:1156
		switch (m.data)[(m.p)] {
		case 32:
			goto tr64
		case 58:
			goto tr64
		case 93:
			goto tr65
		}
		goto tr63
	tr63:
//line machine.go.rl:21

		m.pb = m.p

		goto st31
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line machine.go:1177
		switch (m.data)[(m.p)] {
		case 32:
			goto tr67
		case 58:
			goto tr67
		case 93:
			goto tr68
		}
		goto st31
	tr65:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st32
	tr68:
//line machine.go.rl:108

		output.contentSet = true
		output.content = string(m.text())

		goto st32
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:1210
		switch (m.data)[(m.p)] {
		case 32:
			goto st105
		case 58:
			goto st105
		}
		goto st0
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

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st111
	tr105:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st111
	tr107:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st111
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st111
	st111:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof111
		}
	st_case_111:
//line machine.go:1299
		switch (m.data)[(m.p)] {
		case 32:
			goto tr105
		case 58:
			goto tr151
		case 91:
			goto tr139
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr150
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr150
			}
		default:
			goto tr150
		}
		goto tr149
	tr149:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st112
	st112:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof112
		}
	st_case_112:
//line machine.go:1336
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr61
		case 91:
			goto tr141
		}
		goto st112
	tr150:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st113
	st113:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof113
		}
	st_case_113:
//line machine.go:1361
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st114
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st114
			}
		default:
			goto st114
		}
		goto st112
	st114:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st115
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st115
			}
		default:
			goto st115
		}
		goto st112
	st115:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st116
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st116
			}
		default:
			goto st116
		}
		goto st112
	st116:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		goto st112
	tr57:
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st117
	tr71:
//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st117
	tr151:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st117
	st117:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof117
		}
	st_case_117:
//line machine.go:1503
		switch (m.data)[(m.p)] {
		case 32:
			goto tr143
		case 58:
			goto tr157
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr156
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr156
			}
		default:
			goto tr156
		}
		goto tr142
	tr156:
//line machine.go.rl:21

		m.pb = m.p

		goto st118
	st118:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof118
		}
	st_case_118:
//line machine.go:1534
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st123
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st120
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st120
			}
		default:
			goto st120
		}
		goto st106
	tr158:
//line machine.go.rl:98

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

		goto st119
	st119:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof119
		}
	st_case_119:
//line machine.go:1566
		switch (m.data)[(m.p)] {
		case 32:
			goto tr137
		case 58:
			goto tr138
		case 91:
			goto tr139
		}
		goto tr136
	st120:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof120
		}
	st_case_120:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st123
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st121
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st121
			}
		default:
			goto st121
		}
		goto st106
	st121:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof121
		}
	st_case_121:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st123
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st122
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st122
			}
		default:
			goto st122
		}
		goto st106
	st122:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof122
		}
	st_case_122:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st123
		}
		goto st106
	tr157:
//line machine.go.rl:21

		m.pb = m.p

		goto st123
	st123:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof123
		}
	st_case_123:
//line machine.go:1647
		if (m.data)[(m.p)] == 58 {
			goto st128
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st124
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st124
			}
		default:
			goto st124
		}
		goto st106
	st124:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st128
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st125
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st125
			}
		default:
			goto st125
		}
		goto st106
	st125:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof125
		}
	st_case_125:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st128
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st126
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto st106
	st126:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st128
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st127
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st127
			}
		default:
			goto st127
		}
		goto st106
	st127:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st128
		}
		goto st106
	st128:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof128
		}
	st_case_128:
		if (m.data)[(m.p)] == 58 {
			goto st133
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st129
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st129
			}
		default:
			goto st129
		}
		goto st106
	st129:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st133
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st130
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st130
			}
		default:
			goto st130
		}
		goto st106
	st130:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st133
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st131
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st131
			}
		default:
			goto st131
		}
		goto st106
	st131:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st133
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st132
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st132
			}
		default:
			goto st132
		}
		goto st106
	st132:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st133
		}
		goto st106
	st133:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof133
		}
	st_case_133:
		if (m.data)[(m.p)] == 58 {
			goto st138
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st134
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st134
			}
		default:
			goto st134
		}
		goto st106
	st134:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st138
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st135
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st135
			}
		default:
			goto st135
		}
		goto st106
	st135:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st138
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st136
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st136
			}
		default:
			goto st136
		}
		goto st106
	st136:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st138
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st137
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st137
			}
		default:
			goto st137
		}
		goto st106
	st137:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st138
		}
		goto st106
	st138:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof138
		}
	st_case_138:
		if (m.data)[(m.p)] == 58 {
			goto st143
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st139
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st139
			}
		default:
			goto st139
		}
		goto st106
	st139:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof139
		}
	st_case_139:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st143
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st140
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st140
			}
		default:
			goto st140
		}
		goto st106
	st140:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof140
		}
	st_case_140:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st143
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st141
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st141
			}
		default:
			goto st141
		}
		goto st106
	st141:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof141
		}
	st_case_141:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st143
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st142
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st142
			}
		default:
			goto st142
		}
		goto st106
	st142:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof142
		}
	st_case_142:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st143
		}
		goto st106
	st143:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof143
		}
	st_case_143:
		if (m.data)[(m.p)] == 58 {
			goto st148
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st144
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st144
			}
		default:
			goto st144
		}
		goto st106
	st144:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st148
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st145
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st145
			}
		default:
			goto st145
		}
		goto st106
	st145:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof145
		}
	st_case_145:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st148
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st146
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st146
			}
		default:
			goto st146
		}
		goto st106
	st146:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof146
		}
	st_case_146:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st148
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st147
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st147
			}
		default:
			goto st147
		}
		goto st106
	st147:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr158
		case 58:
			goto st148
		}
		goto st106
	st148:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof148
		}
	st_case_148:
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st149
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st149
			}
		default:
			goto st149
		}
		goto st106
	st149:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof149
		}
	st_case_149:
		if (m.data)[(m.p)] == 32 {
			goto tr158
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st150
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st150
			}
		default:
			goto st150
		}
		goto st106
	st150:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof150
		}
	st_case_150:
		if (m.data)[(m.p)] == 32 {
			goto tr158
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st151
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st151
			}
		default:
			goto st151
		}
		goto st106
	st151:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof151
		}
	st_case_151:
		if (m.data)[(m.p)] == 32 {
			goto tr158
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st152
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st152
			}
		default:
			goto st152
		}
		goto st106
	st152:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof152
		}
	st_case_152:
		if (m.data)[(m.p)] == 32 {
			goto tr158
		}
		goto st106
	tr56:
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st33
	tr106:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st33
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

//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st33
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
//line machine.go:2317
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st34
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st34
			}
		default:
			goto st34
		}
		goto st29
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st35
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st35
			}
		default:
			goto st35
		}
		goto st29
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr62
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st36
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st36
			}
		default:
			goto st36
		}
		goto st29
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr62
		}
		goto st29
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
//line machine.go:2584
		switch (m.data)[(m.p)] {
		case 112:
			goto st50
		case 117:
			goto st78
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
			goto st77
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st76
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
			goto st70
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
			goto tr103
		case 58:
			goto tr103
		}
		goto st0
	tr103:
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
//line machine.go:2799
		switch (m.data)[(m.p)] {
		case 32:
			goto tr105
		case 58:
			goto tr107
		case 91:
			goto tr108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr106
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr106
			}
		default:
			goto tr106
		}
		goto tr104
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof70
		}
	st_case_70:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st71
		}
		goto st0
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
		if (m.data)[(m.p)] == 58 {
			goto st73
		}
		goto st0
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof73
		}
	st_case_73:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
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
		switch (m.data)[(m.p)] {
		case 32:
			goto tr115
		case 58:
			goto tr117
		case 91:
			goto tr118
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr116
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr116
			}
		default:
			goto tr116
		}
		goto tr114
	tr115:
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

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st153
	st153:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof153
		}
	st_case_153:
//line machine.go:2920
		switch (m.data)[(m.p)] {
		case 32:
			goto tr105
		case 58:
			goto tr151
		case 91:
			goto tr139
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr192
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr150
			}
		default:
			goto tr150
		}
		goto tr149
	tr192:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st154
	st154:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof154
		}
	st_case_154:
//line machine.go:2957
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st155
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st114
			}
		default:
			goto st114
		}
		goto st112
	st155:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof155
		}
	st_case_155:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st156
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st115
			}
		default:
			goto st115
		}
		goto st112
	st156:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof156
		}
	st_case_156:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr60
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st157
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st116
			}
		default:
			goto st116
		}
		goto st112
	st157:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof157
		}
	st_case_157:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr196
		case 58:
			goto tr71
		case 91:
			goto tr141
		}
		goto st112
	tr196:
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
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:103

		output.tagSet = true
		output.tag = string(m.text())

		goto st158
	st158:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof158
		}
	st_case_158:
//line machine.go:3079
		switch (m.data)[(m.p)] {
		case 32:
			goto tr105
		case 58:
			goto tr107
		case 91:
			goto tr139
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr150
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr150
			}
		default:
			goto tr150
		}
		goto tr149
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st54
		}
		goto st0
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof77
		}
	st_case_77:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st54
		}
		goto st0
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof78
		}
	st_case_78:
		if (m.data)[(m.p)] == 103 {
			goto st51
		}
		goto st0
	tr7:
//line machine.go.rl:21

		m.pb = m.p

		goto st79
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

		goto st79
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof79
		}
	st_case_79:
//line machine.go:3154
		if (m.data)[(m.p)] == 101 {
			goto st80
		}
		goto st0
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof80
		}
	st_case_80:
		if (m.data)[(m.p)] == 99 {
			goto st51
		}
		goto st0
	tr8:
//line machine.go.rl:21

		m.pb = m.p

		goto st81
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

		goto st81
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof81
		}
	st_case_81:
//line machine.go:3194
		if (m.data)[(m.p)] == 101 {
			goto st82
		}
		goto st0
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof82
		}
	st_case_82:
		if (m.data)[(m.p)] == 98 {
			goto st51
		}
		goto st0
	tr9:
//line machine.go.rl:21

		m.pb = m.p

		goto st83
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

		goto st83
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof83
		}
	st_case_83:
//line machine.go:3234
		switch (m.data)[(m.p)] {
		case 97:
			goto st84
		case 117:
			goto st85
		}
		goto st0
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof84
		}
	st_case_84:
		if (m.data)[(m.p)] == 110 {
			goto st51
		}
		goto st0
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof85
		}
	st_case_85:
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

		goto st86
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

		goto st86
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof86
		}
	st_case_86:
//line machine.go:3289
		if (m.data)[(m.p)] == 97 {
			goto st87
		}
		goto st0
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof87
		}
	st_case_87:
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

		goto st88
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

		goto st88
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof88
		}
	st_case_88:
//line machine.go:3332
		if (m.data)[(m.p)] == 111 {
			goto st89
		}
		goto st0
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof89
		}
	st_case_89:
		if (m.data)[(m.p)] == 118 {
			goto st51
		}
		goto st0
	tr12:
//line machine.go.rl:21

		m.pb = m.p

		goto st90
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

		goto st90
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof90
		}
	st_case_90:
//line machine.go:3372
		if (m.data)[(m.p)] == 99 {
			goto st91
		}
		goto st0
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof91
		}
	st_case_91:
		if (m.data)[(m.p)] == 116 {
			goto st51
		}
		goto st0
	tr13:
//line machine.go.rl:21

		m.pb = m.p

		goto st92
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

		goto st92
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof92
		}
	st_case_92:
//line machine.go:3412
		if (m.data)[(m.p)] == 101 {
			goto st93
		}
		goto st0
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof93
		}
	st_case_93:
		if (m.data)[(m.p)] == 112 {
			goto st51
		}
		goto st0
	tr3:
//line machine.go.rl:21

		m.pb = m.p

		goto st94
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof94
		}
	st_case_94:
//line machine.go:3437
		if (m.data)[(m.p)] == 58 {
			goto tr128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st95
		}
		goto st0
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof95
		}
	st_case_95:
		if (m.data)[(m.p)] == 58 {
			goto tr128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st96
		}
		goto st0
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof96
		}
	st_case_96:
		if (m.data)[(m.p)] == 58 {
			goto tr128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st97
		}
		goto st0
	st97:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch (m.data)[(m.p)] {
		case 45:
			goto st8
		case 58:
			goto tr128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st98
		}
		goto st0
	st98:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof98
		}
	st_case_98:
		if (m.data)[(m.p)] == 58 {
			goto tr128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st98
		}
		goto st0
	tr4:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:34

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st99
	tr128:
//line machine.go.rl:34

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st99
	st99:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof99
		}
	st_case_99:
//line machine.go:3519
		if (m.data)[(m.p)] == 32 {
			goto st2
		}
		goto st0
	st100:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof100
		}
	st_case_100:
		if (m.data)[(m.p)] == 62 {
			goto tr133
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr132
		}
		goto st0
	tr132:
//line machine.go.rl:21

		m.pb = m.p

		goto st101
	st101:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof101
		}
	st_case_101:
//line machine.go:3547
		if (m.data)[(m.p)] == 62 {
			goto tr135
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st101
		}
		goto st0
	tr133:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:29

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st102
	tr135:
//line machine.go.rl:29

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st102
	st102:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof102
		}
	st_case_102:
//line machine.go:3578
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
	_test_eof109:
		m.cs = 109
		goto _test_eof
	_test_eof110:
		m.cs = 110
		goto _test_eof
	_test_eof30:
		m.cs = 30
		goto _test_eof
	_test_eof31:
		m.cs = 31
		goto _test_eof
	_test_eof32:
		m.cs = 32
		goto _test_eof
	_test_eof111:
		m.cs = 111
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
	_test_eof122:
		m.cs = 122
		goto _test_eof
	_test_eof123:
		m.cs = 123
		goto _test_eof
	_test_eof124:
		m.cs = 124
		goto _test_eof
	_test_eof125:
		m.cs = 125
		goto _test_eof
	_test_eof126:
		m.cs = 126
		goto _test_eof
	_test_eof127:
		m.cs = 127
		goto _test_eof
	_test_eof128:
		m.cs = 128
		goto _test_eof
	_test_eof129:
		m.cs = 129
		goto _test_eof
	_test_eof130:
		m.cs = 130
		goto _test_eof
	_test_eof131:
		m.cs = 131
		goto _test_eof
	_test_eof132:
		m.cs = 132
		goto _test_eof
	_test_eof133:
		m.cs = 133
		goto _test_eof
	_test_eof134:
		m.cs = 134
		goto _test_eof
	_test_eof135:
		m.cs = 135
		goto _test_eof
	_test_eof136:
		m.cs = 136
		goto _test_eof
	_test_eof137:
		m.cs = 137
		goto _test_eof
	_test_eof138:
		m.cs = 138
		goto _test_eof
	_test_eof139:
		m.cs = 139
		goto _test_eof
	_test_eof140:
		m.cs = 140
		goto _test_eof
	_test_eof141:
		m.cs = 141
		goto _test_eof
	_test_eof142:
		m.cs = 142
		goto _test_eof
	_test_eof143:
		m.cs = 143
		goto _test_eof
	_test_eof144:
		m.cs = 144
		goto _test_eof
	_test_eof145:
		m.cs = 145
		goto _test_eof
	_test_eof146:
		m.cs = 146
		goto _test_eof
	_test_eof147:
		m.cs = 147
		goto _test_eof
	_test_eof148:
		m.cs = 148
		goto _test_eof
	_test_eof149:
		m.cs = 149
		goto _test_eof
	_test_eof150:
		m.cs = 150
		goto _test_eof
	_test_eof151:
		m.cs = 151
		goto _test_eof
	_test_eof152:
		m.cs = 152
		goto _test_eof
	_test_eof33:
		m.cs = 33
		goto _test_eof
	_test_eof34:
		m.cs = 34
		goto _test_eof
	_test_eof35:
		m.cs = 35
		goto _test_eof
	_test_eof36:
		m.cs = 36
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
	_test_eof153:
		m.cs = 153
		goto _test_eof
	_test_eof154:
		m.cs = 154
		goto _test_eof
	_test_eof155:
		m.cs = 155
		goto _test_eof
	_test_eof156:
		m.cs = 156
		goto _test_eof
	_test_eof157:
		m.cs = 157
		goto _test_eof
	_test_eof158:
		m.cs = 158
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

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 104, 106, 108, 109, 110, 112, 113, 114, 115, 116, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 154, 155, 156, 157:
//line machine.go.rl:113

				output.messageSet = true
				output.message = string(m.text())

			case 103, 105, 107, 111, 117, 153, 158:
//line machine.go.rl:21

				m.pb = m.p

//line machine.go.rl:113

				output.messageSet = true
				output.message = string(m.text())

//line machine.go:3787
			}
		}

	_out:
		{
		}
	}

//line machine.go.rl:277

	return output.export(), m.err
}
