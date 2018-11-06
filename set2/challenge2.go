package set2

import (
	"crypto/aes"
	"cryptopals/crypt/modes"
	"cryptopals/utilities"
	"fmt"
	"log"
)

func Challenge2() {
	aesKey, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatalln(err)
	}

	iv, err := utilities.GenerateEqKey([]byte{1, 2, 3, 4}, aesKey.BlockSize())
	if err != nil {
		log.Fatalln(err)
	}

	encrypter := modes.NewCBCEncrypter(aesKey, iv)
	encrypted := encrypter.Encrypt([]byte("TESTTESTTESTTEST"))

	decryptering := modes.NewCBCDecrypter(aesKey, iv)
	decrypted := decryptering.Decrypt(encrypted)
	fmt.Println(string(decrypted))
}
