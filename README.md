# datediff

[![codecov](https://codecov.io/gh/antklim/datediff/branch/main/graph/badge.svg?token=WEV1D7P7WX)](https://codecov.io/gh/antklim/datediff)

The `datediff` is the package to calculate dates difference. The standard `time` package provides functionality to calculate time difference. The returned result is `time.Duration`. Working with durations can be tricky, especially when we need to represent duration in full days, weeks, etc. This package provides the functionality to represent difference between two dates as the amount of full years, months, weeks, or days passed since the start date.

# Installation
`go get github.com/antklim/datediff`

# Usage
```go
package main

import(
  "fmt"

  "github.com/antklim/datediff"
)

func main() {
  dateOfBirth, _ := time.Parse("2006-01-02", "2000-10-01")
  today, _ := time.Parse("2006-01-02", "2010-11-30")
  age, _ := datediff.NewDiff(dateOfBirth, today, "%Y %M %D")
  fmt.Printf("Today I'm %s old", age)

  // Output:
  // Today I'm 10 years 1 month 29 days old
}
```
