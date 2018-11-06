package set1

import (
	"cryptopals/utilities"
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

	// Create a base64 decoder for the file, which implements io.reader, which we pass to be read
	decodedBytes, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, file))
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
	c := make(chan utilities.GuessstringByte)
	for _, s := range transposed {
		wg.Add(1)
		go utilities.ConMMostEnglish(c, wg, s) // concurrent implementation of MostEnglish
	}

	go monitorWG(wg, c) // close channel when its empty by keeping an eye on wg size, only launch after filling!

	var guesses []byte
	for ret := range c {
		guesses = append(guesses, ret.XorByte)
	}

	result := utilities.RepKeyXor(decodedBytes, guesses)
	fmt.Println(string(result))

}

func monitorWG(wg *sync.WaitGroup, c chan utilities.GuessstringByte) {
	wg.Wait()
	close(c)
}
