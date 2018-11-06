package set1

import (
	"crypto/aes"
	"cryptopals/crypt/modes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Challenge7(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	decodedBytes, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, file))
	if err != nil {
		log.Fatalln(err)
	}

	key := ("YELLOW SUBMARINE")
	ret := AESECBEncrypt(decodedBytes, []byte(key))

	fmt.Println(string(ret))
}

func AESECBEncrypt(src []byte, key []byte) []byte {
	aesKey, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}

	cryptor := modes.NewECBEncrypter(aesKey)
	ret := make([]byte, len(src))
	cryptor.CryptBlocks(ret, src)
	return ret
}
