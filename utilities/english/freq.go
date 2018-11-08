package english

import (
	"crypto/aes"
	"cryptopals/utilities"
	"unicode"
)

// Freq takes a string and returns the  value representing
// the likelihood the string is a valid English sentence based on word frequency
func Freq(s string) uint64 {
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

func DetectECB(line []byte) uint64 {
	var ret uint64
	for i := 0; i+aes.BlockSize < len(line); i += aes.BlockSize {
		firstBlock := line[i : i+aes.BlockSize]
		secondBlock := line[i+aes.BlockSize : i+(aes.BlockSize*2)]
		ret += utilities.HamUnsafe(firstBlock, secondBlock)
	}
	return ret
}
