package date_test

import (
	"sort"
	"testing"

	. "github.com/elecbug/irregular/en/date"
)

func TestDateRaw(t *testing.T) {
	dateStr := "12/31/2020"
	date := NewDate(dateStr)

	if date.Raw() != dateStr {
		t.Errorf("Date.Raw() = %s; expected %s", date.Raw(), dateStr)
	}
}

func TestDateParse(t *testing.T) {
	dateStrs := []string{
		"2020-12-31",
		"31/12/2020",
		"12/31/2020",
		"99/99/9999",
		"03-04-05",
		"03.04.05",
		"03.04.05.",
		"97/08/09",
		"08/09/97",
		"09-97-08",
		"29/02/2020",
		"29/02/2019",
		"31/04/2021",
	}

	expectedFormats := [][]DateFormat{
		{FormatYMD},
		{FormatDMY},
		{FormatMDY},
		{FormatInvalid},
		{FormatYMD, FormatDMY, FormatMDY},
		{FormatYMD, FormatDMY, FormatMDY},
		{FormatYMD, FormatDMY, FormatMDY},
		{FormatYMD},
		{FormatDMY, FormatMDY},
		{FormatInvalid},
		{FormatDMY},
		{FormatInvalid},
		{FormatInvalid},
	}

	for i, dateStr := range dateStrs {
		date := NewDate(dateStr)
		results, err := date.Parse(&Config{BaseYear: 1970})

		if err != nil {
			t.Errorf("Date.Parse(%s) unexpected error: %v", dateStr, err)
			continue
		}

		if len(results) != len(expectedFormats[i]) {
			t.Errorf("Date.Parse(%s) = %d results; expected %d", dateStr, len(results), len(expectedFormats[i]))
			continue
		}

		sort.Slice(results, func(a, b int) bool {
			return results[a].Format < results[b].Format
		})
		sort.Slice(expectedFormats[i], func(a, b int) bool {
			return expectedFormats[i][a] < expectedFormats[i][b]
		})

		for j, result := range results {
			if result.Format != expectedFormats[i][j] {
				t.Errorf("Date.Parse(%s) result %d format = %s; expected %s", dateStr, j, result.Format.String(), expectedFormats[i][j].String())
			}
		}
	}
}
