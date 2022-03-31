package datediff_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/antklim/datediff"
)

const dateFmt = "2006-01-02"

// these are the fields of testdata/dates.csv
const (
	startFld = iota
	endFld
	modeFld
	formatFld
	yearsFld
	monthsFld
	weeksFld
	daysFld
	printFld
	printWzerosFld
)

type datediffRecord struct {
	start          time.Time
	end            time.Time
	mode           datediff.DiffMode
	format         string
	diff           datediff.Diff
	print          string
	printWithZeros string
}

func loadDatediffRecord(r []string) (datediffRecord, error) {
	start, err := time.Parse(dateFmt, r[startFld])
	if err != nil {
		return datediffRecord{}, err
	}
	end, err := time.Parse(dateFmt, r[endFld])
	if err != nil {
		return datediffRecord{}, err
	}
	mode, err := strconv.Atoi(r[modeFld])
	if err != nil {
		return datediffRecord{}, err
	}
	years, err := strconv.Atoi(r[yearsFld])
	if err != nil {
		return datediffRecord{}, err
	}
	months, err := strconv.Atoi(r[monthsFld])
	if err != nil {
		return datediffRecord{}, err
	}
	weeks, err := strconv.Atoi(r[weeksFld])
	if err != nil {
		return datediffRecord{}, err
	}
	days, err := strconv.Atoi(r[daysFld])
	if err != nil {
		return datediffRecord{}, err
	}

	return datediffRecord{
		start:          start,
		end:            end,
		mode:           datediff.DiffMode(mode),
		format:         r[formatFld],
		diff:           datediff.Diff{Years: years, Months: months, Weeks: weeks, Days: days},
		print:          r[printFld],
		printWithZeros: r[printWzerosFld],
	}, nil
}

func loadDatediffRecordsForTest() ([]datediffRecord, error) {
	f, err := os.Open("testdata/datediff.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comment = '#'
	rawRecords, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var dr []datediffRecord
	for _, rr := range rawRecords {
		r, err := loadDatediffRecord(rr)
		if err != nil {
			return nil, err
		}
		dr = append(dr, r)
	}
	return dr, nil
}

var testInvalidFormat = []struct {
	format   string
	expected string
}{
	{
		format:   "%X%L %S",
		expected: `format "%X%L %S" has unknown verb X`,
	},
	{
		format:   "   ",
		expected: "undefined dates difference mode",
	},
	{
		format:   "Years and months",
		expected: "undefined dates difference mode",
	},
}

func TestNewDiff(t *testing.T) {
	testCases, err := loadDatediffRecordsForTest()
	if err != nil {
		t.Fatal(err)
	}

	for _, tC := range testCases {
		desc := fmt.Sprintf("NewDiff(%s, %s, %s)", tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.format)
		t.Run(desc, func(t *testing.T) {
			got, err := datediff.NewDiff(tC.start, tC.end, tC.format)
			if err != nil {
				t.Errorf("failed: %v", err)
			} else if !got.Equal(tC.diff) {
				t.Errorf("got %#v, want %#v", got, tC.diff)
			}
		})

		desc = fmt.Sprintf("NewDiffWithMode(%s, %s, %d)", tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.mode)
		t.Run(desc, func(t *testing.T) {
			got, err := datediff.NewDiffWithMode(tC.start, tC.end, tC.mode)
			if err != nil {
				t.Errorf("failed: %v", err)
			} else if !got.Equal(tC.diff) {
				t.Errorf("got %#v, want %#v", got, tC.diff)
			}
		})
	}
}

func TestNewDiffFails(t *testing.T) {
	type testcase struct {
		start    time.Time
		end      time.Time
		format   string
		expected string
	}
	testCases := []testcase{{
		start:    time.Now().Add(time.Hour),
		end:      time.Now(),
		format:   "%D",
		expected: "start date is after end date",
	}}

	for _, v := range testInvalidFormat {
		tC := testcase{
			start:    time.Now(),
			end:      time.Now().Add(time.Hour),
			format:   v.format,
			expected: v.expected,
		}
		testCases = append(testCases, tC)
	}

	for _, tC := range testCases {
		got, err := datediff.NewDiff(tC.start, tC.end, tC.format)
		if err == nil {
			t.Errorf("NewDiff(%s, %s, %s) = %v, want to fail due to %s",
				tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.format, got, tC.expected)
		} else if err.Error() != tC.expected {
			t.Errorf("NewDiff(%s, %s, %s) failed: %v, want to fail due to %s",
				tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.format, err, tC.expected)
		}
	}
}

func TestString(t *testing.T) {
	testCases, err := loadDatediffRecordsForTest()
	if err != nil {
		t.Fatal(err)
	}

	// TODO: add t.Run for NewDiff and NewDiffWithMode
	for _, tC := range testCases {
		diff, err := datediff.NewDiff(tC.start, tC.end, tC.format)
		if err != nil {
			t.Errorf("NewDiff(%s, %s, %s) failed: %v",
				tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.format, err)
		}
		got := diff.String()
		if got != tC.print {
			t.Errorf("String() = %s, want %s", got, tC.print)
		}
	}
}

func TestStringWithZeros(t *testing.T) {
	testCases, err := loadDatediffRecordsForTest()
	if err != nil {
		t.Fatal(err)
	}

	for _, tC := range testCases {
		diff, err := datediff.NewDiff(tC.start, tC.end, tC.format)
		if err != nil {
			t.Errorf("NewDiff(%s, %s, %s) failed: %v",
				tC.start.Format(dateFmt), tC.end.Format(dateFmt), tC.format, err)
		}
		got := diff.StringWithZeros()
		if got != tC.printWithZeros {
			t.Errorf("StringWithZeros() = %s, want %s", got, tC.printWithZeros)
		}
	}
}

func TestFormats(t *testing.T) {
	start := time.Date(2000, time.April, 17, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(3, -1, -1)
	format := "%Y, %M, %W and %D"
	diff, err := datediff.NewDiff(start, end, format)
	if err != nil {
		t.Errorf("NewDiff(%s, %s, %s) failed: %v",
			start.Format(dateFmt), end.Format(dateFmt), format, err)
	}

	{
		format = "%Y"
		expected := "2 years"
		got, err := diff.Format(format)
		if err != nil {
			t.Errorf("Format(%s) failed: %v", format, err)
		} else if got != expected {
			t.Errorf("Format(%s) = %s, want %s", format, got, expected)
		}
	}

	{
		format = "%Y"
		expected := "2 years"
		got, err := diff.FormatWithZeros(format)
		if err != nil {
			t.Errorf("FormatWithZeros(%s) failed: %v", format, err)
		} else if got != expected {
			t.Errorf("FormatWithZeros(%s) = %s, want %s", format, got, expected)
		}
	}
}

func TestFormatFails(t *testing.T) {
	start := time.Date(2000, time.April, 17, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(3, -1, -1)
	format := "%Y, %M, %W and %D"
	diff, err := datediff.NewDiff(start, end, format)
	if err != nil {
		t.Errorf("NewDiff(%s, %s, %s) failed: %v",
			start.Format(dateFmt), end.Format(dateFmt), format, err)
	}

	for _, tC := range testInvalidFormat {
		got, err := diff.Format(tC.format)
		if err == nil {
			t.Errorf("Format(%s) = %v, want to fail due to %s", tC.format, got, tC.expected)
		} else if err.Error() != tC.expected {
			t.Errorf("Format(%s) failed: %v, want to fail due to %s", tC.format, err, tC.expected)
		}
	}
}
