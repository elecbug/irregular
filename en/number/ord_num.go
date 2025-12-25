package number

import "fmt"

// OrdinalNum converts an integer to its ordinal string representation.
func OrdinalNum(n int) string {
	if n%100 >= 11 && n%100 <= 13 {
		return fmt.Sprintf("%dth", n)
	}

	switch n % 10 {
	case 1:
		return fmt.Sprintf("%dst", n)
	case 2:
		return fmt.Sprintf("%dnd", n)
	case 3:
		return fmt.Sprintf("%drd", n)
	default:
		return fmt.Sprintf("%dth", n)
	}
}

// ValidateOrdinalNum parses an ordinal string and returns the corresponding integer.
func ValidateOrdinalNum(s string) (int, error) {
	var n int
	var suffix string

	_, err := fmt.Sscanf(s, "%d%s", &n, &suffix)
	if err != nil {
		return 0, fmt.Errorf("invalid ordinal number format: %s", s)
	}

	expectedSuffix := OrdinalNum(n)[len(fmt.Sprintf("%d", n)):]
	if suffix != expectedSuffix {
		return 0, fmt.Errorf("invalid ordinal suffix for number %d: got %s, expected %s", n, suffix, expectedSuffix)
	}

	return n, nil
}
