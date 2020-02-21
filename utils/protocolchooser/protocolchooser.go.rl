package protocolchooser

import "errors"

%%{
	machine protocolchooser;
	write data;
}%%

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

	%%{
		write init;
		write exec;
		variable data buf;

		action see_octet_count { octetCount = true }
		action see_version { version = true }
		action see_syslog_header { syslogHeader = true }
		action err_eof { return Unknown, ErrShortBuf }

		octet_count = digit+ ' ' @see_octet_count $eof(err_eof);
		header = '<' $eof(err_eof) digit{1,3} '>' @see_syslog_header $eof(err_eof);
		version = '1'..'9' digit{0,2} ' ' @see_version $eof(err_eof);
	
		main := octet_count? header (version | any - '1'..'9');
	}%%

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
