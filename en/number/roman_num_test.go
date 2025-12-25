package number_test

import (
	"testing"

	. "github.com/elecbug/irregular/en/number"
)

func TestRomanNum(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1987, "MCMLXXXVII"},
	}

	for _, test := range tests {
		result, err := RomanNum(test.input)
		if err != nil {
			t.Errorf("RomanNum(%d) returned an error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("RomanNum(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestValidRomanNum(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"I", 1, false},
		{"IV", 4, false},
		{"V", 5, false},
		{"IX", 9, false},
		{"X", 10, false},
		{"XL", 40, false},
		{"L", 50, false},
		{"XC", 90, false},
		{"C", 100, false},
		{"CD", 400, false},
		{"D", 500, false},
		{"CM", 900, false},
		{"M", 1000, false},
		{"MCMLXXXVII", 1987, false},
		{"INVALID", 0, true},
	}

	for _, test := range tests {
		result, err := ValidRomanNum(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("ValidRomanNum(%s) unexpected error state: got %v, expected error: %v", test.input, err, test.expectError)
			continue
		}
		if result != test.expected {
			t.Errorf("ValidRomanNum(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
