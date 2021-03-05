package socks

import (
	"io"
	"net"

	"github.com/jWhisper/ssrlocal/internal/ssr/encrypt"
)

var _ io.ReadWriter = (*SSRTcpConn)(nil)

// SSRTcpConn is a ssr connenct from proxy to dst
type SSRTcpConn struct {
	net.TCPConn
	*encrypt.StreamCipher
}

func NewSSRTcpConn(conn net.TCPConn) *SSRTcpConn {
	return &SSRTcpConn{
		TCPConn:      conn,
		StreamCipher: nil,
	}
}

func (ss *SSRTcpConn) Write(b []byte) (n int, err error) {
	return 1, nil
}

func (ss *SSRTcpConn) Read(b []byte) (n int, err error) {
	return 1, nil
}
