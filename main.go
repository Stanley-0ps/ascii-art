package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	charHeight = 8
	charBlock  = 9
	firstASCII = 32
	lastASCII  = 126
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
	if err := ValidateBanner(banner); err != nil {
		fmt.Println("Error:", err)
		return
	}
	PrintAscii(input, banner)
}

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file: %w", err)
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines, nil
}

func ValidateBanner(banner []string) error {
	expectedChars := lastASCII - firstASCII + 1
	expectedLines := expectedChars * charBlock
	if len(banner) < expectedLines {
		return fmt.Errorf("invalid banner: too few lines (%d < %d)", len(banner), expectedLines)
	}
	return nil
}

func PrintAscii(input string, banner []string) {
	if input == "" {
		fmt.Println()
		return
	}
	parts := strings.Split(input, "\\n")

	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < charHeight; row++ {
			var line strings.Builder

			for _, r := range part {
				if !IsPrintable(r) {
					line.WriteString("?")
					continue
				}
				offset := int(r) - firstASCII
				base := offset * charBlock
				index := base + row + 1
				if index < 0 || index >= len(banner) {
					line.WriteString("?")
					continue
				}

				line.WriteString(banner[index])
			}
			fmt.Println(line.String())
		}
	}
}

func IsPrintable(r rune) bool {
	return r >= firstASCII && r <= lastASCII
}
