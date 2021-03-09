package socks5

import (
	"io"
	"net"
	"net/url"

	"github.com/jWhisper/ssrlocal/configs"
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

func NewSSRTcpConn(ra, port string, o ...configs.Option) (*SSRTcpConn, error) {
	opt := &configs.Options{
		Type:    "ssr",
		Timeout: 5,
	}
	for _, f := range o {
		f(opt)
	}
	host := ra + port
	u := &url.URL{
		Scheme: opt.Type,
		Host:   host,
	}
	q := u.Query()
	q.Set("encrypt-method", opt.Method)
	u.RawQuery = q.Encode()

	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return nil, err
	}

	sc, err := encrypt.NewStreamCipher(opt.Method, opt.Password)
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
