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
			titles := []string{"", "Ⅰ", "Ⅱ", "Ⅲ", "Ⅳ", "Ⅴ", "Ⅵ", "Ⅶ", "Ⅷ", "Ⅸ", "Ⅹ"}
			return titles[n], nil
		} else {
			titles := []string{"", "ⅰ", "ⅱ", "ⅲ", "ⅳ", "ⅴ", "ⅵ", "ⅶ", "ⅷ", "ⅸ", "ⅹ"}
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
