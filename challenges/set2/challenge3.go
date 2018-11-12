package set2

import (
	"cryptopals/utilities/generate"
	"fmt"
)

func Challenge3() {
	fmt.Printf("% 02X\n", generate.RandomAESEncryption([]byte("YELLOW SUBMARINE")))
}
