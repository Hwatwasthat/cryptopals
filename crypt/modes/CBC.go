package modes

import (
	"crypto/cipher"
	"cryptopals/crypt/xor"
	"cryptopals/utilities"
	"log"
)

type CBC struct {
	block     cipher.Block
	blockSize int
	iv        []byte
}

type CBCEncrypter CBC

func NewCBCEncrypter(cipher cipher.Block, iv []byte) *CBCEncrypter {
	return &CBCEncrypter{block: cipher,
		blockSize: cipher.BlockSize(),
		iv:        iv}
}

func (c *CBCEncrypter) Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	iv := make([]byte, c.blockSize)
	copy(iv, c.iv)
	chunks, err := utilities.Chunkify(src, c.blockSize)
	if err != nil {
		log.Fatalln(err)
	}

	for i, chunk := range chunks {
		tmp := xor.Slices(chunk, iv)                                       // XOR block against current IV
		c.block.Encrypt(dst[i*c.blockSize:i*c.blockSize+c.blockSize], tmp) // Encryption
		copy(iv, dst[len(dst)-c.blockSize:])                               // Set next IV
	}
	return dst
}

type CBCDecrypter CBC

func NewCBCDecrypter(cipher cipher.Block, iv []byte) *CBCDecrypter {
	return &CBCDecrypter{block: cipher,
		blockSize: cipher.BlockSize(),
		iv:        iv}
}

func (c *CBCDecrypter) Decrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	iv := make([]byte, c.blockSize)
	copy(iv, c.iv)
	chunksIn, err := utilities.Chunkify(src, c.blockSize)
	if err != nil {
		log.Fatalln(err)
	}

	tmp := make([]byte, c.blockSize)
	for i, chunk := range chunksIn {
		if i != 0 {
			copy(iv, chunk) // Set next IV for iterations > 1
		}

		c.block.Decrypt(tmp, chunk)
		start := i * c.blockSize
		end := (i * c.blockSize) + c.blockSize
		copy(dst[start:end], xor.Slices(tmp, iv))
	}
	return dst
}
