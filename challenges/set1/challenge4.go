package set1

import (
	"cryptopals/utilities/english"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Challenge4(filename string, test english.Tester) {
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	filestring := string(filebytes)
	lines := strings.Split(filestring, "\n")
	bestGuess := guessFile(lines, test)
	fmt.Println(bestGuess)
}

func guessFile(lines []string, test english.Tester) string {
	guessArray := make([][]byte, len(lines))
	for i, line := range lines {
		bytes, err := hex.DecodeString(line)
		if err != nil {
			log.Fatal(err)
		}
		_, guess := test.GuessString(bytes)
		guessArray[i] = []byte(guess)
	}
	_, ret := english.Standard{}.GuessString(guessArray...)
	return ret
}
