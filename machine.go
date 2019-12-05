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

//line machine.go.rl:227

//line machine.go:22
const start int = 1
const first_final int = 51

const en_main int = 1

//line machine.go.rl:230

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	// general mark
	pb int
	// hostname machine can run in parallel to the tag machine since it is not clear from the beginning
	// if it is a hostname or tag
	pHostname int
	// same for the timestamp
	pTimestamp int

	err        error
	bestEffort bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine() *machine {
	m := &machine{}

//line machine.go.rl:252

//line machine.go.rl:253

//line machine.go.rl:254

//line machine.go.rl:255

//line machine.go.rl:256

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
	m.pTimestamp = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

	output := &syslogMessage{}

//line machine.go:89
	{
		m.cs = start
	}

//line machine.go.rl:284

//line machine.go:96
	{
		if (m.p) == (m.pe) {
			goto _test_eof
		}
		switch m.cs {
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 0:
			goto st_case_0
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
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
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
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
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
		case 32:
			goto st_case_32
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
		case 192:
			goto st_case_192
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		}
		goto st_out
	st_case_1:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr2
		case 46:
			goto tr2
		case 58:
			goto tr4
		case 60:
			goto tr5
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
		case 91:
			goto tr14
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto tr0
	tr0:
//line machine.go.rl:21

		m.pb = m.p

		goto st2
	tr26:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
//line machine.go:551
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	tr82:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	tr16:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	tr19:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st51
	tr22:
//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st51
	tr27:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st51
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof51
		}
	st_case_51:
//line machine.go:622
		if (m.data)[(m.p)] == 32 {
			goto tr77
		}
		goto tr76
	tr76:
//line machine.go.rl:21

		m.pb = m.p

		goto st52
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof52
		}
	st_case_52:
//line machine.go:638
		goto st52
	tr77:
//line machine.go.rl:21

		m.pb = m.p

		goto st53
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof53
		}
	st_case_53:
//line machine.go:651
		goto tr76
	tr14:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	tr17:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	tr37:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st3
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line machine.go:696
		switch (m.data)[(m.p)] {
		case 32:
			goto tr19
		case 58:
			goto tr19
		case 93:
			goto tr20
		}
		goto tr18
	tr18:
//line machine.go.rl:21

		m.pb = m.p

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:717
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 58:
			goto tr22
		case 93:
			goto tr23
		}
		goto st4
	tr20:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st5
	tr23:
//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st5
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
//line machine.go:750
		switch (m.data)[(m.p)] {
		case 32:
			goto st51
		case 58:
			goto st51
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr245:
//line machine.go.rl:21

		m.pb = m.p

		goto st54
	tr1:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st54
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof54
		}
	st_case_54:
//line machine.go:784
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr80
		case 46:
			goto tr80
		case 58:
			goto tr82
		case 65:
			goto tr83
		case 68:
			goto tr84
		case 70:
			goto tr85
		case 74:
			goto tr86
		case 77:
			goto tr87
		case 78:
			goto tr88
		case 79:
			goto tr89
		case 83:
			goto tr90
		case 91:
			goto tr91
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr81
		}
		goto tr79
	tr79:
//line machine.go.rl:21

		m.pb = m.p

		goto st55
	tr98:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

		goto st55
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof55
		}
	st_case_55:
//line machine.go:843
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	tr91:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr93:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr108:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr142:
//line machine.go.rl:52

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	tr230:
//line machine.go.rl:64

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st56
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof56
		}
	st_case_56:
//line machine.go:942
		switch (m.data)[(m.p)] {
		case 32:
			goto tr19
		case 58:
			goto tr19
		case 93:
			goto tr95
		}
		goto tr94
	tr94:
//line machine.go.rl:21

		m.pb = m.p

		goto st57
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:963
		switch (m.data)[(m.p)] {
		case 32:
			goto tr22
		case 58:
			goto tr22
		case 93:
			goto tr97
		}
		goto st57
	tr95:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st58
	tr97:
//line machine.go.rl:112

		output.contentSet = true
		output.content = string(m.text())

		goto st58
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof58
		}
	st_case_58:
//line machine.go:996
		switch (m.data)[(m.p)] {
		case 32:
			goto st51
		case 58:
			goto st51
		}
		goto st52
	tr80:
//line machine.go.rl:21

		m.pb = m.p

		goto st59
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:1015
		switch (m.data)[(m.p)] {
		case 32:
			goto tr27
		case 58:
			goto tr27
		case 65:
			goto tr100
		case 68:
			goto tr101
		case 70:
			goto tr102
		case 74:
			goto tr103
		case 77:
			goto tr104
		case 78:
			goto tr105
		case 79:
			goto tr106
		case 83:
			goto tr107
		case 91:
			goto tr108
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr99
		}
		goto tr98
	tr81:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st60
	tr99:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st60
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof60
		}
	st_case_60:
//line machine.go:1078
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st61
		}
		goto st55
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st62
		}
		goto st55
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st63
		}
		goto st55
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st64
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st65
		case 49:
			goto st139
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st66
		}
		goto st55
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st67
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st68
		case 51:
			goto st138
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st137
		}
		goto st55
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st69
		}
		goto st55
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st70
		case 91:
			goto tr93
		}
		goto st55
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 50:
			goto st136
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st71
		}
		goto st55
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st72
		}
		goto st55
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr54
		case 91:
			goto tr93
		}
		goto st55
	tr54:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st73
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof73
		}
	st_case_73:
//line machine.go:1308
		if (m.data)[(m.p)] == 32 {
			goto tr77
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto tr125
		}
		goto tr76
	tr125:
//line machine.go.rl:21

		m.pb = m.p

		goto st74
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof74
		}
	st_case_74:
//line machine.go:1327
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st75
		}
		goto st52
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof75
		}
	st_case_75:
		if (m.data)[(m.p)] == 58 {
			goto st76
		}
		goto st52
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st77
		}
		goto st52
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof77
		}
	st_case_77:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st78
		}
		goto st52
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 46:
			goto st129
		case 90:
			goto st84
		}
		goto st52
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof79
		}
	st_case_79:
		if (m.data)[(m.p)] == 50 {
			goto st128
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st80
		}
		goto st52
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof80
		}
	st_case_80:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st81
		}
		goto st52
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof81
		}
	st_case_81:
		if (m.data)[(m.p)] == 58 {
			goto st82
		}
		goto st52
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof82
		}
	st_case_82:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 53 {
			goto st83
		}
		goto st52
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof83
		}
	st_case_83:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st84
		}
		goto st52
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr139
		case 58:
			goto tr141
		case 91:
			goto tr142
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr140
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr140
			}
		default:
			goto tr140
		}
		goto tr138
	tr138:
//line machine.go.rl:52

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
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

		goto st85
	tr146:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st85
	tr226:
//line machine.go.rl:64

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
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

		goto st85
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof85
		}
	st_case_85:
//line machine.go:1508
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st85
	tr145:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st86
	tr144:
//line machine.go.rl:102

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st86
	st86:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof86
		}
	st_case_86:
//line machine.go:1546
		switch (m.data)[(m.p)] {
		case 32:
			goto tr145
		case 58:
			goto tr82
		case 91:
			goto tr91
		}
		goto tr79
	tr147:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st87
	tr139:
//line machine.go.rl:52

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st87
	tr220:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st87
	tr229:
//line machine.go.rl:64

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
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

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st87
	st87:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof87
		}
	st_case_87:
//line machine.go:1637
		switch (m.data)[(m.p)] {
		case 32:
			goto tr147
		case 58:
			goto tr149
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr148
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto tr146
	tr140:
//line machine.go.rl:52

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
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

		goto st88
	tr148:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st88
	tr228:
//line machine.go.rl:64

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
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

		goto st88
	st88:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof88
		}
	st_case_88:
//line machine.go:1718
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st89
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st89
			}
		default:
			goto st89
		}
		goto st85
	st89:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st90
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st90
			}
		default:
			goto st90
		}
		goto st85
	st90:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st91
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st91
			}
		default:
			goto st91
		}
		goto st85
	st91:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		goto st85
	tr151:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st92
	tr141:
//line machine.go.rl:52

		{
			t, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", string(m.data[m.pTimestamp:m.p]))
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

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st92
	tr149:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st92
	st92:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof92
		}
	st_case_92:
//line machine.go:1860
		switch (m.data)[(m.p)] {
		case 32:
			goto tr77
		case 58:
			goto tr155
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr154
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr154
			}
		default:
			goto tr154
		}
		goto tr76
	tr154:
//line machine.go.rl:21

		m.pb = m.p

		goto st93
	st93:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof93
		}
	st_case_93:
//line machine.go:1891
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st98
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st95
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st95
			}
		default:
			goto st95
		}
		goto st52
	tr156:
//line machine.go.rl:102

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

		goto st94
	st94:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof94
		}
	st_case_94:
//line machine.go:1923
		switch (m.data)[(m.p)] {
		case 32:
			goto tr145
		case 58:
			goto tr82
		case 91:
			goto tr91
		}
		goto tr79
	st95:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st98
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st96
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto st52
	st96:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st98
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st97
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto st52
	st97:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st98
		}
		goto st52
	tr155:
//line machine.go.rl:21

		m.pb = m.p

		goto st98
	st98:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof98
		}
	st_case_98:
//line machine.go:2004
		if (m.data)[(m.p)] == 58 {
			goto st103
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st99
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st99
			}
		default:
			goto st99
		}
		goto st52
	st99:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st103
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st100
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st100
			}
		default:
			goto st100
		}
		goto st52
	st100:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st103
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st101
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st101
			}
		default:
			goto st101
		}
		goto st52
	st101:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof101
		}
	st_case_101:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st103
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st102
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st102
			}
		default:
			goto st102
		}
		goto st52
	st102:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st103
		}
		goto st52
	st103:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof103
		}
	st_case_103:
		if (m.data)[(m.p)] == 58 {
			goto st108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st104
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st104
			}
		default:
			goto st104
		}
		goto st52
	st104:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st105
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st105
			}
		default:
			goto st105
		}
		goto st52
	st105:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st106
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st106
			}
		default:
			goto st106
		}
		goto st52
	st106:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st108
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st107
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st107
			}
		default:
			goto st107
		}
		goto st52
	st107:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st108
		}
		goto st52
	st108:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof108
		}
	st_case_108:
		if (m.data)[(m.p)] == 58 {
			goto st113
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st109
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st109
			}
		default:
			goto st109
		}
		goto st52
	st109:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st113
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st110
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		goto st52
	st110:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st113
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st111
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st111
			}
		default:
			goto st111
		}
		goto st52
	st111:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st113
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st112
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st112
			}
		default:
			goto st112
		}
		goto st52
	st112:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st113
		}
		goto st52
	st113:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof113
		}
	st_case_113:
		if (m.data)[(m.p)] == 58 {
			goto st118
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
		goto st52
	st114:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st118
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
		goto st52
	st115:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st118
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
		goto st52
	st116:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st118
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st117
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st117
			}
		default:
			goto st117
		}
		goto st52
	st117:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st118
		}
		goto st52
	st118:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof118
		}
	st_case_118:
		if (m.data)[(m.p)] == 58 {
			goto st123
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st119
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st119
			}
		default:
			goto st119
		}
		goto st52
	st119:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof119
		}
	st_case_119:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
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
		goto st52
	st120:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof120
		}
	st_case_120:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
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
		goto st52
	st121:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof121
		}
	st_case_121:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
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
		goto st52
	st122:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof122
		}
	st_case_122:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr156
		case 58:
			goto st123
		}
		goto st52
	st123:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof123
		}
	st_case_123:
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
		goto st52
	st124:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof124
		}
	st_case_124:
		if (m.data)[(m.p)] == 32 {
			goto tr156
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
		goto st52
	st125:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof125
		}
	st_case_125:
		if (m.data)[(m.p)] == 32 {
			goto tr156
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
		goto st52
	st126:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof126
		}
	st_case_126:
		if (m.data)[(m.p)] == 32 {
			goto tr156
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
		goto st52
	st127:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof127
		}
	st_case_127:
		if (m.data)[(m.p)] == 32 {
			goto tr156
		}
		goto st52
	st128:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof128
		}
	st_case_128:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st81
		}
		goto st52
	st129:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof129
		}
	st_case_129:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st130
		}
		goto st52
	st130:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st131
		}
		goto st52
	st131:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st132
		}
		goto st52
	st132:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st133
		}
		goto st52
	st133:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st134
		}
		goto st52
	st134:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st135
		}
		goto st52
	st135:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch (m.data)[(m.p)] {
		case 43:
			goto st79
		case 45:
			goto st79
		case 90:
			goto st84
		}
		goto st52
	st136:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st72
		}
		goto st55
	st137:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st69
		}
		goto st55
	st138:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof138
		}
	st_case_138:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st69
		}
		goto st55
	st139:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof139
		}
	st_case_139:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st66
		}
		goto st55
	tr83:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st140
	tr100:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st140
	st140:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof140
		}
	st_case_140:
//line machine.go:2834
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st141
		case 85:
			goto st176
		case 91:
			goto tr93
		case 112:
			goto st141
		case 117:
			goto st176
		}
		goto st55
	st141:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof141
		}
	st_case_141:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st142
		case 91:
			goto tr93
		case 114:
			goto st142
		}
		goto st55
	st142:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof142
		}
	st_case_142:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr16
		case 91:
			goto tr93
		}
		goto st55
	tr58:
//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st143
	st143:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof143
		}
	st_case_143:
//line machine.go:2896
		switch (m.data)[(m.p)] {
		case 32:
			goto tr199
		case 48:
			goto tr200
		case 51:
			goto tr202
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto tr201
		}
		goto tr76
	tr199:
//line machine.go.rl:21

		m.pb = m.p

		goto st144
	st144:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof144
		}
	st_case_144:
//line machine.go:2920
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr203
		}
		goto tr76
	tr203:
//line machine.go.rl:21

		m.pb = m.p

		goto st145
	st145:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof145
		}
	st_case_145:
//line machine.go:2936
		if (m.data)[(m.p)] == 32 {
			goto st146
		}
		goto st52
	st146:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof146
		}
	st_case_146:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st147
		}
		goto st52
	st147:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof147
		}
	st_case_147:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st148
		}
		goto st52
	st148:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof148
		}
	st_case_148:
		if (m.data)[(m.p)] == 58 {
			goto st161
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st149
		}
		goto st52
	st149:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof149
		}
	st_case_149:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st150
		}
		goto st52
	st150:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof150
		}
	st_case_150:
		if (m.data)[(m.p)] == 32 {
			goto st151
		}
		goto st52
	st151:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof151
		}
	st_case_151:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st152
		}
		goto st52
	st152:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof152
		}
	st_case_152:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st153
		}
		goto st52
	st153:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof153
		}
	st_case_153:
		if (m.data)[(m.p)] == 58 {
			goto st154
		}
		goto st52
	st154:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof154
		}
	st_case_154:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st155
		}
		goto st52
	st155:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof155
		}
	st_case_155:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st156
		}
		goto st52
	st156:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof156
		}
	st_case_156:
		if (m.data)[(m.p)] == 58 {
			goto st157
		}
		goto st52
	st157:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof157
		}
	st_case_157:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st158
		}
		goto st52
	st158:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof158
		}
	st_case_158:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st159
		}
		goto st52
	st159:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof159
		}
	st_case_159:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr219
		case 58:
			goto tr219
		}
		goto st52
	tr219:
//line machine.go.rl:76

		{
			t, err := time.Parse("Jan _2 2006 15:04:05", string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:153
		(m.p)--

		goto st160
	st160:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof160
		}
	st_case_160:
//line machine.go:3095
		switch (m.data)[(m.p)] {
		case 32:
			goto tr147
		case 58:
			goto tr220
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr148
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto tr146
	st161:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof161
		}
	st_case_161:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st162
		}
		goto st52
	st162:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof162
		}
	st_case_162:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st163
		}
		goto st52
	st163:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof163
		}
	st_case_163:
		if (m.data)[(m.p)] == 58 {
			goto st164
		}
		goto st52
	st164:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof164
		}
	st_case_164:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st165
		}
		goto st52
	st165:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof165
		}
	st_case_165:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st166
		}
		goto st52
	st166:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr227
		case 58:
			goto tr229
		case 91:
			goto tr230
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr228
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr228
			}
		default:
			goto tr228
		}
		goto tr226
	tr227:
//line machine.go.rl:64

		{
			t, err := time.Parse(time.Stamp, string(m.data[m.pTimestamp:m.p]))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestampSet = true
			output.timestamp = t
		}

//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st167
	st167:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof167
		}
	st_case_167:
//line machine.go:3216
		switch (m.data)[(m.p)] {
		case 32:
			goto tr147
		case 58:
			goto tr149
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr231
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto tr146
	tr231:
//line machine.go.rl:25

		m.pHostname = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st168
	st168:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof168
		}
	st_case_168:
//line machine.go:3253
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st169
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st89
			}
		default:
			goto st89
		}
		goto st85
	st169:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st170
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st90
			}
		default:
			goto st90
		}
		goto st85
	st170:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr144
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto st171
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto st91
			}
		default:
			goto st91
		}
		goto st85
	st171:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr235
		case 58:
			goto tr151
		case 91:
			goto tr93
		}
		goto st85
	tr235:
//line machine.go.rl:88

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

//line machine.go.rl:164
		(m.p)--

//line machine.go.rl:102

		output.hostnameSet = true
		output.hostname = string(m.data[m.pHostname:m.p])

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st172
	st172:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof172
		}
	st_case_172:
//line machine.go:3375
		switch (m.data)[(m.p)] {
		case 32:
			goto tr147
		case 58:
			goto tr220
		case 91:
			goto tr91
		}
		switch {
		case (m.data)[(m.p)] < 65:
			if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
				goto tr148
			}
		case (m.data)[(m.p)] > 70:
			if 97 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto tr146
	tr200:
//line machine.go.rl:21

		m.pb = m.p

		goto st173
	st173:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof173
		}
	st_case_173:
//line machine.go:3408
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st145
		}
		goto st52
	tr201:
//line machine.go.rl:21

		m.pb = m.p

		goto st174
	st174:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof174
		}
	st_case_174:
//line machine.go:3424
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st145
		}
		goto st52
	tr202:
//line machine.go.rl:21

		m.pb = m.p

		goto st175
	st175:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof175
		}
	st_case_175:
//line machine.go:3440
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st145
		}
		goto st52
	st176:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof176
		}
	st_case_176:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 71:
			goto st142
		case 91:
			goto tr93
		case 103:
			goto st142
		}
		goto st55
	tr84:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st177
	tr101:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st177
	st177:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof177
		}
	st_case_177:
//line machine.go:3497
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st178
		case 91:
			goto tr93
		case 101:
			goto st178
		}
		goto st55
	st178:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof178
		}
	st_case_178:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st142
		case 91:
			goto tr93
		case 99:
			goto st142
		}
		goto st55
	tr85:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st179
	tr102:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st179
	st179:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof179
		}
	st_case_179:
//line machine.go:3563
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st180
		case 91:
			goto tr93
		case 101:
			goto st180
		}
		goto st55
	st180:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof180
		}
	st_case_180:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 66:
			goto st142
		case 91:
			goto tr93
		case 98:
			goto st142
		}
		goto st55
	tr86:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st181
	tr103:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st181
	st181:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof181
		}
	st_case_181:
//line machine.go:3629
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st182
		case 85:
			goto st183
		case 91:
			goto tr93
		case 97:
			goto st182
		case 117:
			goto st183
		}
		goto st55
	st182:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof182
		}
	st_case_182:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 78:
			goto st142
		case 91:
			goto tr93
		case 110:
			goto st142
		}
		goto st55
	st183:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof183
		}
	st_case_183:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 76:
			goto st142
		case 78:
			goto st142
		case 91:
			goto tr93
		case 108:
			goto st142
		case 110:
			goto st142
		}
		goto st55
	tr87:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st184
	tr104:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st184
	st184:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof184
		}
	st_case_184:
//line machine.go:3721
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st185
		case 91:
			goto tr93
		case 97:
			goto st185
		}
		goto st55
	st185:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st142
		case 89:
			goto st142
		case 91:
			goto tr93
		case 114:
			goto st142
		case 121:
			goto st142
		}
		goto st55
	tr88:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st186
	tr105:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st186
	st186:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof186
		}
	st_case_186:
//line machine.go:3791
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 79:
			goto st187
		case 91:
			goto tr93
		case 111:
			goto st187
		}
		goto st55
	st187:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof187
		}
	st_case_187:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 86:
			goto st142
		case 91:
			goto tr93
		case 118:
			goto st142
		}
		goto st55
	tr89:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st188
	tr106:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st188
	st188:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof188
		}
	st_case_188:
//line machine.go:3857
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st189
		case 91:
			goto tr93
		case 99:
			goto st189
		}
		goto st55
	st189:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof189
		}
	st_case_189:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st142
		case 91:
			goto tr93
		case 116:
			goto st142
		}
		goto st55
	tr90:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st190
	tr107:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st190
	st190:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof190
		}
	st_case_190:
//line machine.go:3923
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st191
		case 91:
			goto tr93
		case 101:
			goto st191
		}
		goto st55
	st191:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof191
		}
	st_case_191:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st142
		case 91:
			goto tr93
		case 112:
			goto st142
		}
		goto st55
	tr2:
//line machine.go.rl:21

		m.pb = m.p

		goto st6
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
//line machine.go:3966
		switch (m.data)[(m.p)] {
		case 32:
			goto tr27
		case 58:
			goto tr27
		case 65:
			goto tr29
		case 68:
			goto tr30
		case 70:
			goto tr31
		case 74:
			goto tr32
		case 77:
			goto tr33
		case 78:
			goto tr34
		case 79:
			goto tr35
		case 83:
			goto tr36
		case 91:
			goto tr37
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr28
		}
		goto tr26
	tr28:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st7
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof7
		}
	st_case_7:
//line machine.go:4019
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st8
		}
		goto st2
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st9
		}
		goto st2
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st10
		}
		goto st2
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st11
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st12
		case 49:
			goto st23
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st13
		}
		goto st2
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st14
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 48:
			goto st15
		case 51:
			goto st22
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st21
		}
		goto st2
	st15:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st16
		}
		goto st2
	st16:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st17
		case 91:
			goto tr17
		}
		goto st2
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 50:
			goto st20
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st18
		}
		goto st2
	st18:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st19
		}
		goto st2
	st19:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr54
		case 91:
			goto tr17
		}
		goto st2
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 51 {
			goto st19
		}
		goto st2
	st21:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st16
		}
		goto st2
	st22:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st16
		}
		goto st2
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st13
		}
		goto st2
	tr6:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st24
	tr29:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st24
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:4339
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st25
		case 85:
			goto st27
		case 91:
			goto tr17
		case 112:
			goto st25
		case 117:
			goto st27
		}
		goto st2
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st26
		case 91:
			goto tr17
		case 114:
			goto st26
		}
		goto st2
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr16
		case 91:
			goto tr17
		}
		goto st2
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 71:
			goto st26
		case 91:
			goto tr17
		case 103:
			goto st26
		}
		goto st2
	tr7:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st28
	tr30:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st28
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof28
		}
	st_case_28:
//line machine.go:4441
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st29
		case 91:
			goto tr17
		case 101:
			goto st29
		}
		goto st2
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st26
		case 91:
			goto tr17
		case 99:
			goto st26
		}
		goto st2
	tr8:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st30
	tr31:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st30
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line machine.go:4507
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st31
		case 91:
			goto tr17
		case 101:
			goto st31
		}
		goto st2
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 66:
			goto st26
		case 91:
			goto tr17
		case 98:
			goto st26
		}
		goto st2
	tr9:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st32
	tr32:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st32
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:4573
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st33
		case 85:
			goto st34
		case 91:
			goto tr17
		case 97:
			goto st33
		case 117:
			goto st34
		}
		goto st2
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 78:
			goto st26
		case 91:
			goto tr17
		case 110:
			goto st26
		}
		goto st2
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 76:
			goto st26
		case 78:
			goto st26
		case 91:
			goto tr17
		case 108:
			goto st26
		case 110:
			goto st26
		}
		goto st2
	tr10:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st35
	tr33:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st35
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line machine.go:4665
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 65:
			goto st36
		case 91:
			goto tr17
		case 97:
			goto st36
		}
		goto st2
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 82:
			goto st26
		case 89:
			goto st26
		case 91:
			goto tr17
		case 114:
			goto st26
		case 121:
			goto st26
		}
		goto st2
	tr11:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st37
	tr34:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st37
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof37
		}
	st_case_37:
//line machine.go:4735
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 79:
			goto st38
		case 91:
			goto tr17
		case 111:
			goto st38
		}
		goto st2
	st38:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 86:
			goto st26
		case 91:
			goto tr17
		case 118:
			goto st26
		}
		goto st2
	tr12:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st39
	tr35:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st39
	st39:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof39
		}
	st_case_39:
//line machine.go:4801
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 67:
			goto st40
		case 91:
			goto tr17
		case 99:
			goto st40
		}
		goto st2
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 84:
			goto st26
		case 91:
			goto tr17
		case 116:
			goto st26
		}
		goto st2
	tr13:
//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st41
	tr36:
//line machine.go.rl:43

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:29

		m.pTimestamp = m.p

//line machine.go.rl:21

		m.pb = m.p

		goto st41
	st41:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof41
		}
	st_case_41:
//line machine.go:4867
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 69:
			goto st42
		case 91:
			goto tr17
		case 101:
			goto st42
		}
		goto st2
	st42:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 80:
			goto st26
		case 91:
			goto tr17
		case 112:
			goto st26
		}
		goto st2
	tr3:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:29

		m.pTimestamp = m.p

		goto st43
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof43
		}
	st_case_43:
//line machine.go:4914
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st44
		}
		goto st2
	st44:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st45
		}
		goto st2
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st46
		}
		goto st2
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 45:
			goto st11
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st47
		}
		goto st2
	st47:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr68
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st47
		}
		goto st2
	tr4:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:38

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st192
	tr68:
//line machine.go.rl:38

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

//line machine.go.rl:107

		output.tagSet = true
		output.tag = string(m.text())

		goto st192
	st192:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof192
		}
	st_case_192:
//line machine.go:5030
		if (m.data)[(m.p)] == 32 {
			goto tr245
		}
		goto tr76
	tr5:
//line machine.go.rl:21

		m.pb = m.p

		goto st48
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:5046
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 62:
			goto tr73
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr72
		}
		goto st2
	tr72:
//line machine.go.rl:21

		m.pb = m.p

		goto st49
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof49
		}
	st_case_49:
//line machine.go:5072
		switch (m.data)[(m.p)] {
		case 32:
			goto tr16
		case 58:
			goto tr16
		case 62:
			goto tr75
		case 91:
			goto tr17
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st49
		}
		goto st2
	tr73:
//line machine.go.rl:21

		m.pb = m.p

//line machine.go.rl:33

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st50
	tr75:
//line machine.go.rl:33

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st50
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:5110
		switch (m.data)[(m.p)] {
		case 32:
			goto tr1
		case 42:
			goto tr2
		case 46:
			goto tr2
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
		case 91:
			goto tr14
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr3
		}
		goto tr0
	st_out:
	_test_eof2:
		m.cs = 2
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
	_test_eof3:
		m.cs = 3
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof5:
		m.cs = 5
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
	_test_eof109:
		m.cs = 109
		goto _test_eof
	_test_eof110:
		m.cs = 110
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
	_test_eof159:
		m.cs = 159
		goto _test_eof
	_test_eof160:
		m.cs = 160
		goto _test_eof
	_test_eof161:
		m.cs = 161
		goto _test_eof
	_test_eof162:
		m.cs = 162
		goto _test_eof
	_test_eof163:
		m.cs = 163
		goto _test_eof
	_test_eof164:
		m.cs = 164
		goto _test_eof
	_test_eof165:
		m.cs = 165
		goto _test_eof
	_test_eof166:
		m.cs = 166
		goto _test_eof
	_test_eof167:
		m.cs = 167
		goto _test_eof
	_test_eof168:
		m.cs = 168
		goto _test_eof
	_test_eof169:
		m.cs = 169
		goto _test_eof
	_test_eof170:
		m.cs = 170
		goto _test_eof
	_test_eof171:
		m.cs = 171
		goto _test_eof
	_test_eof172:
		m.cs = 172
		goto _test_eof
	_test_eof173:
		m.cs = 173
		goto _test_eof
	_test_eof174:
		m.cs = 174
		goto _test_eof
	_test_eof175:
		m.cs = 175
		goto _test_eof
	_test_eof176:
		m.cs = 176
		goto _test_eof
	_test_eof177:
		m.cs = 177
		goto _test_eof
	_test_eof178:
		m.cs = 178
		goto _test_eof
	_test_eof179:
		m.cs = 179
		goto _test_eof
	_test_eof180:
		m.cs = 180
		goto _test_eof
	_test_eof181:
		m.cs = 181
		goto _test_eof
	_test_eof182:
		m.cs = 182
		goto _test_eof
	_test_eof183:
		m.cs = 183
		goto _test_eof
	_test_eof184:
		m.cs = 184
		goto _test_eof
	_test_eof185:
		m.cs = 185
		goto _test_eof
	_test_eof186:
		m.cs = 186
		goto _test_eof
	_test_eof187:
		m.cs = 187
		goto _test_eof
	_test_eof188:
		m.cs = 188
		goto _test_eof
	_test_eof189:
		m.cs = 189
		goto _test_eof
	_test_eof190:
		m.cs = 190
		goto _test_eof
	_test_eof191:
		m.cs = 191
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
	_test_eof32:
		m.cs = 32
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
	_test_eof192:
		m.cs = 192
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

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 52, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 88, 89, 90, 91, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 168, 169, 170, 171, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191:
//line machine.go.rl:117

				output.messageSet = true
				output.message = string(m.text())

			case 51, 53, 54, 73, 86, 87, 92, 143, 144, 167, 172, 192:
//line machine.go.rl:21

				m.pb = m.p

//line machine.go.rl:117

				output.messageSet = true
				output.message = string(m.text())

//line machine.go:5355
			}
		}

	_out:
		{
		}
	}

//line machine.go.rl:285

	return output.export(), m.err
}
