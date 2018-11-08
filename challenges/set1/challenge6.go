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

func Challenge6(filename string) {
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

	wg := &sync.WaitGroup{}
	c := make(chan english.GuessstringByte)
	for _, s := range transposed {
		wg.Add(1)
		go english.ConcurrentGuessString(c, wg, s) // concurrent implementation of MostEnglish
	}

	go monitorWG(wg, c)

	var guesses []byte
	for ret := range c {
		guesses = append(guesses, ret.XorByte)
	}

	result := xor.RepKey(decodedBytes, guesses)
	fmt.Println(string(result))

}

func monitorWG(wg *sync.WaitGroup, c chan english.GuessstringByte) {
	wg.Wait()
	close(c)
}
