package utilities

import (
	"errors"
	"sort"
)

const (
	maxKeysize    = 32   // limit keysearch space
	maxIterations = 8    // limit amount of repeat checks
	accuracyMod   = 1000 // to avoid rounding problems
)

// HammingDistance returns the hamming distance of two passed in byte slices.
// Must be of equal length. Returns a uint64 of the hamming distance and an
// error passed up from an interior function.
func HammingDistance(x, y []byte) (uint64, error) {
	if len(x) != len(y) {
		return 0, errors.New("hamming error: must be of same length")
	}
	ret := hamUnsafe(x, y)
	return ret, nil
}

func FindKeysize(x []byte) ([]uint64, error) {
	possibleKeysizes := make([]uint64, maxKeysize)
	for i, j := 2, 0; i*maxIterations < len(x) && i < maxKeysize; i, j = i+1, 0 {
		var temp uint64
		for ; j < maxIterations; j++ {
			offset := i * j
			nextOffset := i * (j + 1)
			temp += hamUnsafe(x[offset:offset+i], x[nextOffset:nextOffset+i])
		}
		possibleKeysizes[i] = (temp * accuracyMod) / uint64(i) / uint64(j)
	}

	sort.Slice(possibleKeysizes, func(i, j int) bool { return possibleKeysizes[i] < possibleKeysizes[j] })
	startOfReturn := 0 // we want to ignore 0 values
	for i := 0; i < len(possibleKeysizes); i++ {
		if possibleKeysizes[i] > 0 {
			startOfReturn = i
			break
		}
	}
	return possibleKeysizes[startOfReturn : startOfReturn+5], nil
}

func hamUnsafe(x, y []byte) uint64 {
	var ret uint64
	xored, _ := XorBytes(x, y) // safe to ignore error, sepereate check in wrapper
	for _, v := range xored {
		var i uint
		for ; i < 8; i++ {
			ret += uint64((v >> i) & 1)
		}
	}
	return ret
}
