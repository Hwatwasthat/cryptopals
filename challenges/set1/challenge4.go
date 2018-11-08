package set1

import (
	"cryptopals/utilities/english"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Challenge4() {
	filebytes, err := ioutil.ReadFile("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	filestring := string(filebytes)
	lines := strings.Split(filestring, "\n")
	bestGuess := guessFile(lines)
	fmt.Println(bestGuess)
}

func guessFile(lines []string) string {
	guessArray := make([][]byte, len(lines))
	for i, line := range lines {
		bytes, err := hex.DecodeString(line)
		if err != nil {
			log.Fatal(err)
		}
		_, guess := english.GuessString(bytes)
		guessArray[i] = []byte(guess)
	}
	_, ret := english.GuessString(guessArray...)
	return ret
}
