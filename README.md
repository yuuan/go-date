# Date

## Description

`Date` is an immutable struct designed for handling dates without time components.
Using `time.Time` to represent dates requires considering the time part during comparisons, which can lead to unintended issues.
Date eliminates this concern by storing only the date information and ensuring immutability, preventing accidental modifications.
It also provides useful methods for common date operations such as determination, comparison, addition, subtraction, conversion, and marshalling.

## Usage

```go
// Create a new Date instance with today's date.
today := date.Today()

// Create a new Date instance with a specific date.
specificDate := date.New(2024, 3, 15)

// Parse from string.
d, _ := date.Parse("2024-03-15")

// Date calculations.
tomorrow := today.AddDay()
yesterday := today.SubDay()

// Determine if a date is in the past.
if yesterday.IsPast() {
    // do something
}

// Compare dates.
if today.After(yesterday) {
    // do something
}

// Format to string.
str := d.String() // "2024-03-15"
```

# Month

## Description

Month is an immutable struct for handling year and month values without considering specific dates or times.
It allows precise month-based calculations, such as adding or subtracting months, without dealing with days or time components.

## Usage

```go
// Create a new Month instance with the current month.
currentMonth := date.CurrentMonth()

// Create a new Month instance with a specific year and month.
specificMonth := date.NewMonth(2024, 3)

// Parse from string.
m, _ := date.ParseMonth("2024-03")

// Month calculations.
nextMonth := currentMonth.AddMonth()
previousMonth := currentMonth.SubMonth()

// Determine if a month is in the past.
if previousMonth.IsPast() {
    // do something
}

// Compare months.
if currentMonth.After(previousMonth) {
    // do something
}

// Get the DateRange for the month.
dateRange := currentMonth.ToDateRange() // DateRange{start: 2024-03-01, end: 2024-03-31}

// Format to string.
str := m.String() // "2024-03"
```

# DateRange

## Description

DateRange is an immutable struct for handling a date range from a start date to an end date.
It ensures that once created, the range remains unchanged, preventing unintended modifications.
It also allows you to easily determine whether two date ranges overlap.

## Usage

```go
// Create a new DateRange.
start := date.New(2024, 3, 1)
end := date.New(2024, 3, 31)
march := date.NewRange(start, end)

// Check if a date is within range.
isInRange := march.Contains(date.New(2024, 3, 15)) // true

// Get the number of days in the range.
days := march.Days() // 31

// Get all dates in the range.
dates := march.Dates() // []date.Date{date.New(2024, 3, 1), date.New(2024, 3, 2), ...}

// Iterate over dates.
march.Each(func(d date.Date) {
    // Process each date.
})

// Check for overlap.
other := date.NewRange(date.New(2024, 3, 15), date.New(2024, 4, 15))
overlaps := march.Overlaps(other) // true
```

# Installation

```shell
$ go get github.com/yuuan/go-date
```

# Tests

The package provides a way to mock `time.Now()` for testing purposes:

```go
// Set mock time
date.SetNow(func() time.Time {
    return time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
})

// Reset to actual time
defer date.ResetNow()

// Now your tests will use the mocked time
today := date.Today() // 2024-03-15
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
