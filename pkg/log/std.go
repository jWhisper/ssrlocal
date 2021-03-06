package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var _ Logger = (*stdlogger)(nil)

// DefaultLogger Default
var DefaultLogger = newStd(os.Stderr)

type stdlogger struct {
	log  *log.Logger
	pool *sync.Pool
}

func (std *stdlogger) Print(msg ...interface{}) {
	// TODO: use buffer
	std.log.Println(msg...)
}

func newStd(w io.Writer) *stdlogger {
	return &stdlogger{
		log: log.New(w, "", log.LstdFlags),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(sync.Pool)
			},
		},
	}
}
