package set1

import (
	"cryptopals/crypt/xor"
	"encoding/hex"
	"fmt"
	"os"
)

func Challenge2() {
	first := "1c0111001f010100061a024b53535009181c"
	second := "686974207468652062756c6c277320657965"
	fbytes, err := hex.DecodeString(first)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	sbytes, err := hex.DecodeString(second)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	xorbytes := xor.Slices(fbytes, sbytes)
	for _, v := range xorbytes {
		fmt.Printf("%02X", v)
	}
	fmt.Printf("\n")
}
