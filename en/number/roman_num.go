package number

import "fmt"

// RomanNum converts an integer to its Roman numeral representation.
func RomanNum(n int) (string, error) {
	if n <= 0 || n >= 4000 {
		return "", fmt.Errorf("number out of range (1-3999): %d", n)
	}

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(vals); i++ {
		for n >= vals[i] {
			n -= vals[i]
			result += syms[i]
		}
	}

	return result, nil
}

// TitleRomanNum converts an integer to its Roman numeral representation
func TitleRomanNum(n int, upper bool) (string, error) {
	if roman, err := RomanNum(n); err != nil {
		return roman, err
	} else {
		if n < 0 || n > 10 {
			return roman, nil
		}

		if upper {
			titles := []string{"", "\u2160", "\u2161", "\u2162", "\u2163", "\u2164", "\u2165", "\u2166", "\u2167", "\u2168", "\u2169"}
			return titles[n], nil
		} else {
			titles := []string{"", "\u2170", "\u2171", "\u2172", "\u2173", "\u2174", "\u2175", "\u2176", "\u2177", "\u2178", "\u2179"}
			return titles[n], nil
		}
	}
}

// ValidRomanNum parses a Roman numeral string and returns the corresponding integer.
func ValidRomanNum(s string) (int, error) {
	vals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		val, ok := vals[s[i]]
		if !ok {
			return 0, fmt.Errorf("invalid Roman numeral character: %c", s[i])
		}

		if val < prev {
			n -= val
		} else {
			n += val
		}
		prev = val
	}

	reconstructed, err := RomanNum(n)
	if err != nil || reconstructed != s {
		return 0, fmt.Errorf("invalid Roman numeral format: %s", s)
	}

	return n, nil
}
