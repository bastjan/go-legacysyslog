package legacysyslog_test

import "testing"
import "github.com/bastjan/go-legacysyslog"

func TestMachineParse(t *testing.T) {
	line := []byte(`<13>Dec  1 09:15:22 0304eebf3c1e root[3037]: yolo`)
	err := legacysyslog.NewMachine().Parse(line)
	if err != nil {
		t.Error(err)
	}
}

func TestMachineParseCiscoSequence(t *testing.T) {
	line := []byte(`000123: *Dec  1 09:15:22: 0304eebf3c1e root[3037]: yolo`)
	legacysyslog.NewMachine().Parse(line)
	err := legacysyslog.NewMachine().Parse(line)
	if err != nil {
		t.Error(err)
	}
}
