package utilities

import (
	"unicode"
)

type Stringval struct {
	guess string
	val   int
}

// MostEnglish takes a slice of strings and returns the string most likely to
// be an English sentence.
func MostEnglish(bytes ...[]byte) (int, string) {
	var maxVal uint64
	var maxStr string
	var maxIdx int
	for _, b := range bytes {
		for i := 0; i < 128; i++ {
			guess := string(SbXor(b, byte(i)))
			val := EnglishFreq(guess)
			if val > maxVal {
				maxIdx, maxVal, maxStr = i, val, guess
				//fmt.Printf("%v ", maxIdx)
			}
		}
		//fmt.Println()
	}
	return maxIdx, maxStr
}

// EnglishFreq takes a string and returns the  value representing
// the likelihood the string is a valid English sentence based on word frequency
func EnglishFreq(s string) uint64 {
	var total uint64
	alphaFreq := map[rune]uint64{
		'A': 816,
		'B': 149,
		'C': 278,
		'D': 425,
		'E': 1270,
		'F': 222,
		'G': 201,
		'H': 609,
		'I': 696,
		'J': 15,
		'K': 77,
		'L': 402,
		'M': 240,
		'N': 674,
		'O': 750,
		'P': 192,
		'Q': 9,
		'R': 598,
		'S': 632,
		'T': 905,
		'U': 275,
		'V': 236,
		'W': 20,
		'X': 15,
		'Y': 197,
		'Z': 7,
		' ': 1200,
	}

	for _, char := range s {
		if val, ok := alphaFreq[unicode.ToUpper(char)]; ok {
			total += val
		}
	}
	return total
}
