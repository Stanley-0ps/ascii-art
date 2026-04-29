package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	data, err := os.ReadFile("shadow.txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}

	lines := strings.Split(string(data), "\n")
	font := make(map[rune][]string)

	ascii := 32
	index := 1

	for ascii <= 126 {
		if index+8 <= len(lines) {
			font[rune(ascii)] = lines[index : index+8]
		}
		index += 9
		ascii++
	}

	if input == "" {
		return
	}

	if input == "\\n" {
		fmt.Println()
		return
	}

	words := strings.Split(input, "\\n")

	for _, word := range words {

		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				if art, ok := font[ch]; ok {
					fmt.Print(art[i])
				}
			}
			fmt.Println()
		}
	}
}
