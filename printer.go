package main

import (
	"fmt"
	"strings"
)

func PrintAscii(text string, banner []string) {
	words := strings.Split(text, "\\n")

	for w, word := range words {

		if word == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			line := ""

			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				index := (int(char)-32)*9 + row + 1
				line += banner[index]
			}
			fmt.Println(line)
		}
		if w!=len(words)-1{
			
		}
	}
}
