package set1

import (
	"bufio"
	"crypto/aes"
	"cryptopals/utilities"
	"fmt"
	"log"
	"os"
	"sort"
)

type hamLine struct {
	line int
	ham  uint64
}

func Challenge8(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 40)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var lineNo int
	hammed := make([]hamLine, len(lines))
	for _, line := range lines {
		for i := 0; i+aes.BlockSize < len(line); i += aes.BlockSize {
			chunk1 := line[i : i+aes.BlockSize]
			chunk2 := line[i+aes.BlockSize : i+(aes.BlockSize*2)]
			temp, err := utilities.HammingDistance([]byte(chunk1), []byte(chunk2))
			if err != nil {
				log.Fatalln(err)
			}
			hammed[lineNo].line = lineNo
			hammed[lineNo].ham += temp
		}
		lineNo++
	}
	sort.Slice(hammed, func(i, j int) bool { return hammed[i].ham < hammed[j].ham })
	fmt.Println(hammed[0].line, hammed[0].ham)
}
