package log

import (
	"os"
	"testing"
)

func Testlogger(t *testing.T) {
	stdlog := newStd(os.Stderr)
	log := NewLogger(stdlog)

	log.Lv(2).Print("test")
}
