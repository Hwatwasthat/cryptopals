package generate

import (
	"crypto/aes"
	"crypto/cipher"
	"cryptopals/utilities/crypt/modes"
	"cryptopals/utilities/crypt/padding"
	"log"
	"math/rand"
	"time"
)

// RandomAESEncryption Takes a provided plaintext as a slice of bytes and generates a random AES key and randomly
// decides between CBC and EBC mode. Returns an encrypted byte slice.
func RandomAESEncryption(plaintext []byte) []byte {
	rand.Seed(time.Now().Unix()) // Seed RNG for goodness

	key := make([]byte, aes.BlockSize)
	rand.Read(key) // Randomly fill key slice

	aesKey, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}

	temp := make([]byte, rand.Intn(6)) // Initialise with random len header
	end := make([]byte, rand.Intn(6))  // Create end bytes
	rand.Read(temp)                    // Fill header
	rand.Read(end)                     // Fill end bytes
	temp = append(temp, plaintext...)  // Build slice to be encrypted
	temp = append(temp, end...)

	if len(temp)%aes.BlockSize != 0 {
		temp = padding.Pkcs7(temp, aes.BlockSize) // Pad to a full 16 byte block
	}

	ret := make([]byte, len(temp))
	switch rand.Intn(2) { // Decide what mode of encryption to use
	case 0: // EBC
		block := modes.NewECBEncrypter(aesKey)
		block.CryptBlocks(ret, temp)
	case 1: // CBC
		iv := make([]byte, aesKey.BlockSize())
		rand.Read(iv) // create random IV
		block := cipher.NewCBCEncrypter(aesKey, iv)
		block.CryptBlocks(ret, temp)
	}
	return ret
}
