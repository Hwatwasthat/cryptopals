package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func Challenge1() {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := StrHexToBase64(s)
	fmt.Println(b64)
}
func StrHexToBase64(s string) string {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
