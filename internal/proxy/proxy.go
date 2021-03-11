package proxy

import (
	"github.com/jWhisper/ssrlocal/pkg/log"
)

var (
	tcpKeepAlive = false
	tcpSndBuf    = 4 * 1024
	tcpRcvBuf    = 4 * 1024
)

type server struct {
	lp     string
	err    error
	logger log.Wrapper
}

func NewServer(o ...Option) (s *server, err error) {
	opt := &options{
		lp:      ":1080",
		timeout: 3,
	}

	for _, f := range o {
		f(opt)
	}

	s = &server{
		lp:     opt.lp,
		logger: log.NewWrapper("ssrlocal:"),
	}
	s.logger.Info("listen at:", opt.lp)
	return
}
