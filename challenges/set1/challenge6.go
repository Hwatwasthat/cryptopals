package set1

import (
	"cryptopals/utilities"
	"cryptopals/utilities/crypt/xor"
	"cryptopals/utilities/english"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Challenge6(filename string, test english.Tester) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	b64 := base64.NewDecoder(base64.StdEncoding, file)
	decodedBytes, err := ioutil.ReadAll(b64)
	if err != nil {
		log.Fatalln(err)
	}

	keysize, err := utilities.FindKeySize(decodedBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	transposed := utilities.Transpose(decodedBytes, keysize[0]) // dangerous assumption that keysize[0] is accurate

	var guesses []byte
	for _, s := range transposed {
		ret, _ := test.GuessString(s)
		guesses = append(guesses, ret)
	}

	result := xor.RepKey(decodedBytes, guesses)
	fmt.Println(string(result))

}
