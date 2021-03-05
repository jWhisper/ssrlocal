package log

import (
	"io"
	"log"
	"sync"
)

var _ Logger = (*stdlogger)(nil)

type stdlogger struct {
	log  *log.Logger
	pool *sync.Pool
	inLv Level
}

func (std *stdlogger) Lv(lv Level) Logger {
	std.inLv = lv
	return std
}

func (std *stdlogger) Print(msg ...interface{}) {
	// TODO: use buffer
	std.log.Println(msg)
}

func newStd(w io.Writer) *stdlogger {
	return &stdlogger{
		inLv: LvDebug,
		log:  log.New(w, "", log.LstdFlags),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(sync.Pool)
			},
		},
	}
}
