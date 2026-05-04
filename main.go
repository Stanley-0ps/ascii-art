package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run() error {
	input, err := validateArgs(os.Args)
	if err != nil {
		return err
	}

	font, err := loadFont("shadow.txt")
	if err != nil {
		return err
	}

	return render(input, font)
}

func validateArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("program expects exactly one argument")
	}
	return args[1], nil
}

func loadFont(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	font := make(map[rune][]string)

	ascii := 32
	index := 1

	for ascii <= 126 {
		if index+8 > len(lines) {
			return nil, errors.New("invalid font file format")
		}
		font[rune(ascii)] = lines[index : index+8]
		index += 9
		ascii++
	}

	return font, nil
}

func render(input string, font map[rune][]string) error {
	if input == "" {
		return nil
	}

	if input == "\\n" {
		fmt.Println()
		return nil
	}

	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		if err := renderWord(word, font); err != nil {
			return err
		}
	}

	return nil
}

func renderWord(word string, font map[rune][]string) error {
	for i := 0; i < 8; i++ {
		for _, ch := range word {
			art, ok := font[ch]
			if !ok {
				return fmt.Errorf("unsupported character: %q", ch)
			}
			fmt.Print(art[i])
		}
		fmt.Println()
	}
	return nil
}