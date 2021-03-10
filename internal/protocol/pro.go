package protocol

type IProtocol interface {
}

type Protocol struct {
}

func NewPro(n string) *Protocol {
	return &Protocol{}
}
