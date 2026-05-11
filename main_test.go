package main

import (
	"strings"
	"testing"
)

// helper: count occurrences of a substring in s.
func countOccurrences(s, sub string) int {
	return strings.Count(s, sub)
}

// TestGenerateArt_EmptyInputProducesNoOutput checks that an empty string
// returns "" and not 8 blank lines.
func TestGenerateArt_EmptyInputProducesNoOutput(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt("", banner)
	if got != "" {
		t.Errorf("expected empty output for empty input, got %q", got)
	}
}

// TestGenerateArt_SingleNewlineProducesOneLine checks that a lone \n
// produces exactly one newline in the output — not 8 blank lines.
func TestGenerateArt_SingleNewlineProducesOneLine(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt(`\n`, banner)
	if got != "\n" {
		t.Errorf("expected exactly one newline for input %q, got %q", `\n`, got)
	}
}

// TestGenerateArt_SingleWordProducesEightLines checks that a plain word
// produces exactly 8 output lines.
func TestGenerateArt_SingleWordProducesEightLines(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt("Hi", banner)
	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines for 'Hi', got %d:\n%s", len(lines), got)
	}
}

// TestGenerateArt_TwoWordsProducesSixteenLines checks that two words split by
// \n together produce 16 lines (8 each), no blank line between them.
func TestGenerateArt_TwoWordsProducesSixteenLines(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt(`A\nB`, banner)
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 16 {
		t.Errorf("expected 16 newlines for 'A\\nB', got %d\noutput:\n%s", newlineCount, got)
	}
}

// TestGenerateArt_DoubleNewlineProducesBlankLineBetween is the most critical
// test: \n\n must produce one blank line between the two word blocks —
// NOT 8 blank lines.
func TestGenerateArt_DoubleNewlineProducesBlankLineBetween(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt(`A\n\nB`, banner)
	// Expected: 8 lines for A + 1 blank line + 8 lines for B = 17 newlines
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 17 {
		t.Errorf("expected 17 newlines for 'A\\n\\nB' (8+1+8), got %d\noutput:\n%s",
			newlineCount, got)
	}
}

// TestGenerateArt_TrailingNewlineAddsEightBlankLines checks that a trailing
// \n after a word adds 8 blank lines (the empty segment rendered as 8 blank).
// NOTE: per the spec, "Hello\n" produces Hello's 8 lines THEN 8 blank lines.
func TestGenerateArt_TrailingNewlineAddsEightBlankLines(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt(`Hello\n`, banner)
	// 8 lines for Hello + 8 blank lines = 16 newlines total
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 9 {
		t.Errorf("expected 9 newlines for 'Hello\\n', got %d\noutput:\n%s",
			newlineCount, got)
	}
}

// TestGenerateArt_LeadingNewlineAddsBlankLineFirst checks that a leading \n
// produces a blank line before the word block.
func TestGenerateArt_LeadingNewlineAddsBlankLineFirst(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt(`\nHello`, banner)
	// 1 blank line + 8 lines for Hello = 9 newlines
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 9 {
		t.Errorf("expected 9 newlines for '\\nHello', got %d\noutput:\n%s",
			newlineCount, got)
	}
}

// TestGenerateArt_EachLineEndsWithNewline checks that every line in the output
// ends with \n (required by the spec — verified by cat -e showing $).
func TestGenerateArt_EachLineEndsWithNewline(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt("Hi", banner)
	if !strings.HasSuffix(got, "\n") {
		t.Error("output must end with a newline")
	}
	// Every line separator must be \n not \r\n
	if strings.Contains(got, "\r\n") {
		t.Error("output contains \\r\\n — use Unix line endings only")
	}
}

// TestGenerateArt_ContentMatchesRenderLine checks that the output of
// GenerateArt for a single word matches manually joining RenderLine output.
func TestGenerateArt_ContentMatchesRenderLine(t *testing.T) {
	banner := loadStandard(t)

	rendered := RenderLine("Hello", banner)
	var want strings.Builder
	for _, line := range rendered {
		want.WriteString(line + "\n")
	}

	got := GenerateArt("Hello", banner)
	if got != want.String() {
		t.Errorf("GenerateArt(\"Hello\") does not match manually joined RenderLine output\ngot:\n%s\nwant:\n%s",
			got, want.String())
	}
}

// TestGenerateArt_SpaceOnlyInput checks that a string of spaces renders
// correctly — spaces must not be silently dropped.
func TestGenerateArt_SpaceOnlyInput(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt("   ", banner)
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 8 {
		t.Errorf("expected 8 lines for 3 spaces, got %d newlines\noutput:\n%q", newlineCount, got)
	}
}

// TestGenerateArt_NumbersAndLetters checks that mixed numeric and alphabetic
// input renders without error and produces 8 lines.
func TestGenerateArt_NumbersAndLetters(t *testing.T) {
	banner := loadStandard(t)
	got := GenerateArt("1Hello 2There", banner)
	newlineCount := countOccurrences(got, "\n")
	if newlineCount != 8 {
		t.Errorf("expected 8 lines for '1Hello 2There', got %d\noutput:\n%s",
			newlineCount, got)
	}
}

func loadStandard(t *testing.T) map[rune][]string {
	t.Helper()

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	return banner
}

// TestValidateInput_ValidStrings checks that fully valid inputs return 0, nil.
func TestValidateInput_ValidStrings(t *testing.T) {
	valid := []string{
		"Hello",
		"hello world",
		"Hello, World!",
		"1234567890",
		"~`!@#$%^&*()-_=+[]{}|;':\",./<>?",
		" ",
		"",
	}
	for _, s := range valid {
		r, err := ValidateInput(s)
		if err != nil {
			t.Errorf("ValidateInput(%q): unexpected error %v (rune %q)", s, err, r)
		}
		if r != 0 {
			t.Errorf("ValidateInput(%q): expected rune 0, got %q", s, r)
		}
	}
}

// TestValidateInput_InvalidCharacters checks that each non-ASCII or
// below-32 character is caught and returned.
func TestValidateInput_InvalidCharacters(t *testing.T) {
	cases := []struct {
		input       string
		invalidRune rune
	}{
		{"héllo", 'é'},
		{"naïve", 'ï'},
		{"café", 'é'},
		{"こんにちは", 'こ'},
		{"\x01hello", '\x01'},
		{"\x00", '\x00'},
		{"\x1F", '\x1F'},      // ASCII 31 — just below valid range
		{"hello\x7F", '\x7F'}, // ASCII 127 — just above valid range
		{"😀", '😀'},
	}
	for _, tc := range cases {
		r, err := ValidateInput(tc.input)
		if err == nil {
			t.Errorf("ValidateInput(%q): expected error, got nil", tc.input)
			continue
		}
		if r != tc.invalidRune {
			t.Errorf("ValidateInput(%q): expected rune %q, got %q", tc.input, tc.invalidRune, r)
		}
	}
}

// TestValidateInput_ReturnsFirstInvalidOnly checks that when there are
// multiple invalid characters, only the first one is returned.
func TestValidateInput_ReturnsFirstInvalidOnly(t *testing.T) {
	// 'é' comes before 'ñ' — must return 'é'
	r, err := ValidateInput("héñ")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if r != 'é' {
		t.Errorf("expected first invalid rune 'é', got %q", r)
	}
}

// TestValidateInput_BoundaryASCII32And126 checks that ASCII codes 32 (space)
// and 126 (~) are accepted — they are the inclusive edges of the valid range.
func TestValidateInput_BoundaryASCII32And126(t *testing.T) {
	boundaries := []rune{32, 126}
	for _, code := range boundaries {
		r, err := ValidateInput(string(code))
		if err != nil {
			t.Errorf("ASCII %d (%q) should be valid, got error: %v (rune %q)", code, code, err, r)
		}
	}
}

// TestValidateInput_BoundaryASCII31And127 checks that ASCII codes 31 and 127
// are rejected — they sit just outside the valid range.
func TestValidateInput_BoundaryASCII31And127(t *testing.T) {
	for _, code := range []rune{31, 127} {
		_, err := ValidateInput(string(code))
		if err == nil {
			t.Errorf("ASCII %d should be invalid, but got no error", code)
		}
	}
}

// TestValidateInput_ValidPrefixBeforeInvalid checks that even if most of the
// string is valid, the function still catches an invalid char at the end.
func TestValidateInput_ValidPrefixBeforeInvalid(t *testing.T) {
	_, err := ValidateInput("Hello World é")
	if err == nil {
		t.Error("expected error for string with invalid char at end, got nil")
	}
}

// TestValidateInput_EmptyString expects 0, nil for empty input.
func TestValidateInput_EmptyString(t *testing.T) {
	r, err := ValidateInput("")
	if err != nil {
		t.Errorf("ValidateInput(\"\"): unexpected error %v", err)
	}
	if r != 0 {
		t.Errorf("ValidateInput(\"\"): expected rune 0, got %q", r)
	}
}
