package ssr

import (
	"context"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jWhisper/ssrlocal/internal/ssr/socks"
)

func (s *server) StartTCP() (err error) {
	var bind *net.TCPAddr
	if bind, err = net.ResolveTCPAddr("tcp", s.lp); err != nil {
		return
	}
	ln, err := net.ListenTCP("tcp", bind)
	if err != nil {
		logger.Print("failed to listen local port", s.lp)
		os.Exit(1)
	}

	logger.Print("start ssrlocal server; listening... (curl --socks5 127.0.0.1:1080 http://www.google.com) for test")

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

func handTcpConn(s *server, conn *net.TCPConn) (err error) {
	defer conn.Close()
	if err = conn.SetKeepAlive(tcpKeepAlive); err != nil {
		return
	}
	if err = conn.SetReadBuffer(tcpRcvBuf); err != nil {
		return
	}
	if err = conn.SetWriteBuffer(tcpSndBuf); err != nil {
		return
	}
	dstAddr, err := socks.HandShake(conn)
	return
}
