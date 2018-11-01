package utilities

import (
	"errors"
)

/*XorBytes takes two byte slices and returns the XORed product.
Both slices must be of equal length. */
func XorBytes(a, b []byte) []byte {
	ret := make([]byte, len(a))
	for i := range a {
		ret[i] = a[i] ^ b[i]
	}
	return ret
}

//SbXor does a single byte xor against a provided byte array
func SbXor(arr []byte, b byte) []byte {
	ret := make([]byte, len(arr))
	for i, v := range arr {
		ret[i] = v ^ b
	}
	return ret
}

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

// RepKeyXor takes two provided byte arrays, one key and one to be ciphered and
// repeatedly XORs, repeating the key when necessary to match the length of the
// array to be ciphered.
func RepKeyXor(b []byte, key []byte) []byte {
	ret := make([]byte, len(b))
	keylen := len(key)
	for i, v := range b {
		ret[i] = key[i%keylen] ^ v
	}
	return ret
}

// Chunkify takes a given byte slice and returns a slice of slices, each the size of the provided chunk size.
func Chunkify(bytes []byte, chunkSize int) ([][]byte, error) {
	if len(bytes) < chunkSize {
		return [][]byte{}, errors.New("error: slice to be chunked must be larger than chunk size")
	}

	var ret [][]byte
	var i int
	for ; i+chunkSize < len(bytes); i += chunkSize {
		ret = append(ret, bytes[i:i+chunkSize])
	}

	if i < len(bytes) { // Check if we have leftovers
		remainder := len(bytes[i:])  // find out how many bytes are left
		ret = append(ret, bytes[i:]) // fill slice with leftovers
		for i := 0; i < len(ret[0])-remainder; i++ {
			ret[len(ret)-1] = append(ret[len(ret)-1], 0) // pad remainder with 0s
		}

	}
	return ret, nil
}

func Transpose(bytes []byte, chunkSize uint16) [][]byte {
	ret := make([][]byte, chunkSize)
	for i := 0; i < len(bytes); i++ {
		pos := i % int(chunkSize)
		ret[pos] = append(ret[pos], bytes[i])
	}
	return ret
}
