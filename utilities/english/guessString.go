package english

import (
	"cryptopals/utilities/crypt/xor"
	"sync"
)

type GuessstringByte struct { // struct to hold the return values from GuessString
	Guess   string
	XorByte byte
}

func ConcurrentGuessString(c chan GuessstringByte, wg *sync.WaitGroup, bytes ...[]byte) {
	defer wg.Done()
	var ret GuessstringByte
	ret.XorByte, _ = GuessString(bytes...)
	c <- ret
}

// GuessString takes a slice of strings and returns the string most likely to
// be an English sentence.
func GuessString(bytesSlice ...[]byte) (byte, string) {
	var maxVal uint64
	var maxStr string
	var maxIdx int
	for _, b := range bytesSlice {
		for i := 0; i < 128; i++ {
			guess := string(xor.SingleByte(b, byte(i)))
			val := Freq(guess)
			if val > maxVal {
				maxIdx, maxVal, maxStr = i, val, guess
				//fmt.Printf("%v ", maxIdx)
			}
		}
		//fmt.Println()
	}
	return byte(maxIdx), maxStr
}
