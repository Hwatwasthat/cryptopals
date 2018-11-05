package set1

import (
	"crypto/aes"
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

	aesKey, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	ret := make([]byte, len(decodedBytes))
	for i := 0; i < len(decodedBytes); i += aes.BlockSize {
		aesKey.Decrypt(ret[i:i+aes.BlockSize], decodedBytes[i:i+aes.BlockSize])
	}

	fmt.Println(string(ret))
}
