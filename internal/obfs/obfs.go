package obfs

import (
	"strings"

	"github.com/jWhisper/ssrlocal/errs"
)

type creator func() IObfs

var obfsSupported = make(map[string]creator)

func register(n string, c creator) {
	obfsSupported[n] = c
}

type IObfs interface {
	SetInfo(...Setor)
	GetInfo() *obfs
	Encode(data []byte) ([]byte, error)
	Decode(data []byte) ([]byte, uint64, error)
}

type obfs struct {
	host      string
	port      uint16
	param     string
	iV        []byte
	iVLen     int
	recvIV    []byte
	recvIVLen int
	key       []byte
	keyLen    int
	headLen   int
	tcpMss    int
}

func NewObfs(n string) (IObfs, error) {
	if c, ok := obfsSupported[strings.ToLower(n)]; !ok {
		return nil, errs.ErrObfsNotSupported
	} else {
		return c(), nil
	}
}
