package number_test

import (
	"testing"

	. "github.com/elecbug/irregular/en/number"
)

func TestOrdinalNum(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1st"},
		{2, "2nd"},
		{3, "3rd"},
		{4, "4th"},
		{11, "11th"},
		{12, "12th"},
		{13, "13th"},
		{21, "21st"},
		{22, "22nd"},
		{23, "23rd"},
		{24, "24th"},
		{51, "51st"},
		{52, "52nd"},
		{53, "53rd"},
		{54, "54th"},
		{101, "101st"},
		{111, "111th"},
	}

	for _, test := range tests {
		result := OrdinalNum(test.input)
		if result != test.expected {
			t.Errorf("OrdinalNum(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestValidOrdinalNum(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"1st", 1, false},
		{"2nd", 2, false},
		{"3rd", 3, false},
		{"4th", 4, false},
		{"11th", 11, false},
		{"12th", 12, false},
		{"13th", 13, false},
		{"21st", 21, false},
		{"22nd", 22, false},
		{"23rd", 23, false},
		{"24th", 24, false},
		{"51st", 51, false},
		{"52nd", 52, false},
		{"53rd", 53, false},
		{"54th", 54, false},
		{"101st", 101, false},
		{"111th", 111, false},
		{"5st", 0, true},
		{"2rd", 0, true},
		{"13st", 0, true},
		{"abc", 0, true},
	}

	for _, test := range tests {
		result, err := ValidateOrdinalNum(test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("ValidOrdinalNum(%s) expected error; got %d", test.input, result)
			}
		} else {
			if err != nil {
				t.Errorf("ValidOrdinalNum(%s) unexpected error: %v", test.input, err)
			} else if result != test.expected {
				t.Errorf("ValidOrdinalNum(%s) = %d; expected %d", test.input, result, test.expected)
			}
		}
	}
}
