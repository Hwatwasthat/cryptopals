package utilities

import (
	"errors"
)

/*XorBytes takes two byte slices and returns the XORed product and an error.
Both slices must be of equal length. */
func XorBytes(a, b []byte) ([]byte, error) {
	ret := make([]byte, len(a))
	if len(a) != len(b) {
		return []byte{}, errors.New("xorBytes: bytes provided must be of equal length")
	}
	for i := range a {
		ret[i] = a[i] ^ b[i]
	}
	return ret, nil
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
// of the array reopeated to match the length provided.
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
