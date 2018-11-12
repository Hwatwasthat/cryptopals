package set1

import (
	"cryptopals/utilities/english"
	"encoding/hex"
	"fmt"
)

func Challenge3(test english.Tester) {
	bytes, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(test.GuessString(bytes))
}
