
//line machine.go.rl:1
package legacysyslog

import (
	"fmt"
)


//line machine.go.rl:89



//line machine.go:15
const start int = 1
const first_final int = 58

const en_main int = 1


//line machine.go.rl:92

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

	
//line machine.go.rl:107
	
//line machine.go.rl:108
	
//line machine.go.rl:109
	
//line machine.go.rl:110
	
//line machine.go.rl:111

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

	
//line machine.go:71
	{
	 m.cs = start
	}

//line machine.go.rl:135
	
//line machine.go:78
	{
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
	switch  m.cs {
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
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
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
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
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
	}
	goto st_out
	st_case_1:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st2
		case 42:
			goto tr2
		case 46:
			goto tr2
		case 58:
			goto tr4
		case 60:
			goto st55
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
			goto st50
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr3
		}
		goto st0
tr47:
//line machine.go.rl:21

	// cs:1 p:26 pe:55 eof:55 pb:8
	fmt.Println("-- ERROR --")
	fmt.Println("POS: ", m.p)
	fmt.Println("Last mark: ", m.pb)
	fmt.Printf("Str: '%s'\n", string(m.text()))
	fmt.Printf("Before: '%s'\n", string(m.data[:m.p]))
	fmt.Printf("After: '%s'\n", string(m.data[m.p:]))

	return fmt.Errorf("failed at %d", m.p)

	goto st0
//line machine.go:263
st_case_0:
	st0:
		 m.cs = 0
		goto _out
	st2:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st2
		case 42:
			goto tr2
		case 46:
			goto tr2
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
			goto st50
		}
		goto st0
tr2:
//line machine.go.rl:13

	m.pb = m.p

	goto st3
	st3:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line machine.go:311
		switch ( m.data)[( m.p)] {
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
//line machine.go.rl:13

	m.pb = m.p

	goto st4
tr15:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st4
	st4:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:354
		switch ( m.data)[( m.p)] {
		case 112:
			goto st5
		case 117:
			goto st34
		}
		goto st0
	st5:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof5
		}
	st_case_5:
		if ( m.data)[( m.p)] == 114 {
			goto st6
		}
		goto st0
	st6:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof6
		}
	st_case_6:
		if ( m.data)[( m.p)] == 32 {
			goto st7
		}
		goto st0
	st7:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st8
		case 48:
			goto st8
		case 51:
			goto st33
		}
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 50 {
			goto st32
		}
		goto st0
	st8:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 49 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof9
		}
	st_case_9:
		if ( m.data)[( m.p)] == 32 {
			goto st10
		}
		goto st0
	st10:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof12
		}
	st_case_12:
		if ( m.data)[( m.p)] == 58 {
			goto st25
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st14
		}
		goto st0
	st14:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof14
		}
	st_case_14:
		if ( m.data)[( m.p)] == 32 {
			goto st15
		}
		goto st0
	st15:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st16
		}
		goto st0
	st16:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st17
		}
		goto st0
	st17:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof17
		}
	st_case_17:
		if ( m.data)[( m.p)] == 58 {
			goto st18
		}
		goto st0
	st18:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st19
		}
		goto st0
	st19:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st20
		}
		goto st0
	st20:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof20
		}
	st_case_20:
		if ( m.data)[( m.p)] == 58 {
			goto st21
		}
		goto st0
	st21:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st22
		}
		goto st0
	st22:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof22
		}
	st_case_22:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st23
		}
		goto st0
	st23:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr48
		case 58:
			goto tr50
		}
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto tr49
		}
		goto tr47
tr48:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st24
	st24:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:561
		if ( m.data)[( m.p)] == 32 {
			goto st24
		}
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto tr52
		}
		goto tr47
tr52:
//line machine.go.rl:13

	m.pb = m.p

	goto st58
tr49:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st58
	st58:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof58
		}
	st_case_58:
//line machine.go:590
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto st58
		}
		goto tr47
tr50:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st59
	st59:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:610
		if ( m.data)[( m.p)] == 32 {
			goto st24
		}
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto tr52
		}
		goto tr47
	st25:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof25
		}
	st_case_25:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st26
		}
		goto st0
	st26:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof26
		}
	st_case_26:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st27
		}
		goto st0
	st27:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof27
		}
	st_case_27:
		if ( m.data)[( m.p)] == 58 {
			goto st28
		}
		goto st0
	st28:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof28
		}
	st_case_28:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st29
		}
		goto st0
	st29:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof29
		}
	st_case_29:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st30
		}
		goto st0
	st30:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr58
		case 58:
			goto tr50
		}
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto tr49
		}
		goto tr47
tr58:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st31
	st31:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof31
		}
	st_case_31:
//line machine.go:689
		if ( m.data)[( m.p)] == 32 {
			goto st24
		}
		switch {
		case ( m.data)[( m.p)] < 48:
			if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 47 {
				goto tr52
			}
		case ( m.data)[( m.p)] > 57:
			if 58 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
				goto tr52
			}
		default:
			goto tr59
		}
		goto tr47
tr59:
//line machine.go.rl:13

	m.pb = m.p

	goto st60
	st60:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof60
		}
	st_case_60:
//line machine.go:717
		switch {
		case ( m.data)[( m.p)] < 48:
			if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 47 {
				goto st58
			}
		case ( m.data)[( m.p)] > 57:
			if 58 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
				goto st58
			}
		default:
			goto st61
		}
		goto tr47
	st61:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch {
		case ( m.data)[( m.p)] < 48:
			if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 47 {
				goto st58
			}
		case ( m.data)[( m.p)] > 57:
			if 58 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
				goto st58
			}
		default:
			goto st62
		}
		goto tr47
	st62:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch {
		case ( m.data)[( m.p)] < 48:
			if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 47 {
				goto st58
			}
		case ( m.data)[( m.p)] > 57:
			if 58 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
				goto st58
			}
		default:
			goto st63
		}
		goto tr47
	st63:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof63
		}
	st_case_63:
		if ( m.data)[( m.p)] == 32 {
			goto st23
		}
		if 33 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 126 {
			goto st58
		}
		goto tr47
	st32:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof32
		}
	st_case_32:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st9
		}
		goto st0
	st33:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof33
		}
	st_case_33:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 49 {
			goto st9
		}
		goto st0
	st34:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof34
		}
	st_case_34:
		if ( m.data)[( m.p)] == 103 {
			goto st6
		}
		goto st0
tr7:
//line machine.go.rl:13

	m.pb = m.p

	goto st35
tr16:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st35
	st35:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof35
		}
	st_case_35:
//line machine.go:827
		if ( m.data)[( m.p)] == 101 {
			goto st36
		}
		goto st0
	st36:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof36
		}
	st_case_36:
		if ( m.data)[( m.p)] == 99 {
			goto st6
		}
		goto st0
tr8:
//line machine.go.rl:13

	m.pb = m.p

	goto st37
tr17:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st37
	st37:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof37
		}
	st_case_37:
//line machine.go:862
		if ( m.data)[( m.p)] == 101 {
			goto st38
		}
		goto st0
	st38:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof38
		}
	st_case_38:
		if ( m.data)[( m.p)] == 98 {
			goto st6
		}
		goto st0
tr9:
//line machine.go.rl:13

	m.pb = m.p

	goto st39
tr18:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st39
	st39:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof39
		}
	st_case_39:
//line machine.go:897
		switch ( m.data)[( m.p)] {
		case 97:
			goto st40
		case 117:
			goto st41
		}
		goto st0
	st40:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof40
		}
	st_case_40:
		if ( m.data)[( m.p)] == 110 {
			goto st6
		}
		goto st0
	st41:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch ( m.data)[( m.p)] {
		case 108:
			goto st6
		case 110:
			goto st6
		}
		goto st0
tr10:
//line machine.go.rl:13

	m.pb = m.p

	goto st42
tr19:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st42
	st42:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof42
		}
	st_case_42:
//line machine.go:947
		if ( m.data)[( m.p)] == 97 {
			goto st43
		}
		goto st0
	st43:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch ( m.data)[( m.p)] {
		case 114:
			goto st6
		case 121:
			goto st6
		}
		goto st0
tr11:
//line machine.go.rl:13

	m.pb = m.p

	goto st44
tr20:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st44
	st44:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof44
		}
	st_case_44:
//line machine.go:985
		if ( m.data)[( m.p)] == 111 {
			goto st45
		}
		goto st0
	st45:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof45
		}
	st_case_45:
		if ( m.data)[( m.p)] == 118 {
			goto st6
		}
		goto st0
tr12:
//line machine.go.rl:13

	m.pb = m.p

	goto st46
tr21:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st46
	st46:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof46
		}
	st_case_46:
//line machine.go:1020
		if ( m.data)[( m.p)] == 99 {
			goto st47
		}
		goto st0
	st47:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof47
		}
	st_case_47:
		if ( m.data)[( m.p)] == 116 {
			goto st6
		}
		goto st0
tr13:
//line machine.go.rl:13

	m.pb = m.p

	goto st48
tr22:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

//line machine.go.rl:13

	m.pb = m.p

	goto st48
	st48:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:1055
		if ( m.data)[( m.p)] == 101 {
			goto st49
		}
		goto st0
	st49:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof49
		}
	st_case_49:
		if ( m.data)[( m.p)] == 112 {
			goto st6
		}
		goto st0
tr23:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st50
	st50:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:1080
		if ( m.data)[( m.p)] == 79 {
			goto st51
		}
		goto st0
	st51:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof51
		}
	st_case_51:
		if ( m.data)[( m.p)] == 68 {
			goto st52
		}
		goto st0
	st52:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof52
		}
	st_case_52:
		if ( m.data)[( m.p)] == 79 {
			goto st24
		}
		goto st0
tr3:
//line machine.go.rl:13

	m.pb = m.p

	goto st53
	st53:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof53
		}
	st_case_53:
//line machine.go:1114
		if ( m.data)[( m.p)] == 58 {
			goto tr71
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st53
		}
		goto st0
tr4:
//line machine.go.rl:13

	m.pb = m.p

//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st54
tr71:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st54
	st54:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof54
		}
	st_case_54:
//line machine.go:1143
		if ( m.data)[( m.p)] == 32 {
			goto st2
		}
		goto st0
	st55:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof55
		}
	st_case_55:
		if ( m.data)[( m.p)] == 62 {
			goto tr73
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr72
		}
		goto st0
tr72:
//line machine.go.rl:13

	m.pb = m.p

	goto st56
	st56:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof56
		}
	st_case_56:
//line machine.go:1171
		if ( m.data)[( m.p)] == 62 {
			goto tr75
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st56
		}
		goto st0
tr73:
//line machine.go.rl:13

	m.pb = m.p

//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st57
tr75:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

	goto st57
	st57:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:1200
		switch ( m.data)[( m.p)] {
		case 32:
			goto st2
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
		case 84:
			goto st50
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr3
		}
		goto st0
	st_out:
	_test_eof2:  m.cs = 2; goto _test_eof
	_test_eof3:  m.cs = 3; goto _test_eof
	_test_eof4:  m.cs = 4; goto _test_eof
	_test_eof5:  m.cs = 5; goto _test_eof
	_test_eof6:  m.cs = 6; goto _test_eof
	_test_eof7:  m.cs = 7; goto _test_eof
	_test_eof8:  m.cs = 8; goto _test_eof
	_test_eof9:  m.cs = 9; goto _test_eof
	_test_eof10:  m.cs = 10; goto _test_eof
	_test_eof11:  m.cs = 11; goto _test_eof
	_test_eof12:  m.cs = 12; goto _test_eof
	_test_eof13:  m.cs = 13; goto _test_eof
	_test_eof14:  m.cs = 14; goto _test_eof
	_test_eof15:  m.cs = 15; goto _test_eof
	_test_eof16:  m.cs = 16; goto _test_eof
	_test_eof17:  m.cs = 17; goto _test_eof
	_test_eof18:  m.cs = 18; goto _test_eof
	_test_eof19:  m.cs = 19; goto _test_eof
	_test_eof20:  m.cs = 20; goto _test_eof
	_test_eof21:  m.cs = 21; goto _test_eof
	_test_eof22:  m.cs = 22; goto _test_eof
	_test_eof23:  m.cs = 23; goto _test_eof
	_test_eof24:  m.cs = 24; goto _test_eof
	_test_eof58:  m.cs = 58; goto _test_eof
	_test_eof59:  m.cs = 59; goto _test_eof
	_test_eof25:  m.cs = 25; goto _test_eof
	_test_eof26:  m.cs = 26; goto _test_eof
	_test_eof27:  m.cs = 27; goto _test_eof
	_test_eof28:  m.cs = 28; goto _test_eof
	_test_eof29:  m.cs = 29; goto _test_eof
	_test_eof30:  m.cs = 30; goto _test_eof
	_test_eof31:  m.cs = 31; goto _test_eof
	_test_eof60:  m.cs = 60; goto _test_eof
	_test_eof61:  m.cs = 61; goto _test_eof
	_test_eof62:  m.cs = 62; goto _test_eof
	_test_eof63:  m.cs = 63; goto _test_eof
	_test_eof32:  m.cs = 32; goto _test_eof
	_test_eof33:  m.cs = 33; goto _test_eof
	_test_eof34:  m.cs = 34; goto _test_eof
	_test_eof35:  m.cs = 35; goto _test_eof
	_test_eof36:  m.cs = 36; goto _test_eof
	_test_eof37:  m.cs = 37; goto _test_eof
	_test_eof38:  m.cs = 38; goto _test_eof
	_test_eof39:  m.cs = 39; goto _test_eof
	_test_eof40:  m.cs = 40; goto _test_eof
	_test_eof41:  m.cs = 41; goto _test_eof
	_test_eof42:  m.cs = 42; goto _test_eof
	_test_eof43:  m.cs = 43; goto _test_eof
	_test_eof44:  m.cs = 44; goto _test_eof
	_test_eof45:  m.cs = 45; goto _test_eof
	_test_eof46:  m.cs = 46; goto _test_eof
	_test_eof47:  m.cs = 47; goto _test_eof
	_test_eof48:  m.cs = 48; goto _test_eof
	_test_eof49:  m.cs = 49; goto _test_eof
	_test_eof50:  m.cs = 50; goto _test_eof
	_test_eof51:  m.cs = 51; goto _test_eof
	_test_eof52:  m.cs = 52; goto _test_eof
	_test_eof53:  m.cs = 53; goto _test_eof
	_test_eof54:  m.cs = 54; goto _test_eof
	_test_eof55:  m.cs = 55; goto _test_eof
	_test_eof56:  m.cs = 56; goto _test_eof
	_test_eof57:  m.cs = 57; goto _test_eof

	_test_eof: {}
	if ( m.p) == ( m.eof) {
		switch  m.cs {
		case 58, 59, 60, 61, 62, 63:
//line machine.go.rl:17

	fmt.Println("Val: ", string(m.text()))

		case 23, 24, 30, 31:
//line machine.go.rl:21

	// cs:1 p:26 pe:55 eof:55 pb:8
	fmt.Println("-- ERROR --")
	fmt.Println("POS: ", m.p)
	fmt.Println("Last mark: ", m.pb)
	fmt.Printf("Str: '%s'\n", string(m.text()))
	fmt.Printf("Before: '%s'\n", string(m.data[:m.p]))
	fmt.Printf("After: '%s'\n", string(m.data[m.p:]))

	return fmt.Errorf("failed at %d", m.p)

//line machine.go:1318
		}
	}

	_out: {}
	}

//line machine.go.rl:136

	return nil
}
