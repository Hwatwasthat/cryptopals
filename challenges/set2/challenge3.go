package set2

import (
	"cryptopals/utilities/generate"
	"fmt"
)

func Challenge3() {
	fmt.Println(generate.RandomAESEncryption([]byte("YELLOW SUBMARINE")))
}
