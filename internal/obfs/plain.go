package obfs

func init() {
	register("plain", newPlain)
}

type plain struct {
	obfs
}

func newPlain() IObfs {
	return new(plain)
}

func (p *plain) SetInfo(o ...Setor) {
	for _, f := range o {
		f(&p.obfs)
	}
}

func (p *plain) GetInfo() *obfs {
	return &p.obfs
}

func (p *plain) Encode(d []byte) ([]byte, error) {
	return d, nil
}

func (p *plain) Decode(d []byte) ([]byte, uint64, error) {
	return d, 0, nil
}
