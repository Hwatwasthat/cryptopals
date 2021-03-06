/*
Package utilities provides a grouping of functions to assist in the Cryptopals challenges. Root utilities handle basic
functionality that is generic across embedded packages.
*/
package utilities

import (
	"crypto/aes"
	"cryptopals/utilities/crypt/xor"
	"errors"
	"sort"
)

const (
	maxKeysize    = 40 // limit keysearch space
	maxIterations = 16
	accuracyMod   = 1000 // to avoid rounding problems
)

type key struct {
	size    uint16 // No key is larger than a 16byte number
	hamDist uint64
}

// HammingDistance returns the hamming distance of two passed in byte slices.
// Must be of equal length. Returns a uint64 of the hamming distance and an
// error passed up from an interior function.
func HammingDistance(x, y []byte) (uint64, error) {
	if len(x) != len(y) {
		return 0, errors.New("hamming error: must be of same length")
	}
	ret := HamUnsafe(x, y)
	return ret, nil
}

// FindKeySize takes a byte slice and works out the likely size of the key it has been Xored against. takes two
// byte slices and returns a slice of uint16s representing the potential keysizes and an error.
func FindKeySize(x []byte) ([]uint16, error) {
	keyFreq := make([]key, maxKeysize)
	for i := 2; i*2 < len(x) && i < maxKeysize; i++ {
		keyFreq[i].size = uint16(i)

		chunked, err := Chunkify(x, i)
		if err != nil {
			break // we've passed the slice size, time to leave
		}

		var j int
		var temp uint64
		for ; j+1 < len(chunked) && j < maxIterations; j++ {
			temp += HamUnsafe(chunked[j], chunked[j+1])
		}
		keyFreq[i].hamDist = (temp * accuracyMod) / uint64(i) / uint64(j)
	}

	sort.Slice(keyFreq, func(i, j int) bool { return keyFreq[i].hamDist < keyFreq[j].hamDist })
	var possibleKeySizes []uint16
	for i := 0; i < len(keyFreq); i++ {
		if keyFreq[i].hamDist > 0 { // we want to ignore 0 values
			possibleKeySizes = append(possibleKeySizes, keyFreq[i].size)
		}
	}
	return possibleKeySizes, nil
}

func DetectECB(line []byte) uint64 {
	var ret uint64
	for i := 0; i+aes.BlockSize < len(line); i += aes.BlockSize {
		firstBlock := line[i : i+aes.BlockSize]
		secondBlock := line[i+aes.BlockSize : i+(aes.BlockSize*2)]
		ret += HamUnsafe(firstBlock, secondBlock)
	}
	return ret
}

// HamUnsafe provides a hamming function without length checks, to provide easier integration at the expense
// of guaranteeing safety. Use HammingDistance in preference where possible. Takes two byte slices and returns
// a uint64 of the bit level hamming distance.
func HamUnsafe(x, y []byte) uint64 {
	var ret uint64
	xored := xor.Slices(x, y)
	for _, v := range xored {
		var i uint
		for ; i < 8; i++ {
			ret += uint64((v >> i) & 0x01)
		}
	}
	return ret
}
