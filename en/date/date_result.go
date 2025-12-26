package date

// DateResult represents the result of parsing a date string.
type DateResult struct {
	Format     DateFormat
	Year       int
	Month      int
	Day        int
	Confidence float64
}
