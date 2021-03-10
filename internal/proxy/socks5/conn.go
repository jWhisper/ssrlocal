package socks5

import (
	"io"
	"net"
	"net/url"

	"github.com/jWhisper/ssrlocal/internal/encrypt"
	"github.com/jWhisper/ssrlocal/internal/obfs"
	"github.com/jWhisper/ssrlocal/internal/protocol"
)

var _ io.ReadWriter = (*SSRTcpConn)(nil)

// SSRTcpConn is a ssr connenct from proxy to dst
type SSRTcpConn struct {
	net.Conn
	sc   *encrypt.StreamCipher
	obfs obfs.Obfs
	pro  protocol.Protocol
}

func DialOpt(o ...Option) (*SSRTcpConn, error) {
	opt := &options{
		server: nil,
		sp:     ":1020",
	}
	for _, f := range o {
		f(opt)
	}
	return Dial(opt.server[0], opt.sp, o...)
}

func Dial(ra, port string, o ...Option) (*SSRTcpConn, error) {
	opt := &options{
		typeof:  "ssr",
		timeout: 5,
	}
	for _, f := range o {
		f(opt)
	}
	host := ra + port
	u := &url.URL{
		Scheme: opt.typeof,
		Host:   host,
	}
	q := u.Query()
	q.Set("encrypt-method", opt.method)
	u.RawQuery = q.Encode()

	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return nil, err
	}

	sc, err := encrypt.NewStreamCipher(opt.method, opt.password)
	if err != nil {
		return nil, err
	}
	//logger.Info("connect....:", host)
	return &SSRTcpConn{
		Conn: conn,
		sc:   sc,
		obfs: obfs.ObfsImp{},
		pro:  protocol.ProImp{},
	}, nil
}

func (ss *SSRTcpConn) Write(b []byte) (n int, err error) {
	return 1, nil
}

func (ss *SSRTcpConn) Read(b []byte) (n int, err error) {
	return 1, nil
}

func (ss *SSRTcpConn) Close() {
}
