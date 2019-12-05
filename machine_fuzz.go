// +build gofuzz

package legacysyslog

func Fuzz(data []byte) int {
	m := NewMachine()
	m.Parse(data)
	// TODO: valid input should return 1, but we don't really know what is valid right now
	return 0
}
