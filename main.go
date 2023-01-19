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

	fmt.Println(len(strings.Split(sentence, " ")))
}
