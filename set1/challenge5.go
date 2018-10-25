package set1

import (
	"cryptopals/utilities"
	"encoding/hex"
	"fmt"
)

func Challenge5() {
	s := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	key := "ICE"
	result := utilities.RepKeyXor([]byte(s), []byte(key))
	rs := hex.EncodeToString(result)
	fmt.Println(rs)
}
