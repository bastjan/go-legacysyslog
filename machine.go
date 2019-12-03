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

//line machine.go.rl:207

//line machine.go:22
const start int = 1
const first_final int = 71

const en_main int = 1

//line machine.go.rl:210

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

//line machine.go.rl:225

//line machine.go.rl:226

//line machine.go.rl:227

//line machine.go.rl:228

//line machine.go.rl:229

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

//line machine.go.rl:255

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
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
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
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
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
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
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
			goto st68
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
		case 84:
			goto st63
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
		case 84:
			goto st63
		}
		goto st0
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch (m.data)[(m.p)] {
		case 65:
			goto tr15
		case 68:
			goto tr16
		case 70:
			goto tr17
		case 74:
			goto tr18
		case 77:
			goto tr19
		case 78:
			goto tr20
		case 79:
			goto tr21
		case 83:
			goto tr22
		case 84:
			goto tr23
		}
		goto st0
	tr6:
//line machine.go.rl:20

		m.pb = m.p

		goto st4
	tr15:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:390
		switch (m.data)[(m.p)] {
		case 112:
			goto st5
		case 117:
			goto st47
		}
		goto st0
	st5:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof5
		}
	st_case_5:
		if (m.data)[(m.p)] == 114 {
			goto st6
		}
		goto st0
	st6:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof6
		}
	st_case_6:
		if (m.data)[(m.p)] == 32 {
			goto st7
		}
		goto st0
	st7:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch (m.data)[(m.p)] {
		case 32:
			goto st8
		case 48:
			goto st8
		case 51:
			goto st46
		}
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 50 {
			goto st45
		}
		goto st0
	st8:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 49 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof9
		}
	st_case_9:
		if (m.data)[(m.p)] == 32 {
			goto st10
		}
		goto st0
	st10:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof12
		}
	st_case_12:
		if (m.data)[(m.p)] == 58 {
			goto st34
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st14
		}
		goto st0
	st14:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof14
		}
	st_case_14:
		if (m.data)[(m.p)] == 32 {
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
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st17
		}
		goto st0
	st17:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof17
		}
	st_case_17:
		if (m.data)[(m.p)] == 58 {
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
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st20
		}
		goto st0
	st20:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof20
		}
	st_case_20:
		if (m.data)[(m.p)] == 58 {
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
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st23
		}
		goto st0
	st23:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr47
		case 58:
			goto tr47
		}
		goto st0
	tr47:
//line machine.go.rl:59

		{
			output.timestampSet = true
			t, err := time.Parse("Jan _2 2006 15:04:05", string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestamp = t
		}

//line machine.go.rl:136
		(m.p)--

		goto st24
	st24:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:605
		switch (m.data)[(m.p)] {
		case 32:
			goto tr49
		case 58:
			goto tr50
		case 91:
			goto st0
		}
		goto tr48
	tr48:
//line machine.go.rl:20

		m.pb = m.p

		goto st25
	tr78:
//line machine.go.rl:47

		{
			output.timestampSet = true
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st25
	st25:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof25
		}
	st_case_25:
//line machine.go:644
		switch (m.data)[(m.p)] {
		case 32:
			goto tr52
		case 91:
			goto st0
		}
		goto st25
	tr52:
//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st26
	st26:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof26
		}
	st_case_26:
//line machine.go:664
		switch (m.data)[(m.p)] {
		case 32:
			goto tr54
		case 58:
			goto tr55
		case 91:
			goto tr56
		}
		goto tr53
	tr53:
//line machine.go.rl:20

		m.pb = m.p

		goto st27
	st27:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof27
		}
	st_case_27:
//line machine.go:685
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr58
		case 91:
			goto tr59
		}
		goto st27
	tr55:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st71
	tr58:
//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st71
	tr61:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st71
	tr64:
//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st71
	st71:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof71
		}
	st_case_71:
//line machine.go:736
		if (m.data)[(m.p)] == 32 {
			goto tr104
		}
		goto tr103
	tr103:
//line machine.go.rl:20

		m.pb = m.p

		goto st72
	st72:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof72
		}
	st_case_72:
//line machine.go:752
		goto st72
	tr104:
//line machine.go.rl:20

		m.pb = m.p

		goto st73
	st73:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof73
		}
	st_case_73:
//line machine.go:765
		goto tr103
	tr56:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st28
	tr59:
//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st28
	st28:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof28
		}
	st_case_28:
//line machine.go:790
		switch (m.data)[(m.p)] {
		case 32:
			goto tr61
		case 58:
			goto tr61
		case 93:
			goto tr62
		}
		goto tr60
	tr60:
//line machine.go.rl:20

		m.pb = m.p

		goto st29
	st29:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof29
		}
	st_case_29:
//line machine.go:811
		switch (m.data)[(m.p)] {
		case 32:
			goto tr64
		case 58:
			goto tr64
		case 93:
			goto tr65
		}
		goto st29
	tr62:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st30
	tr65:
//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st30
	st30:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof30
		}
	st_case_30:
//line machine.go:844
		switch (m.data)[(m.p)] {
		case 32:
			goto st71
		case 58:
			goto st71
		}
		goto st0
	tr54:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st74
	tr71:
//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st74
	tr115:
//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:20

		m.pb = m.p

		goto st74
	st74:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof74
		}
	st_case_74:
//line machine.go:891
		switch (m.data)[(m.p)] {
		case 32:
			goto tr54
		case 58:
			goto tr55
		case 91:
			goto tr107
		}
		goto tr106
	tr106:
//line machine.go.rl:20

		m.pb = m.p

		goto st75
	st75:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof75
		}
	st_case_75:
//line machine.go:912
		switch (m.data)[(m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr58
		case 91:
			goto tr109
		}
		goto st75
	tr107:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st76
	tr109:
//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st76
	st76:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof76
		}
	st_case_76:
//line machine.go:945
		switch (m.data)[(m.p)] {
		case 32:
			goto tr61
		case 58:
			goto tr61
		case 93:
			goto tr111
		}
		goto tr110
	tr110:
//line machine.go.rl:20

		m.pb = m.p

		goto st77
	st77:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof77
		}
	st_case_77:
//line machine.go:966
		switch (m.data)[(m.p)] {
		case 32:
			goto tr64
		case 58:
			goto tr64
		case 93:
			goto tr113
		}
		goto st77
	tr111:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st78
	tr113:
//line machine.go.rl:95

		output.contentSet = true
		output.content = string(m.text())

		goto st78
	st78:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof78
		}
	st_case_78:
//line machine.go:999
		switch (m.data)[(m.p)] {
		case 32:
			goto st71
		case 58:
			goto st71
		}
		goto st72
	tr49:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st31
	st31:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line machine.go:1023
		switch (m.data)[(m.p)] {
		case 32:
			goto tr68
		case 58:
			goto tr69
		case 91:
			goto tr56
		}
		goto tr67
	tr67:
//line machine.go.rl:20

		m.pb = m.p

		goto st32
	st32:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:1044
		switch (m.data)[(m.p)] {
		case 32:
			goto tr71
		case 58:
			goto tr72
		case 91:
			goto tr59
		}
		goto st32
	tr69:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st79
	tr72:
//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st79
	st79:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof79
		}
	st_case_79:
//line machine.go:1077
		switch (m.data)[(m.p)] {
		case 32:
			goto tr115
		case 91:
			goto tr103
		}
		goto tr114
	tr114:
//line machine.go.rl:20

		m.pb = m.p

		goto st80
	st80:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof80
		}
	st_case_80:
//line machine.go:1096
		switch (m.data)[(m.p)] {
		case 32:
			goto tr117
		case 91:
			goto st72
		}
		goto st80
	tr117:
//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st81
	st81:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof81
		}
	st_case_81:
//line machine.go:1116
		switch (m.data)[(m.p)] {
		case 32:
			goto tr54
		case 58:
			goto tr55
		case 91:
			goto tr107
		}
		goto tr106
	tr121:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st82
	tr68:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st82
	st82:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof82
		}
	st_case_82:
//line machine.go:1158
		switch (m.data)[(m.p)] {
		case 32:
			goto tr68
		case 58:
			goto tr69
		case 91:
			goto tr107
		}
		goto tr118
	tr118:
//line machine.go.rl:20

		m.pb = m.p

		goto st83
	st83:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof83
		}
	st_case_83:
//line machine.go:1179
		switch (m.data)[(m.p)] {
		case 32:
			goto tr71
		case 58:
			goto tr72
		case 91:
			goto tr109
		}
		goto st83
	tr50:
//line machine.go.rl:20

		m.pb = m.p

		goto st33
	tr80:
//line machine.go.rl:47

		{
			output.timestampSet = true
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st33
	st33:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof33
		}
	st_case_33:
//line machine.go:1218
		switch (m.data)[(m.p)] {
		case 32:
			goto tr49
		case 91:
			goto st0
		}
		goto tr48
	st34:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof34
		}
	st_case_34:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st35
		}
		goto st0
	st35:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof35
		}
	st_case_35:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st36
		}
		goto st0
	st36:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof36
		}
	st_case_36:
		if (m.data)[(m.p)] == 58 {
			goto st37
		}
		goto st0
	st37:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof37
		}
	st_case_37:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st38
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
		case 32:
			goto tr79
		case 58:
			goto tr80
		case 91:
			goto st0
		}
		goto tr78
	tr79:
//line machine.go.rl:47

		{
			output.timestampSet = true
			t, err := time.Parse(time.Stamp, string(m.text()))
			if err != nil {
				m.err = err
				return output.export(), m.err
			}
			output.timestamp = t
		}

//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

		goto st40
	st40:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof40
		}
	st_case_40:
//line machine.go:1313
		switch (m.data)[(m.p)] {
		case 32:
			goto tr68
		case 58:
			goto tr69
		case 91:
			goto tr56
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr81
		}
		goto tr67
	tr81:
//line machine.go.rl:20

		m.pb = m.p

		goto st41
	st41:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof41
		}
	st_case_41:
//line machine.go:1337
		switch (m.data)[(m.p)] {
		case 32:
			goto tr71
		case 58:
			goto tr72
		case 91:
			goto tr59
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st42
		}
		goto st32
	st42:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr71
		case 58:
			goto tr72
		case 91:
			goto tr59
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st43
		}
		goto st32
	st43:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr71
		case 58:
			goto tr72
		case 91:
			goto tr59
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st44
		}
		goto st32
	st44:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch (m.data)[(m.p)] {
		case 32:
			goto tr85
		case 58:
			goto tr72
		case 91:
			goto tr59
		}
		goto st32
	tr85:
//line machine.go.rl:71

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

//line machine.go.rl:147
		(m.p)--

//line machine.go.rl:85

		output.hostnameSet = true
		output.hostname = string(m.text())

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st84
	st84:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof84
		}
	st_case_84:
//line machine.go:1432
		switch (m.data)[(m.p)] {
		case 32:
			goto tr68
		case 58:
			goto tr120
		case 91:
			goto tr107
		}
		goto tr118
	tr120:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:90

		output.tagSet = true
		output.tag = string(m.text())

		goto st85
	st85:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof85
		}
	st_case_85:
//line machine.go:1458
		switch (m.data)[(m.p)] {
		case 32:
			goto tr121
		case 91:
			goto tr103
		}
		goto tr114
	st45:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof45
		}
	st_case_45:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st9
		}
		goto st0
	st46:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof46
		}
	st_case_46:
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 49 {
			goto st9
		}
		goto st0
	st47:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof47
		}
	st_case_47:
		if (m.data)[(m.p)] == 103 {
			goto st6
		}
		goto st0
	tr7:
//line machine.go.rl:20

		m.pb = m.p

		goto st48
	tr16:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st48
	st48:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:1519
		if (m.data)[(m.p)] == 101 {
			goto st49
		}
		goto st0
	st49:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof49
		}
	st_case_49:
		if (m.data)[(m.p)] == 99 {
			goto st6
		}
		goto st0
	tr8:
//line machine.go.rl:20

		m.pb = m.p

		goto st50
	tr17:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st50
	st50:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:1559
		if (m.data)[(m.p)] == 101 {
			goto st51
		}
		goto st0
	st51:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof51
		}
	st_case_51:
		if (m.data)[(m.p)] == 98 {
			goto st6
		}
		goto st0
	tr9:
//line machine.go.rl:20

		m.pb = m.p

		goto st52
	tr18:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st52
	st52:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof52
		}
	st_case_52:
//line machine.go:1599
		switch (m.data)[(m.p)] {
		case 97:
			goto st53
		case 117:
			goto st54
		}
		goto st0
	st53:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof53
		}
	st_case_53:
		if (m.data)[(m.p)] == 110 {
			goto st6
		}
		goto st0
	st54:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch (m.data)[(m.p)] {
		case 108:
			goto st6
		case 110:
			goto st6
		}
		goto st0
	tr10:
//line machine.go.rl:20

		m.pb = m.p

		goto st55
	tr19:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st55
	st55:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof55
		}
	st_case_55:
//line machine.go:1654
		if (m.data)[(m.p)] == 97 {
			goto st56
		}
		goto st0
	st56:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch (m.data)[(m.p)] {
		case 114:
			goto st6
		case 121:
			goto st6
		}
		goto st0
	tr11:
//line machine.go.rl:20

		m.pb = m.p

		goto st57
	tr20:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st57
	st57:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:1697
		if (m.data)[(m.p)] == 111 {
			goto st58
		}
		goto st0
	st58:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof58
		}
	st_case_58:
		if (m.data)[(m.p)] == 118 {
			goto st6
		}
		goto st0
	tr12:
//line machine.go.rl:20

		m.pb = m.p

		goto st59
	tr21:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st59
	st59:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:1737
		if (m.data)[(m.p)] == 99 {
			goto st60
		}
		goto st0
	st60:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof60
		}
	st_case_60:
		if (m.data)[(m.p)] == 116 {
			goto st6
		}
		goto st0
	tr13:
//line machine.go.rl:20

		m.pb = m.p

		goto st61
	tr22:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

//line machine.go.rl:20

		m.pb = m.p

		goto st61
	st61:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof61
		}
	st_case_61:
//line machine.go:1777
		if (m.data)[(m.p)] == 101 {
			goto st62
		}
		goto st0
	st62:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof62
		}
	st_case_62:
		if (m.data)[(m.p)] == 112 {
			goto st6
		}
		goto st0
	tr23:
//line machine.go.rl:38

		switch m.data[m.p-1] {
		case '.':
			output.ciscoTimestampExt = CiscoTimeClockModeSynced
		case '*':
			output.ciscoTimestampExt = CiscoTimeClockModeUnsynced
		}

		goto st63
	st63:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof63
		}
	st_case_63:
//line machine.go:1807
		if (m.data)[(m.p)] == 79 {
			goto st64
		}
		goto st0
	st64:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof64
		}
	st_case_64:
		if (m.data)[(m.p)] == 68 {
			goto st65
		}
		goto st0
	st65:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof65
		}
	st_case_65:
		if (m.data)[(m.p)] == 79 {
			goto st33
		}
		goto st0
	tr3:
//line machine.go.rl:20

		m.pb = m.p

		goto st66
	st66:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof66
		}
	st_case_66:
//line machine.go:1841
		if (m.data)[(m.p)] == 58 {
			goto tr98
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st66
		}
		goto st0
	tr4:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:33

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st67
	tr98:
//line machine.go.rl:33

		output.ciscoSequenceIDSet = true
		output.ciscoSequenceID = string(m.text())

		goto st67
	st67:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof67
		}
	st_case_67:
//line machine.go:1872
		if (m.data)[(m.p)] == 32 {
			goto st2
		}
		goto st0
	st68:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof68
		}
	st_case_68:
		if (m.data)[(m.p)] == 62 {
			goto tr100
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto tr99
		}
		goto st0
	tr99:
//line machine.go.rl:20

		m.pb = m.p

		goto st69
	st69:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof69
		}
	st_case_69:
//line machine.go:1900
		if (m.data)[(m.p)] == 62 {
			goto tr102
		}
		if 48 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 57 {
			goto st69
		}
		goto st0
	tr100:
//line machine.go.rl:20

		m.pb = m.p

//line machine.go.rl:28

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st70
	tr102:
//line machine.go.rl:28

		output.priority = uint8(common.UnsafeUTF8DecimalCodePointsToInt(m.text()))
		output.prioritySet = true

		goto st70
	st70:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof70
		}
	st_case_70:
//line machine.go:1931
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
		case 84:
			goto st63
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
	_test_eof71:
		m.cs = 71
		goto _test_eof
	_test_eof72:
		m.cs = 72
		goto _test_eof
	_test_eof73:
		m.cs = 73
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
	_test_eof31:
		m.cs = 31
		goto _test_eof
	_test_eof32:
		m.cs = 32
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
	_test_eof84:
		m.cs = 84
		goto _test_eof
	_test_eof85:
		m.cs = 85
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

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 72, 75, 76, 77, 78, 80, 81, 83:
//line machine.go.rl:100

				output.messageSet = true
				output.message = string(m.text())

			case 71, 73, 74, 79, 82, 84, 85:
//line machine.go.rl:20

				m.pb = m.p

//line machine.go.rl:100

				output.messageSet = true
				output.message = string(m.text())

//line machine.go:2069
			}
		}

	_out:
		{
		}
	}

//line machine.go.rl:256

	return output.export(), m.err
}
