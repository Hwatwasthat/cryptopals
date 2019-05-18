package utilities

import (
	"errors"
)

// GenerateEqKey takes a byte slice to be used as a key and a length value
// which must be longer than the array.  returns a new byte array consisting
// of the array repeated to match the length provided.
func GenerateEqKey(key []byte, length int) ([]byte, error) {
	keylen := len(key)

	if keylen > length {
		return []byte{}, errors.New("error: key must be shorter than block to cipher")
	}
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = key[i%keylen]
	}
	return ret, nil
}

// Chunkify takes a given byte slice and returns a slice of slices of bytes, each the size of the provided chunk size,
// and an error. chunkSize must be less than the provided slices length. Final slice is padded with 0's if
// len(src) % chunkSize != 0.
func Chunkify(src []byte, chunkSize int) ([][]byte, error) {
	if len(src) < chunkSize {
		return [][]byte{}, errors.New("error: slice to be chunked must be larger than chunk size")
	}

	var ret [][]byte
	var i int
	for ; i+chunkSize < len(src); i += chunkSize {
		ret = append(ret, src[i:i+chunkSize])
	}

	if i < len(src) { // Check if we have leftovers
		remainder := len(src[i:])  // find out how many bytes are left
		ret = append(ret, src[i:]) // fill slice with leftovers
		for i := 0; i < len(ret[0])-remainder; i++ {
			ret[len(ret)-1] = append(ret[len(ret)-1], 0) // pad remainder with 0s
		}

	}
	return ret, nil
}

// Transpose takes a slice of bytes and a chunkSize and returns a slice of slices of bytes. Each slice is composed of
// the respective bytes at the index % chunkSize of the original slice.
// Example:
// slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
// ret := Transpose(slice, 4)
// ret -> [][]int{{0, 4, 8, 12}, {1, 5, 9, 13}, {2, 6, 10, 14}, {3, 7, 11, 15}}
func Transpose(bytes []byte, chunkSize uint16) [][]byte {
	ret := make([][]byte, chunkSize)
	for i := 0; i < len(bytes); i++ {
		pos := i % int(chunkSize)
		ret[pos] = append(ret[pos], bytes[i])
	}
	return ret
}
