package set2

import (
	"cryptopals/crypt/tests"
	"fmt"
)

func Challenge3() {
	fmt.Println(tests.RandomAES([]byte("YELLOW SUBMARINE")))
}
