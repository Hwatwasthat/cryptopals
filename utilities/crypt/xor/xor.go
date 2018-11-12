package xor

// Slices takes two byte slices and returns the XORed product.
// Both slices must be of equal length.
func Slices(a, b []byte) []byte {
	ret := make([]byte, len(a))
	for i := range a {
		ret[i] = a[i] ^ b[i]
	}
	return ret
}

// SingleByte does a single byte xor against a provided byte slice using the passed byte.
func SingleByte(arr []byte, b byte) []byte {
	ret := make([]byte, len(arr))
	for i, v := range arr {
		ret[i] = v ^ b
	}
	return ret
}

// RepKey takes two provided byte arrays, one key and one to be ciphered and
// repeatedly XORs, repeating the key when necessary to match the length of the
// array to be ciphered.
func RepKey(b []byte, key []byte) []byte {
	ret := make([]byte, len(b))
	keylen := len(key)
	for i, v := range b {
		ret[i] = key[i%keylen] ^ v
	}
	return ret
}
