package log

import "fmt"

var globalLevel = LvDebug

func init() {
	globalLevel = LvInfo
}

type Logger interface {
	Lv(lv Level) Logger
	Print(msg ...interface{})
}

type logger struct {
	log  Logger
	inLv Level
}

func (l *logger) Lv(lv Level) Logger {
	// TODO: thread safe
	l.inLv = lv
	return l.log.Lv(lv)
}

func (l *logger) Print(msg ...interface{}) {
	if l.inLv >= globalLevel {
		tmp := make([]interface{}, len(msg)+1)
		tmp[0] = fmt.Sprintf("%s", l.inLv)
		l.log.Print(append(tmp, msg...)...)
	}
}

func With(l Logger) Logger {
	return &logger{
		log:  l,
		inLv: LvDebug,
	}
}
