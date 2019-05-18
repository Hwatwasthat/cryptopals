package set2

import (
	"cryptopals/utilities/crypt/tests"
	"cryptopals/utilities/generate"
	"fmt"
)

func Challenge3() {
	crypted := generate.RandomAESEncryption([]byte("YELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINE"))
	if tests.IsECB(crypted) {
		fmt.Println("ECB encrypted")
	} else {
		fmt.Println("CBC encrypted")
	}
}
