package datediff

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var formatUnits = map[string]string{
	"%Y": "year",
	"%y": "year",
	"%M": "month",
	"%m": "month",
	"%W": "week",
	"%w": "week",
	"%D": "day",
	"%d": "day",
}

// format formats dates difference according to the provided format.
// It trims time units with 0 values.
func format(diff Diff, rawFormat string) string {
	result := rawFormat

	frmt(diff, rawFormat, func(n int, verb, unit string) {
		if n == 0 {
			result = zeroVerbReplace(result, verb)
		} else {
			result = verbReplace(result, n, verb, unit)
		}
	})

	return result
}

// format formats dates difference according to the provided format.
// Since this function is private, it's assumed that format is valid.
func formatWithZeros(diff Diff, rawFormat string) string {
	result := rawFormat

	frmt(diff, rawFormat, func(n int, verb, unit string) {
		result = verbReplace(result, n, verb, unit)
	})

	return result
}

func frmt(diff Diff, rawFormat string, replace func(n int, verb, unit string)) {
	for verb, unit := range formatUnits {
		if strings.Contains(rawFormat, verb) {
			var n int
			switch unit {
			case "year":
				n = diff.Years
			case "month":
				n = diff.Months
			case "week":
				n = diff.Weeks
			case "day":
				n = diff.Days
			}
			replace(n, verb, unit)
		}
	}
}

// formatNoun takes a positive number n and noun s in singular form.
// It returns a number and correct form of noun (singular or plural).
func formatNoun(n int, s string) string {
	f := "%d %s"
	if n != 1 {
		f += "s"
	}
	return fmt.Sprintf(f, n, s)
}

func verbReplace(s string, n int, verb, unit string) string {
	replacement := strconv.Itoa(n)
	if r := rune(verb[1]); unicode.IsUpper(r) {
		replacement = formatNoun(n, unit)
	}
	return strings.ReplaceAll(s, verb, replacement)
}

func zeroVerbReplace(s, verb string) string {
	s = strings.ReplaceAll(s, " "+verb, "")
	return strings.ReplaceAll(s, verb, "")
}
