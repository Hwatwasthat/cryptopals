package set1

import (
	"cryptopals/utilities"
	"encoding/base64"
	"fmt"
	"os"
)

func Challenge6(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoding := base64.NewDecoder(base64.StdEncoding, file)
	decodedBytes := make([]byte, 1E6)
	decoding.Read(decodedBytes)
	keysizes, err := utilities.FindKeySize(decodedBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	transposed := utilities.Transpose(decodedBytes, keysizes[0])
	var guesses []int
	for _, s := range transposed {
		guess, _ := utilities.MostEnglish(s)
		guesses = append(guesses, guess)
	}
	fmt.Println(guesses)
}
