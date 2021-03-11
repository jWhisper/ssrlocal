package protocol

func init() {
	register("auth_aes128_md5", newAES128MD5)
}

type aes128md5 struct {
}

func newAES128MD5() IProtocol {
	return new(aes128md5)
}

func (a *aes128md5) SetInfo() {

}

func (a *aes128md5) GetInfo() {}

func (a *aes128md5) Encode(d []byte) ([]byte, error) {
	return d, nil
}

func (a *aes128md5) Decode(d []byte) ([]byte, uint64, error) {
	return d, 0, nil
}
