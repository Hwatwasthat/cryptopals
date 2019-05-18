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
	"sync"
)

func Challenge6(filename string, test english.Tester, concurrentOrNot int) {
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

	switch concurrentOrNot { // For testing concurrent implentation. Result = runs about the same, but needs syncing
	case 0:
		results(decodedBytes, transposed, test)
	case 1:
		concurrentResults(decodedBytes, transposed, test)
	}

}

func concurrentResults(decodedBytes []byte, transposed [][]byte, test english.Tester) {
	results := make(chan *english.Guesses)
	wg := &sync.WaitGroup{}

	for idx, s := range transposed {
		wg.Add(1)
		go english.ConcurrentGuessString(idx, results, wg, test, s)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	byteGuess := make([]byte, len(transposed))

	for result := range results {
		byteGuess[result.Idx] = result.XorByte
	}

	result := xor.RepKey(decodedBytes, byteGuess)
	fmt.Println(string(result))
}

func results(decodedBytes []byte, transposed [][]byte, test english.Tester) {
	var byteGuess []byte

	for _, s := range transposed {
		ret, _ := test.GuessString(s)
		byteGuess = append(byteGuess, ret)
	}

	result := xor.RepKey(decodedBytes, byteGuess)
	fmt.Println(string(result))
}
