package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"strings"

	"github.com/jWhisper/ssrlocal/errs"
	"golang.org/x/crypto/chacha20"
)

type DecOrEnc int8

const (
	Decrypt DecOrEnc = iota
	Encrypt
)

type cipherInfo struct {
	keyLen    int
	ivLen     int
	newStream func(k, i []byte, doe DecOrEnc) (cipher.Stream, error)
}

var cipherSupported = map[string]*cipherInfo{
	"chacha20-ietf": &cipherInfo{32, 12, newChaCha20_IETF},
}

func register(m string, info *cipherInfo) {
	cipherSupported[m] = info
}

func newChaCha20_IETF(k, i []byte, _ DecOrEnc) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(k, i)
}

type streamCipher struct {
	info    *cipherInfo
	enc     cipher.Stream
	dec     cipher.Stream
	key, iv []byte
}

func NewStreamCipher(method, pass string) (*streamCipher, error) {
	lm := strings.ToLower(method)
	info, ok := cipherSupported[lm]
	if !ok {
		return nil, errs.ErrCipherMethodNotSupported
	}
	key := EVPBytesToKey(pass, info.keyLen)
	return &streamCipher{
		key:  key,
		info: info,
	}, nil
}

// Initializes the block cipher with CFB mode, returns IV.
func (c *streamCipher) initEncrypt() (iv []byte, err error) {
	if c.iv == nil {
		iv = make([]byte, c.info.ivLen)
		rand.Read(iv)
		c.iv = iv
	} else {
		iv = c.iv
	}
	c.enc, err = c.info.newStream(c.key, iv, Encrypt)
	return
}

func (c *streamCipher) initDecrypt(iv []byte) (err error) {
	c.dec, err = c.info.newStream(c.key, iv, Decrypt)
	return
}

func (c *streamCipher) encrypt(dst, src []byte) {
	c.enc.XORKeyStream(dst, src)
}

func (c *streamCipher) decrypt(dst, src []byte) {
	c.dec.XORKeyStream(dst, src)
}

func (c *streamCipher) Copy() *streamCipher {
	cp := *c
	cp.dec, cp.enc = nil, nil
	return &cp
}

func (c *streamCipher) Key() ([]byte, int) {
	return c.key, c.info.keyLen
}

func (c *streamCipher) IV() ([]byte, int) {
	return c.iv, c.info.ivLen
}
