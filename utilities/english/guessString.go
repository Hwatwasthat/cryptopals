package english

import (
	"cryptopals/utilities/crypt/xor"
	"sync"
)
// Empty struct to provide hangar for standard english freq implementation
type Standard struct{}

// Empty struct to provide hangar for chi squared english freq implementation
type ChiSquared struct{}

// tester interface allows for multiple GuessString interfaces to be changed in and out
type Tester interface {
	GuessString(byteSlice ...[]byte) (byte, string)
}


// Guesses holds the return data from a ConcurrentGuessString function run
type Guesses struct { // struct to hold the return values from GuessString
	Idx     int    // Index to ensure data is reconstructed in order
	Guess   string // string of the most probably English sentence from the byte slice provided
	XorByte byte   // The byte that when xored against the byte slice initially provided produces the Guess string
}

func ConcurrentGuessString(idx int, c chan *Guesses, wg *sync.WaitGroup, test Tester, bytes ...[]byte) {
	defer wg.Done()
	ret := &Guesses{Idx: idx}
	ret.XorByte, ret.Guess = test.GuessString(bytes...)
	c <- ret
}

// GuessString takes a slice of strings and returns the guessed keysize and the string most likely to
// be an English sentence. Only works for single byte xor. Uses standard Freq comparison (map of values directly
// translated to score)
func (Standard) GuessString(bytesSlice ...[]byte) (byte, string) {

	var maxVal uint64
	var maxStr string
	var maxIdx int
	for _, b := range bytesSlice {
		for i := 0; i < 128; i++ {
			guess := string(xor.SingleByte(b, byte(i)))
			val := Freq(guess)
			if val > maxVal {
				maxIdx, maxVal, maxStr = i, val, guess
			}
		}
	}
	return byte(maxIdx), maxStr
}

// GuessString takes a slice of strings and returns the guessed keysize and the string most likely to
// be an English sentence. Only works for single byte xor.
// Uses Chi Squared method https://en.wikipedia.org/wiki/Chi-squared_test
func (ChiSquared) GuessString(bytesSlice ...[]byte) (byte, string) {

	var maxVal float64
	var maxStr string
	var maxIdx int
	for _, b := range bytesSlice {
		for i := 0; i < 256; i++ {
			guess := string(xor.SingleByte(b, byte(i)))
			val := FreqChiSquared(guess)
			if val > maxVal {
				maxIdx, maxVal, maxStr = i, val, guess
			}
		}
	}
	return byte(maxIdx), maxStr
}
