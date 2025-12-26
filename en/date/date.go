package date

import (
	"fmt"
	"strings"
	"unicode"
)

// Date represents a date with its raw string representation.
type Date struct {
	raw string
}

// NewDate creates a new Date instance from the given date string.
func NewDate(dateStr string) Date {
	return Date{raw: dateStr}
}

// Raw returns the raw date string.
func (d Date) Raw() string {
	return d.raw
}

// Parse parses the date string and returns possible date interpretations based on the provided configuration.
func (d Date) Parse(cfg *Config) ([]DateResult, error) {
	if cfg == nil {
		defaultCfg := DefaultConfig()
		cfg = defaultCfg
	}

	parsed, err := d.determineDatePattern(*cfg)
	if err != nil {
		return nil, err
	}

	out := make([]DateResult, 0, len(parsed))

	for _, r := range parsed {
		if r.Format == FormatInvalid {
			continue
		}

		if r.Month < 1 || r.Month > 12 {
			continue
		}

		maxDay := LAST_DAYS_IN_MONTH[r.Month]
		if r.Month == 2 && ((r.Year%4 == 0 && r.Year%100 != 0) || (r.Year%400 == 0)) {
			maxDay = 29
		}

		if r.Day < 1 || r.Day > maxDay {
			continue
		}

		out = append(out, r)
	}

	if len(out) == 0 {
		return []DateResult{{Format: FormatInvalid, Confidence: 1}}, nil
	}

	return out, nil
}

// determineDatePattern parses the date string into its components and determines possible date formats.
func (d Date) determineDatePattern(cfg Config) ([]DateResult, error) {
	components, err := parseDateComponents(d.raw)

	if err != nil {
		return nil, err
	}

	if components.first > 12 && components.second > 12 && components.third > 12 {
		return []DateResult{
			{
				Format:     FormatInvalid,
				Confidence: 1,
			},
		}, nil
	}

	if components.first < 100 && components.second < 100 && components.third < 100 {
		// YY MM DD or DD MM YY or MM DD YY
		if components.first > 31 {
			// YY MM DD
			return []DateResult{
				{
					Format:     FormatYMD,
					Year:       yearConvert(components.first, cfg.BaseYear),
					Month:      components.second,
					Day:        components.third,
					Confidence: 1,
				},
			}, nil
		} else if components.third > 31 {
			if components.first > 12 {
				// DD MM YY
				return []DateResult{
					{
						Format:     FormatDMY,
						Year:       yearConvert(components.third, cfg.BaseYear),
						Month:      components.second,
						Day:        components.first,
						Confidence: 1,
					},
				}, nil
			} else if components.second > 12 {
				// MM DD YY
				return []DateResult{
					{
						Format:     FormatMDY,
						Year:       yearConvert(components.third, cfg.BaseYear),
						Month:      components.first,
						Day:        components.second,
						Confidence: 1,
					},
				}, nil
			} else {
				// DD MM YY or MM DD YY
				return []DateResult{
					{
						Format:     FormatDMY,
						Year:       yearConvert(components.third, cfg.BaseYear),
						Month:      components.second,
						Day:        components.first,
						Confidence: 0.5,
					},
					{
						Format:     FormatMDY,
						Year:       yearConvert(components.third, cfg.BaseYear),
						Month:      components.first,
						Day:        components.second,
						Confidence: 0.5,
					},
				}, nil
			}
		} else if components.second > 31 {
			// Invalid date
			return []DateResult{
				{
					Format:     FormatInvalid,
					Confidence: 1,
				},
			}, nil
		} else {
			// DD MM YY or MM DD YY or YY MM DD
			return []DateResult{
				{
					Format:     FormatDMY,
					Year:       yearConvert(components.third, cfg.BaseYear),
					Month:      components.second,
					Day:        components.first,
					Confidence: 1.0 / 3.0,
				},
				{
					Format:     FormatMDY,
					Year:       yearConvert(components.third, cfg.BaseYear),
					Month:      components.first,
					Day:        components.second,
					Confidence: 1.0 / 3.0,
				},
				{
					Format:     FormatYMD,
					Year:       yearConvert(components.first, cfg.BaseYear),
					Month:      components.second,
					Day:        components.third,
					Confidence: 1.0 / 3.0,
				},
			}, nil
		}
	} else if components.first >= 100 {
		// YYYY MM DD
		return []DateResult{
			{
				Format:     FormatYMD,
				Year:       components.first,
				Month:      components.second,
				Day:        components.third,
				Confidence: 1,
			},
		}, nil
	} else if components.third >= 100 {
		if components.first > 12 {
			// DD MM YYYY
			return []DateResult{
				{
					Format:     FormatDMY,
					Year:       components.third,
					Month:      components.second,
					Day:        components.first,
					Confidence: 1,
				},
			}, nil
		} else if components.second > 12 {
			// MM DD YYYY
			return []DateResult{
				{
					Format:     FormatMDY,
					Year:       components.third,
					Month:      components.first,
					Day:        components.second,
					Confidence: 1,
				},
			}, nil
		} else {
			// DD MM YYYY or MM DD YYYY
			return []DateResult{
				{
					Format:     FormatDMY,
					Year:       components.third,
					Month:      components.second,
					Day:        components.first,
					Confidence: 0.5,
				},
				{
					Format:     FormatMDY,
					Year:       components.third,
					Month:      components.first,
					Day:        components.second,
					Confidence: 0.5,
				},
			}, nil
		}
	} else {
		return []DateResult{
			{
				Format:     FormatInvalid,
				Confidence: 1,
			},
		}, nil
	}
}

// yearConvert converts a 2-digit year to a 4-digit year based on the base year.
func yearConvert(year int, baseYear int) int {
	if year >= 100 {
		return year
	}

	century := baseYear / 100
	fullYear := century*100 + year

	if fullYear < baseYear {
		fullYear += 100
	}

	return fullYear
}

// parseDateComponents parses the date string into its numeric components and separators.
func parseDateComponents(dateStr string) (dateComponents, error) {
	s := strings.TrimSpace(dateStr)
	if s == "" {
		return dateComponents{}, fmt.Errorf("empty date string")
	}

	var nums []int
	var seps []rune

	i := 0
	for i < len(s) {
		if !unicode.IsDigit(rune(s[i])) {
			return dateComponents{}, fmt.Errorf("expected digit at position %d", i)
		}

		val := 0
		for i < len(s) && unicode.IsDigit(rune(s[i])) {
			val = val*10 + int(s[i]-'0')
			i++
		}
		nums = append(nums, val)

		if i >= len(s) {
			break
		}

		if unicode.IsDigit(rune(s[i])) {
			return dateComponents{}, fmt.Errorf("expected separator at position %d", i)
		}

		sep := rune(s[i])
		seps = append(seps, sep)
		i++
	}

	if len(nums) != 3 {
		return dateComponents{}, fmt.Errorf("expected 3 numeric components, got %d", len(nums))
	}

	if len(seps) < 2 || len(seps) > 3 {
		return dateComponents{}, fmt.Errorf("invalid separator count: %d", len(seps))
	}

	baseSep := seps[0]
	for _, s := range seps {
		if s != baseSep {
			return dateComponents{}, fmt.Errorf("mixed separators are not allowed")
		}
	}

	if len(seps) == 2 {
		return dateComponents{first: nums[0], second: nums[1], third: nums[2], separator: string(baseSep), hasLastSeparator: false}, nil
	} else if len(seps) == 3 {
		return dateComponents{first: nums[0], second: nums[1], third: nums[2], separator: string(baseSep), hasLastSeparator: true}, nil
	}

	return dateComponents{}, fmt.Errorf("unexpected error in parsing date components")
}
