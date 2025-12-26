package month

import (
	"fmt"
	"strings"
)

var monthNames = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

// Month converts a month number (1-12) to its corresponding month name.
func Month(n int) (string, error) {
	if n < 1 || n > 12 {
		return "", fmt.Errorf("invalid month number: %d", n)
	}

	return monthNames[n-1], nil
}

// ExtractMonth converts a month name to its corresponding month number (1-12).
func ExtractMonth(str string) (int, error) {
	str = strings.ToLower(str)

	for i, name := range monthNames {
		name = strings.ToLower(name)

		if name == str {
			return i + 1, nil
		}

		for l := 3; l <= 5 && l <= len(name); l++ {
			if name[:l] == str || fmt.Sprintf("%s.", name[:l]) == str {
				return i + 1, nil
			}
		}
	}

	return 0, fmt.Errorf("invalid month name: %s", str)
}
