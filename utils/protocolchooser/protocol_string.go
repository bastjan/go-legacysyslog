// Code generated by "stringer -type=Protocol"; DO NOT EDIT.

package protocolchooser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[IETF-1]
	_ = x[BSD-2]
	_ = x[OctetCountingIETF-3]
	_ = x[OctetCountingBSD-4]
}

const _Protocol_name = "UnknownIETFBSDOctetCountingIETFOctetCountingBSD"

var _Protocol_index = [...]uint8{0, 7, 11, 14, 31, 47}

func (i Protocol) String() string {
	if i < 0 || i >= Protocol(len(_Protocol_index)-1) {
		return "Protocol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Protocol_name[_Protocol_index[i]:_Protocol_index[i+1]]
}