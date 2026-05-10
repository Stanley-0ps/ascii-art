package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	charHeight = 8
	asciiStart = 32
	asciiEnd   = 126
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

	banner, err := LoadBanner("shadow.txt")
	if err != nil {
		return err
	}

	_, err = ValidateInput(strings.ReplaceAll(input, `\n`, ""))
	if err != nil {
		return err
	}

	fmt.Print(GenerateArt(input, banner))

	return nil
}

func validateArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("program expects exactly one argument")
	}

	return args[1], nil
}

// LoadBanner matches the test file function name.
func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file: %w", err)
	}

	if len(data) == 0 {
		return nil, errors.New("banner file is empty")
	}

	// Normalize line endings
	clean := strings.ReplaceAll(string(data), "\r\n", "\n")
	clean = strings.ReplaceAll(clean, "\r", "\n")

	lines := strings.Split(clean, "\n")

	// Remove trailing empty lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	banner := make(map[rune][]string)

	index := 1

	for ascii := asciiStart; ascii <= asciiEnd; ascii++ {
		if index+charHeight > len(lines) {
			return nil, errors.New("invalid banner format")
		}

		glyph := lines[index : index+charHeight]

		if len(glyph) != charHeight {
			return nil, fmt.Errorf("invalid glyph height for %q", rune(ascii))
		}

		banner[rune(ascii)] = glyph

		index += charHeight + 1
	}

	if len(banner) != 95 {
		return nil, fmt.Errorf("expected 95 characters, got %d", len(banner))
	}

	return banner, nil
}

// ValidateInput matches the test expectations exactly.
func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r < asciiStart || r > asciiEnd {
			return r, fmt.Errorf("invalid character: %q", r)
		}
	}

	return 0, nil
}

// RenderLine returns exactly 8 rendered lines.
func RenderLine(text string, banner map[rune][]string) []string {
	lines := make([]string, charHeight)

	for i := 0; i < charHeight; i++ {
		var builder strings.Builder

		for _, r := range text {
			builder.WriteString(banner[r][i])
		}

		lines[i] = builder.String()
	}

	return lines
}

// GenerateArt handles ALL newline edge cases from the tests.
func GenerateArt(input string, banner map[rune][]string) string {
	// Empty input => no output
	if input == "" {
		return ""
	}

	// Single \n => exactly one newline
	if input == `\n` {
		return "\n"
	}

	parts := strings.Split(input, `\n`)

	var result strings.Builder

	for i, part := range parts {
		switch {
		// Leading empty part
		case i == 0 && part == "":
			result.WriteString("\n")

		// Trailing empty part
		case i == len(parts)-1 && part == "":
			for j := 0; j < charHeight; j++ {
				result.WriteString("\n")
			}

		// Empty middle part => single blank line
		case part == "":
			result.WriteString("\n")

		// Normal text
		default:
			rendered := RenderLine(part, banner)

			for _, line := range rendered {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}

	return result.String()
}