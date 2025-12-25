package number_test

import (
	"testing"

	. "github.com/elecbug/irregular/en/number"
)

func TestAlphabetNum(t *testing.T) {
	tests := []struct {
		input    int
		expected string
		hasError bool
	}{
		{1, "A", false},
		{26, "Z", false},
		{27, "AA", false},
		{52, "AZ", false},
		{703, "AAA", false},
		{0, "", true},
		{-5, "", true},
	}

	for _, test := range tests {
		result, err := AlphabetNum(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("AlphabetNum(%d) unexpected error state: got %v, want error: %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("AlphabetNum(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestExtractAlphabetNum(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"A", 1, false},
		{"Z", 26, false},
		{"AA", 27, false},
		{"AZ", 52, false},
		{"AAA", 703, false},
		{"A1", 0, true},
		{"a", 1, false},
		{"", 0, true},
	}

	for _, test := range tests {
		result, err := ExtractAlphabetNum(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ExtractAlphabetNum(%q) unexpected error state: got %v, want error: %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("ExtractAlphabetNum(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
