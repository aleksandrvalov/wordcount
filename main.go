package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sentence := os.Args[1]

	if sentence == "" {
		fmt.Println(0)

		return
	}

	words := strings.Split(sentence, " ")

	fmt.Println(len(words))
}
