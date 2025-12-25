package date_test

import (
	"testing"

	. "github.com/elecbug/irregular/en/date"
)

func TestMonth(t *testing.T) {
	tests := []struct {
		input    int
		expected string
		hasError bool
	}{
		{1, "January", false},
		{2, "February", false},
		{3, "March", false},
		{4, "April", false},
		{5, "May", false},
		{6, "June", false},
		{7, "July", false},
		{8, "August", false},
		{9, "September", false},
		{10, "October", false},
		{11, "November", false},
		{12, "December", false},
		{0, "", true},
		{13, "", true},
	}

	for _, test := range tests {
		result, err := Month(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Month(%d) unexpected error state: got %v, expected error: %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("Month(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestExtractMonth(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"January", 1, false},
		{"Feb", 2, false},
		{"Mar.", 3, false},
		{"April", 4, false},
		{"May", 5, false},
		{"Jun", 6, false},
		{"July", 7, false},
		{"Aug.", 8, false},
		{"September", 9, false},
		{"Oct", 10, false},
		{"Nov.", 11, false},
		{"December", 12, false},
		{"Januar", 0, true},
		{"Foo", 0, true},
	}

	for _, test := range tests {
		result, err := ExtractMonth(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ExtractMonth(%s) unexpected error state: got %v, expected error: %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("ExtractMonth(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
