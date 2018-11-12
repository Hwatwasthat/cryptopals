package english

import (
	"cryptopals/utilities/crypt/xor"
	"sync"
)

type Standard struct{}
type ChiSquared struct{}

type Tester interface {
	GuessString(byteSlice ...[]byte) (byte, string)
}

type Guesses struct { // struct to hold the return values from GuessString
	Guess   string
	XorByte byte
}

func ConcurrentGuessString(c chan Guesses, wg *sync.WaitGroup, test Tester, bytes ...[]byte) {
	defer wg.Done()
	var ret Guesses
	ret.XorByte, _ = test.GuessString(bytes...) // Concurrent is only interested in the key
	c <- ret
}

// GuessString takes a slice of strings and returns the guessed keysize and the string most likely to
// be an English sentence. Only works for single byte xor.
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
