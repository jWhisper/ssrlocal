package ssr

import (
	"context"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jWhisper/ssrlocal/configs"
	"github.com/jWhisper/ssrlocal/errs"
	"github.com/jWhisper/ssrlocal/internal/ssr/obfs"
	"github.com/jWhisper/ssrlocal/internal/ssr/pro"
	"github.com/jWhisper/ssrlocal/pkg/log"
)

var logger = log.WithLevelAndMeta(log.DefaultLogger, log.LvInfo, "ssrlocal:")

type server struct {
	addrs  []string
	rp, lp string
	cnf    configs.Cnf
	obfs   obfs.Obfs
	pro    pro.Protocol
	err    error
}

func NewServer(cnf configs.Cnf) (s *server, err error) {
	lp, ok := cnf.Get("local_port").(string)
	if !ok {
		return nil, errs.InvalidLocalPort
	}
	rp, ok := cnf.Get("server_port").(string)
	if !ok {
		return nil, errs.InvalidRemotePort
	}
	addrs := cnf.GetStringSlice("server")
	if len(addrs) < 1 {
		return nil, errs.InvalidRemoteServers
	}
	logger.Print("listen at", lp, ";", "remote servers:", addrs, "remote port", rp)

	s = &server{
		cnf:   cnf,
		addrs: addrs, // a server slice
		rp:    rp,
		lp:    lp,
		obfs:  obfs.ObfsImp{},
		pro:   pro.ProImp{},
		err:   err,
	}
	return
}

func (s *server) Start() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.Print("start ssrlocal server")
	ln, err := net.Listen("tcp", s.lp)
	if err != nil {
		logger.Print("failed to listen port", s.lp)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				conn, err := ln.Accept()
				if err != nil {
					return
				}
				go handTcpConn(s, conn)
			}
		}()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-ctx.Done():
			goto endSig
		case c := <-ch:
			logger.Print("catch a signal:", c)
			switch c {
			case syscall.SIGHUP:
				continue
			default:
				goto endSig
			}
		}
	}
endSig:
	cancel()
	ln.Close()
	return nil
}

func handTcpConn(s *server, conn net.Conn) error {
	defer conn.Close()
	return nil
}

func handUdpConn(s *server, conn net.Conn) error {
	return nil
}
