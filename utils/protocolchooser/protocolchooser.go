//line protocolchooser.go.rl:1
package protocolchooser

import "errors"

//line protocolchooser.go:9
const protocolchooser_start int = 1
const protocolchooser_first_final int = 12
const protocolchooser_error int = 0

const protocolchooser_en_main int = 1

//line protocolchooser.go.rl:8

// ErrShortBuf is returned if the protocol can't be decided because there is not enough characters.
var ErrShortBuf = errors.New("too few characters")

// Choose returns the protocol based on the syslog header. If the input buffer is too short to
// decide on a protocol ErrShortBuf is returned. This is the only possible error.
// I currently use an initial buffer length of 17 (10 for octet count + 7 for IETF syslog header).
func Choose(buf []byte) (Protocol, error) {
	cs, p, pe := 0, 0, len(buf)
	eof := len(buf)
	_ = eof

	var octetCount bool
	var version bool
	var syslogHeader bool

//line protocolchooser.go:36
	{
		cs = protocolchooser_start
	}

//line protocolchooser.go:41
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
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
		case 12:
			goto st_case_12
		case 9:
			goto st_case_9
		case 13:
			goto st_case_13
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		}
		goto st_out
	st_case_1:
		if (buf)[p] == 60 {
			goto st4
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if (buf)[p] == 32 {
			goto tr3
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st2
		}
		goto st0
	tr3:
//line protocolchooser.go.rl:30
		octetCount = true
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line protocolchooser.go:110
		if (buf)[p] == 60 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if (buf)[p] == 62 {
			goto tr6
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if (buf)[p] == 62 {
			goto tr6
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if (buf)[p] == 62 {
			goto tr6
		}
		goto st0
	tr6:
//line protocolchooser.go.rl:32
		syslogHeader = true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line protocolchooser.go:166
		if 49 <= (buf)[p] && (buf)[p] <= 57 {
			goto st9
		}
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if (buf)[p] == 32 {
			goto tr10
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st10
		}
		goto st0
	tr10:
//line protocolchooser.go.rl:31
		version = true
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line protocolchooser.go:198
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if (buf)[p] == 32 {
			goto tr10
		}
		if 48 <= (buf)[p] && (buf)[p] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if (buf)[p] == 32 {
			goto tr10
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13:
//line protocolchooser.go.rl:33
				return Unknown, ErrShortBuf
//line protocolchooser.go:241
			}
		}

	_out:
		{
		}
	}

//line protocolchooser.go.rl:40

	if !syslogHeader {
		return Unknown, nil
	}

	switch {
	case octetCount && version:
		return OctetCountingIETF, nil
	case octetCount:
		return OctetCountingBSD, nil
	case version:
		return IETF, nil
	}
	return BSD, nil
}
