package set2

import (
	"cryptopals/crypt/padding"
	"fmt"
)

func Challenge1() {
	fmt.Printf("% 02X\n", padding.Pkcs7([]byte("YELLOW"), 20))
}
