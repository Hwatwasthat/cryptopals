package set1

import (
	"cryptopals/utilities"
	"fmt"
)

func Challenge6() {
	test1 := "this is a test"
	test2 := "wokka wokka!!!"
	ham, err := utilities.HammingDistance([]byte(test1), []byte(test2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ham)
	}
}
