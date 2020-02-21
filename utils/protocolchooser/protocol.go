//go:generate stringer -type=Protocol

package protocolchooser

// Protocol represents the syslog protocol.
type Protocol int

const (
	Unknown Protocol = iota
	IETF
	BSD
	OctetCountingIETF
	OctetCountingBSD
)
