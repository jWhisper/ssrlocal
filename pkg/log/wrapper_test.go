package log

import "testing"

func TestWrapperLog(t *testing.T) {
	l := NewWrapper("")
	l.Debug("xxx debug")
	l.Info("xxx Info")
	l.Error("xxx Error")
}
