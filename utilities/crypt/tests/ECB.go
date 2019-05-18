package tests

import (
	"crypto/aes"
	"fmt"
)

func IsECB(ciphertext []byte) bool {
	var matches int
	var ret bool

	for i := 0; i < len(ciphertext); i += 1 {
		for j := i + aes.BlockSize; j < len(ciphertext); j += aes.BlockSize {
			if ciphertext[i] == ciphertext[j] { // if it matches, we have an EBC indication
				matches++
			}
		}
	}

	if matches > len(ciphertext)/10 { // Randomly chosen match values, just likelihood with non repeating sentences
		ret = true
	}

	fmt.Println(matches, len(ciphertext))
	return ret
}
