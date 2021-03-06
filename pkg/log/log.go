package log

// Logger is logger interface
type Logger interface {
	Print(msg ...interface{})
}

type lvAndmetaLogger struct {
	l  Logger
	md string
	lv Level
}

func (lm *lvAndmetaLogger) Print(msg ...interface{}) {
	if lm.lv.Active() {
		tmp := []interface{}{lm.lv.String(), lm.md}
		tmp = append(tmp, msg...)
		lm.l.Print(tmp...)
	}
}

// WithLevelAndMeta return a logger has lv and metadata
func WithLevelAndMeta(l Logger, lv Level, md string) Logger {
	return &lvAndmetaLogger{
		l:  l,
		md: md,
		lv: lv,
	}
}
