package number

import (
	"fmt"
	"strings"
)

// AlphabetNum converts an integer to its corresponding alphabetic representation.
func AlphabetNum(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("input must be a positive integer")
	}

	words := []string{}
	for n > 0 {
		alpha := n % 26

		if alpha == 0 {
			alpha = 26
		}
		words = append([]string{digitToWord(alpha)}, words...)
		n = (n - alpha) / 26
	}

	return joinWords(words), nil
}

// ExtractAlphabetNum converts an alphabetic representation back to its integer form.
func ExtractAlphabetNum(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("input string is empty")
	}

	s = strings.ToUpper(s)

	n := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char < 'A' || char > 'Z' {
			return 0, fmt.Errorf("invalid character in alphabetic representation: %c", char)
		}
		n = n*26 + int(char-'A'+1)
	}
	return n, nil
}

// joinWords joins a slice of words into a single string with spaces.
func joinWords(words []string) string {
	result := ""
	for _, word := range words {
		result += word
	}
	return result
}

// digitToWord converts a digit (1-26) to its corresponding alphabet letter.
func digitToWord(digit int) string {
	if digit < 1 || digit > 26 {
		return ""
	}
	return string('A' + rune(digit) - 1)
}
