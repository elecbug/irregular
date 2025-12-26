package date

// LAST_DAYS_IN_MONTH maps month numbers to the number of days in that month (non-leap year).
var LAST_DAYS_IN_MONTH = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// dateComponents holds the parsed components of a date string.
type dateComponents struct {
	first            int
	second           int
	third            int
	separator        string
	hasLastSeparator bool
}
