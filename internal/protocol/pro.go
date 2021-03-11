package protocol

import (
	"strings"

	"github.com/jWhisper/ssrlocal/errs"
)

type creator func() IProtocol

var protocolSupported = make(map[string]creator)

func register(n string, c creator) {
	protocolSupported[n] = c
}

type IProtocol interface {
	SetInfo()
	GetInfo()
	Encode(data []byte) ([]byte, error)
	Decode(data []byte) ([]byte, uint64, error)
}

type Protocol struct {
}

func NewPro(n string) (IProtocol, error) {
	c, ok := protocolSupported[strings.ToLower(n)]
	if !ok {
		return nil, errs.ErrProtocolNotSupported
	}
	return c(), nil
}
