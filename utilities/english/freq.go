package english

import (
	"math"
	"unicode"
)

//Freq takes a string and returns the  value representing
//the likelihood the string is a valid English sentence based on word frequency
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
		' ': 1270,
	}

	for _, char := range s {
		if val, ok := alphaFreq[unicode.ToUpper(char)]; ok {
			total += val
		}
	}
	return total
}

func FreqChiSquared(s string) float64 {
	var total float64
	alphaFreq := map[rune]float64{
		'A': 08.167,
		'B': 01.492,
		'C': 02.782,
		'D': 04.253,
		'E': 12.702,
		'F': 02.228,
		'G': 02.015,
		'H': 06.094,
		'I': 06.966,
		'J': 00.153,
		'K': 00.772,
		'L': 04.025,
		'M': 02.406,
		'N': 06.749,
		'O': 07.507,
		'P': 01.929,
		'Q': 00.095,
		'R': 05.987,
		'S': 06.327,
		'T': 09.056,
		'U': 02.758,
		'V': 00.978,
		'W': 02.360,
		'X': 00.150,
		'Y': 01.974,
		'Z': 00.074,
		' ': 12.702,
	}

	characters := make(map[rune]int)
	for _, char := range s {
		uChar := unicode.ToUpper(char)
		if _, ok := alphaFreq[uChar]; ok { // if character is in alphaFreq map
			if _, ok := characters[uChar]; ok { // if character has been added before
				characters[uChar]++
			} else { // if not we make a new entry
				characters[uChar] = 1
			}
		}
	}

	lenSlice := float64(len(s))
	for k, v := range characters {
		normalised := float64(v) / lenSlice // divide by length of slice to normalise the value
		total += math.Pow(normalised-alphaFreq[k], 2.0) / alphaFreq[k]
	}

	return total
}
