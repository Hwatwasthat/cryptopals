package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"cryptopals/crypt/modes"
	"cryptopals/crypt/padding"
	"log"
	"math/rand"
	"time"
)

func RandomAES(plaintext []byte) []byte {
	rand.Seed(time.Now().Unix()) // Seed RNG for goodness

	key := make([]byte, aes.BlockSize)
	rand.Read(key) // Randomly fill key slice

	aesKey, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}

	start := 5 + rand.Intn(6) // length of random padding for front of plaintext (5 -10)
	end := 5 + rand.Intn(6)   // length of random padding for end of plaintext (5 -10)
	temp := make([]byte, len(plaintext)+start+end)
	rand.Read(temp[:start])
	rand.Read(temp[end:])
	copy(temp[start:], plaintext)

	if len(temp)%aes.BlockSize != 0 {
		temp = padding.Pkcs7(temp, aes.BlockSize) // Pad to a full 16 byte block
	}

	ret := make([]byte, len(temp))
	if rand.Intn(2) == 1 { // If 1 we do CBC, Else we do EBC
		iv := make([]byte, aesKey.BlockSize())
		rand.Read(iv) // create random IV if we're doing CBC
		block := cipher.NewCBCEncrypter(aesKey, iv)
		block.CryptBlocks(ret, temp)
	} else {
		block := modes.NewECBEncrypter(aesKey)
		block.CryptBlocks(ret, temp)
	}
	return ret
}
