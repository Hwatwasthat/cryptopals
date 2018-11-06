package padding

func Pkcs7(src []byte, blockLen int) []byte {
	pad := blockLen - (len(src) % blockLen)
	dst := make([]byte, 0, len(src)+pad)
	dst = append(dst, src...)
	for i := len(src); i < len(src)+pad; i++ {
		dst = append(dst, byte(pad))
	}
	return dst
}
