package socks5

import (
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/jWhisper/ssrlocal/errs"
	"github.com/jWhisper/ssrlocal/internal/encrypt"
	"github.com/jWhisper/ssrlocal/internal/obfs"
	"github.com/jWhisper/ssrlocal/internal/protocol"
	"github.com/jWhisper/ssrlocal/pkg/log"
)

var _ io.ReadWriter = (*SSRTcpConn)(nil)

// SSRTcpConn is a ssr connenct from proxy to dst
type SSRTcpConn struct {
	net.Conn
	sc     *encrypt.StreamCipher
	obfs   obfs.IObfs
	pro    protocol.IProtocol
	logger log.Wrapper
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
		typeof:       "ssr",
		dialtimeout:  5,
		readtimeout:  5,
		writetimeout: 5,
		logger:       log.NewWrapper("ssrConn:"),
	}
	for _, f := range o {
		f(opt)
	}
	opt.logger.Debug("dial opt:", opt)
	host := ra + port
	u := &url.URL{
		Scheme: opt.typeof,
		Host:   host,
	}
	q := u.Query()
	q.Set("encrypt-method", opt.method)
	q.Set("encrypt-key", opt.password)
	q.Set("obfs", opt.obfs)
	q.Set("obfs-param", opt.obfs_param)
	q.Set("protocol", opt.protocol)
	q.Set("protocol", opt.protocol_param)
	u.RawQuery = q.Encode()

	//TODO:use timeout
	conn, err := net.Dial("tcp", u.Host)
	if conn == nil || conn.RemoteAddr() == nil {
		err = errs.ErrNilConn
	}
	if err != nil {
		return nil, err
	}

	sc, err := encrypt.NewStreamCipher(opt.method, opt.password)
	if err != nil {
		return nil, err
	}

	ss := strings.Split(conn.RemoteAddr().String(), ":")
	sp, _ := strconv.Atoi(ss[1])
	ob, err := obfs.NewObfs(opt.obfs)
	if err != nil {
		return nil, err
	}
	ob.SetInfo(obfs.SetHost(ss[0]), obfs.SetPort(uint16(sp)), obfs.SetParam(opt.obfs_param), obfs.SetTcpMss(1460))
	pro := protocol.NewPro(opt.protocol)
	return &SSRTcpConn{
		Conn:   conn,
		sc:     sc,
		obfs:   ob,
		pro:    pro,
		logger: opt.logger,
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
