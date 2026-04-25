package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}
	input := os.Args[1]

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	PrintAscii(input, banner)
}
