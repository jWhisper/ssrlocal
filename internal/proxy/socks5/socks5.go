package socks5

import (
	"io"
	"net"
	"strconv"

	"github.com/jWhisper/ssrlocal/errs"
)

// MaxAddrLen is the maximum size of socks address in bytes
const (
	MaxAddrLen = 1 + 1 + 255 + 2
)

// SOCKS request commands as defined in RFC 1928 section 4.
const (
	CmdConnect      = 1
	CmdBind         = 2
	CmdUDPAssociate = 3
)

// SOCKS address types as defined in RFC 1928 section 5.
const (
	AtypIPv4       = 1
	AtypDomainName = 3
	AtypIPv6       = 4
)

type Addr []byte

func (a Addr) String() string {
	var host, port string

	switch a[0] { // address type
	case AtypDomainName:
		host = string(a[2 : 2+int(a[1])])
		port = strconv.Itoa((int(a[2+int(a[1])]) << 8) | int(a[2+int(a[1])+1]))
	case AtypIPv4:
		host = net.IP(a[1 : 1+net.IPv4len]).String()
		port = strconv.Itoa((int(a[1+net.IPv4len]) << 8) | int(a[1+net.IPv4len+1]))
	case AtypIPv6:
		host = net.IP(a[1 : 1+net.IPv6len]).String()
		port = strconv.Itoa((int(a[1+net.IPv6len]) << 8) | int(a[1+net.IPv6len+1]))
	}

	return net.JoinHostPort(host, port)
	//return string(a)
}

func readAddr(r io.Reader, b []byte) (Addr, error) {
	if len(b) < MaxAddrLen {
		return nil, io.ErrShortBuffer
	}
	_, err := io.ReadFull(r, b[:1]) // read 1st byte for address type
	if err != nil {
		return nil, err
	}

	switch b[0] {
	case AtypDomainName:
		_, err = io.ReadFull(r, b[1:2]) // read 2nd byte for domain length
		if err != nil {
			return nil, err
		}
		_, err = io.ReadFull(r, b[2:2+int(b[1])+2])
		return b[:1+1+int(b[1])+2], err
	case AtypIPv4:
		_, err = io.ReadFull(r, b[1:1+net.IPv4len+2])
		return b[:1+net.IPv4len+2], err
	case AtypIPv6:
		_, err = io.ReadFull(r, b[1:1+net.IPv6len+2])
		return b[:1+net.IPv6len+2], err
	}

	return nil, errs.ErrAddressNotSupported
}

func ReadAddr(r io.Reader) (Addr, error) {
	return readAddr(r, make([]byte, MaxAddrLen))
}

func HandShake(rw io.ReadWriter) (Addr, error) {
	// Read RFC 1928 for request and reply structure and sizes.There are two steps
	// https://www.cnblogs.com/Full--Stack/p/8041880.html
	buf := make([]byte, MaxAddrLen)
	var err error
	// step1: read VER, NMethods, Methods
	if _, err = io.ReadFull(rw, buf[:2]); err != nil {
		return nil, err
	}
	nmethods := buf[1]
	if _, err = io.ReadFull(rw, buf[:nmethods]); err != nil {
		return nil, err
	}
	// write ver:5 method:0
	if _, err = rw.Write([]byte{5, 0}); err != nil {
		return nil, err
	}

	// step2: read ver cmd rsv atyp dst.addr dst.port
	if _, err = io.ReadFull(rw, buf[:3]); err != nil {
		return nil, err
	}
	cmd := buf[1]
	addr, err := readAddr(rw, buf)
	if err != nil {
		return nil, err
	}
	switch cmd {
	case CmdConnect:
		_, err = rw.Write([]byte{5, 0, 0, 1, 0, 0, 0, 0, 0, 0}) // SOCKS v5, reply succeeded
	case CmdUDPAssociate:
		//if !UDPEnabled {
		//	return nil, ErrCommandNotSupported
		//}
		//listenAddr := ParseAddr(rw.(net.Conn).LocalAddr().String())
		//_, err = rw.Write(append([]byte{5, 0, 0}, listenAddr...)) // SOCKS v5, reply succeeded
		//if err != nil {
		//	return nil, ErrCommandNotSupported
		//}
		//err = InfoUDPAssociate
		return nil, errs.ErrCommandNotSupported
	default:
		return nil, errs.ErrCommandNotSupported
	}
	return addr, err // skip ver, cmd, rsv fields
}
