package proxy

import (
	"context"
	"io"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jWhisper/ssrlocal/configs"
	"github.com/jWhisper/ssrlocal/internal/proxy/socks5"
)

func (s *server) StartTCP() (err error) {
	var bind *net.TCPAddr
	if bind, err = net.ResolveTCPAddr("tcp", s.lp); err != nil {
		return
	}
	ln, err := net.ListenTCP("tcp", bind)
	if err != nil {
		logger.Error("failed to listen local port", s.lp)
		os.Exit(1)
	}

	logger.Info("start ssrlocal server; listening... (curl --socks5 127.0.0.1:1080 http://www.google.com) for test")

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				conn, err := ln.AcceptTCP()
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
			logger.Info("catch a signal:", c)
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

func handTcpConn(s *server, lc *net.TCPConn) {
	var (
		err error
		ra  socks5.Addr
		rc  *socks5.SSRTcpConn
	)
	defer lc.Close()
	err = lc.SetKeepAlive(tcpKeepAlive)
	if err != nil {
		return
	}
	err = lc.SetReadBuffer(tcpRcvBuf)
	if err != nil {
		return
	}
	err = lc.SetWriteBuffer(tcpSndBuf)
	if err != nil {
		return
	}
	ra, err = socks5.HandShake(lc)
	if err != nil {
		return
	}
	rc, err = socks5.NewSSRTcpConn(s.server[0], s.rp, configs.GetOption(s.cnf)...)
	if err != nil {
		logger.Error("failed to connect remote server!", err)
		return
	}
	rc.Write(ra)
	defer rc.Close()
	_, _, err = tcpRelay(rc, lc)
}

func tcpRelay(dst, src io.ReadWriter) (int64, int64, error) {
	return 1, 1, nil
}
