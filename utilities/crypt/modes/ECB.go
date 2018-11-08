package modes

import (
	"crypto/cipher"
)

type ecb struct {
	cipher.Block
	blockSize int
}

// logically separate the functionality
type ECBEncrypter ecb
type ECBDecryptor ecb

func NewECBEncrypter(block cipher.Block) *ECBEncrypter {
	return &ECBEncrypter{block,
		block.BlockSize(),
	}
}

func NewECBDecrypter(block cipher.Block) *ECBDecryptor {
	return &ECBDecryptor{block,
		block.BlockSize(),
	}
}

func (c *ECBEncrypter) BlockSize() int {
	return c.blockSize
}

func (c *ECBEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%c.blockSize != 0 {
		panic("modes/ECB: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("modes/ECB: output smaller than input")
	}
	core(c, dst, src, c.Encrypt)
}

func (c *ECBDecryptor) BlockSize() int {
	return c.blockSize
}

func (c *ECBDecryptor) CryptBlocks(dst, src []byte) {
	if len(src)%c.blockSize != 0 {
		panic("modes/ECB: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("modes/ECB: output smaller than input")
	}
	core(c, dst, src, c.Decrypt)
}

func core(c cipher.BlockMode, dst, src []byte, cryptFunction func(x, y []byte)) {
	blocksize := c.BlockSize()
	for i := 0; i < len(src); i += blocksize {
		cryptFunction(dst[i:i+blocksize], src[i:i+blocksize])
	}
}
