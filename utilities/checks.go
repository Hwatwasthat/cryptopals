package utilities

import (
	"unicode"
)

// GuessString takes a byte array and XORs against 00->FF, returning the most
// likely string to be an english sentence prodiced by this.
func GuessString(b []byte) string {
	guessArray := make([]string, 256)
	for i := 0; i < 256; i++ {
		guessArray[i] = string(SbXor(b, byte(i)))
	}
	return MostEnglish(guessArray)
}

// MostEnglish takes a slice of strings and returns the string most likely to
// be an English sentence.
func MostEnglish(s []string) string {
	var maxVal uint64
	var maxStr string
	for _, line := range s {

		val := EnglishFreq(line)
		if val > maxVal {
			maxVal, maxStr = val, line
		}
	}
	return maxStr
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
		' ': 2400,
	}

	for _, char := range s {
		if val, ok := alphaFreq[unicode.ToUpper(char)]; ok {
			total += val
		}
	}
	return total
}
