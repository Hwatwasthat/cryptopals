package modes

import (
	"crypto/cipher"
)

type ECB struct {
	b         cipher.Block
	blockSize int
}

type ECBEncrypter ECB
type ECBDecryptor ECB

func NewECBEncrypter(b cipher.Block) *ECBEncrypter {
	return &ECBEncrypter{b: b,
		blockSize: b.BlockSize(),
	}
}

func NewECBDecrypter(b cipher.Block) *ECBDecryptor {
	return &ECBDecryptor{b: b,
		blockSize: b.BlockSize(),
	}
}

func (c *ECBEncrypter) CryptBlocks(dst, src []byte) {
	core(c, dst, src, c.b.Encrypt)
}

func (c *ECBDecryptor) CryptBlocks(dst, src []byte) {
	core(c, dst, src, c.b.Decrypt)
}

func core(c *ECB, dst, src []byte, cryptFunction func(x, y []byte)) {
	for i := 0; i < len(src); i += c.blockSize {
		cryptFunction(dst[i:i+c.blockSize], src[i:i+c.blockSize])
	}
}
