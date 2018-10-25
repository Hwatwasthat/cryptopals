package set1

import (
	"cryptopals/utilities"
	"encoding/hex"
	"fmt"
)

func Challenge3() {
	bytes, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(GuessString(bytes))
}

func GuessString(b []byte) string {
	guessArray := make([]string, 256)
	for i := 0; i < 256; i++ {
		guessArray[i] = string(utilities.SbXor(b, byte(i)))
	}
	return utilities.MostEnglish(guessArray)
}
