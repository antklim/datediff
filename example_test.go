package datediff_test

import (
	"fmt"
	"time"

	"github.com/antklim/datediff"
)

func ExampleDiff_Equal() {
	d1, _ := time.Parse("2006-01-02", "2000-10-01")
	d2, _ := time.Parse("2006-01-02", "2000-10-30")

	diff1, _ := datediff.NewDiff(d1, d2, "%Y %M %D")
	diff2, _ := datediff.NewDiff(d1, d2, "%Y %M %D")
	fmt.Println(diff1.Equal(diff2))

	diff3, _ := datediff.NewDiff(d1, d2.Add(-48*time.Hour), "%Y %M %D")
	fmt.Println(diff1.Equal(diff3))

	diff4, _ := datediff.NewDiff(d1, d2, "%Y")
	fmt.Println(diff1.Equal(diff4))

	// Output:
	// true
	// false
	// false
}

func ExampleDiff_Format() {
	d1, _ := time.Parse("2006-01-02", "2000-10-01")
	d2, _ := time.Parse("2006-01-02", "2010-11-30")

	diff, _ := datediff.NewDiff(d1, d2, "%Y %M %D")
	s, _ := diff.Format("%y anos")
	fmt.Println(s)

	// Output:
	// 10 anos
}

func ExampleDiff_FormatWithZeros() {
	d1, _ := time.Parse("2006-01-02", "2000-10-01")
	d2, _ := time.Parse("2006-01-02", "2010-10-30")

	diff, _ := datediff.NewDiff(d1, d2, "%Y %M %D")
	s, _ := diff.FormatWithZeros("%y anos %m meses %d dias")
	fmt.Println(s)

	// Output:
	// 10 anos 0 meses 29 dias
}

func ExampleDiff_String() {
	d1, _ := time.Parse("2006-01-02", "2000-10-01")
	d2, _ := time.Parse("2006-01-02", "2010-11-30")

	diff1, _ := datediff.NewDiff(d1, d2, "%Y")
	diff2, _ := datediff.NewDiff(d1, d2, "%Y %M")
	diff3, _ := datediff.NewDiff(d1, d2, "%Y %M %D")

	fmt.Println(diff1)
	fmt.Println(diff2)
	fmt.Println(diff3)

	// Output:
	// 10 years
	// 10 years 1 month
	// 10 years 1 month 29 days
}

func ExampleDiff_StringWithZeros() {
	d1, _ := time.Parse("2006-01-02", "2000-10-01")
	d2, _ := time.Parse("2006-01-02", "2010-10-30")

	diff, _ := datediff.NewDiff(d1, d2, "%Y %M %D")
	fmt.Println(diff.StringWithZeros())

	// Output:
	// 10 years 0 months 29 days
}
