package date

// DateFormat represents the format of the date.
type DateFormat int

const (
	FormatInvalid DateFormat = iota
	FormatDMY
	FormatMDY
	FormatYMD
)

// String returns the string representation of the DateFormat.
func (df DateFormat) String() string {
	switch df {
	case FormatDMY:
		return "DMY"
	case FormatMDY:
		return "MDY"
	case FormatYMD:
		return "YMD"
	default:
		return "Invalid Format"
	}
}
