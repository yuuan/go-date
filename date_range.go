package date

import (
	"fmt"
	"time"
)

var (
	ErrDifferentTimeZone        = fmt.Errorf("start date TimeZone and end date TimeZone did not match")
	ErrOnlyOneSideIsZero        = fmt.Errorf("only one side cannot be zero")
	ErrEndDateIsBeforeStartDate = fmt.Errorf("end date is before start date")
	ErrRangesDontOverlap        = fmt.Errorf("this range and target range don't overlap")
)

type DateRange struct {
	start Date
	end   Date
}

// Factory functions
// --------------------------------------------------

// NewDateRange creates a new DateRange instance with the specified start and end dates.
// It returns an error if Date instances with different Location are passed.
// The comparison is done by comparing the memory addresses of the Location instances.
func NewDateRange(start, end Date) (DateRange, error) {
	if start.Location() != end.Location() {
		return ZeroDateRange(), fmt.Errorf("NewDateRange: %w", ErrDifferentTimeZone)
	}

	if start.IsZero() != end.IsZero() {
		return ZeroDateRange(), fmt.Errorf("NewDateRange: %w", ErrOnlyOneSideIsZero)
	}

	if end.Before(start) {
		return ZeroDateRange(), fmt.Errorf("NewDateRange: %w", ErrEndDateIsBeforeStartDate)
	}

	return DateRange{
		start: start,
		end:   end,
	}, nil
}

// MustNewDateRange creates a new DateRange instance with the specified start and end dates.
// It panics if the creation fails.
// It returns an error if Date instances with different Location are passed.
// The comparison is done by comparing the memory addresses of the Location instances.
func MustNewDateRange(start, end Date) DateRange {
	r, err := NewDateRange(start, end)
	if err != nil {
		panic(err)
	}

	return r
}

// ZeroDateRange returns a zero value DateRange instance.
func ZeroDateRange() DateRange {
	return DateRange{}
}

// ParseDateRange parses start and end date strings in the format "2006-01-02" and returns a DateRange instance.
func ParseDateRange(start, end string) (DateRange, error) {
	s, err := Parse(start)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("ParseDateRange: failed to parse start date: %w", err)
	}

	e, err := Parse(end)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("ParseDateRange: failed to parse end date: %w", err)
	}

	dr, err := NewDateRange(s, e)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("ParseDateRange: %w", err)
	}

	return dr, nil
}

// MustParseDateRange parses start and end date strings in the format "2006-01-02" and returns a DateRange instance.
// It panics if the parsing fails.
func MustParseDateRange(start, end string) DateRange {
	r, err := ParseDateRange(start, end)
	if err != nil {
		panic(err)
	}

	return r
}

// CustomParseDateRange parses start and end date strings using the specified layout and returns a DateRange instance.
func CustomParseDateRange(layout, start, end string) (DateRange, error) {
	s, err := CustomParse(layout, start)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("CustomParseDateRange: failed to parse start date: %w", err)
	}

	e, err := CustomParse(layout, end)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("CustomParseDateRange: failed to parse end date: %w", err)
	}

	dr, err := NewDateRange(s, e)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("CustomParseDateRange: %w", err)
	}

	return dr, nil
}

// MustCustomParseDateRange parses start and end date strings using the specified layout and returns a DateRange instance.
// It panics if the parsing fails.
func MustCustomParseDateRange(layout, start, end string) DateRange {
	r, err := CustomParseDateRange(layout, start, end)
	if err != nil {
		panic(err)
	}

	return r
}

// Determination methods
// --------------------------------------------------

// IsZero checks if the DateRange instance is a zero value.
func (r DateRange) IsZero() bool {
	return r.start.IsZero() && r.end.IsZero()
}

// OnlyOneDay checks if the DateRange instance represents only one day.
func (r DateRange) OnlyOneDay() bool {
	return r.start.Equal(r.end)
}

// Comparison methods
// --------------------------------------------------

// Equal checks if the DateRange instance is equal to another DateRange instance.
func (r DateRange) Equal(target DateRange) bool {
	return r.start.Equal(target.start) &&
		r.end.Equal(target.end)
}

// NotEqual checks if the DateRange instance is not equal to another DateRange instance.
func (r DateRange) NotEqual(target DateRange) bool {
	return !r.Equal(target)
}

// StartsOn checks if the DateRange instance starts on the specified date.
func (r DateRange) StartsOn(date Date) bool {
	return r.start.Equal(date)
}

// EndsOn checks if the DateRange instance ends on the specified date.
func (r DateRange) EndsOn(date Date) bool {
	return r.end.Equal(date)
}

// StartsOnSameDate checks if the DateRange instance starts on the same date as another DateRange instance.
func (r DateRange) StartsOnSameDate(target DateRange) bool {
	return r.start.Equal(target.start)
}

// StartsBefore checks if the DateRange instance starts before another DateRange instance.
func (r DateRange) StartsBefore(target DateRange) bool {
	return r.start.Before(target.start)
}

// StartsBeforeOrEqual checks if the DateRange instance starts before or equal to another DateRange instance.
func (r DateRange) StartsBeforeOrEqual(target DateRange) bool {
	return r.start.BeforeOrEqual(target.start)
}

// StartsAfter checks if the DateRange instance starts after another DateRange instance.
func (r DateRange) StartsAfter(target DateRange) bool {
	return r.start.After(target.start)
}

// StartsAfterOrEqual checks if the DateRange instance starts after or equal to another DateRange instance.
func (r DateRange) StartsAfterOrEqual(target DateRange) bool {
	return r.start.AfterOrEqual(target.start)
}

// EndsOnSameDate checks if the DateRange instance ends on the same date as another DateRange instance.
func (r DateRange) EndsOnSameDate(target DateRange) bool {
	return r.end.Equal(target.end)
}

// EndsBefore checks if the DateRange instance ends before another DateRange instance.
func (r DateRange) EndsBefore(target DateRange) bool {
	return r.end.Before(target.end)
}

// EndsBeforeOrEqual checks if the DateRange instance ends before or equal to another DateRange instance.
func (r DateRange) EndsBeforeOrEqual(target DateRange) bool {
	return r.end.BeforeOrEqual(target.end)
}

// EndsAfter checks if the DateRange instance ends after another DateRange instance.
func (r DateRange) EndsAfter(target DateRange) bool {
	return r.end.After(target.end)
}

// EndsAfterOrEqual checks if the DateRange instance ends after or equal to another DateRange instance.
func (r DateRange) EndsAfterOrEqual(target DateRange) bool {
	return r.end.AfterOrEqual(target.end)
}

// Contains checks if the DateRange instance contains the specified date.
func (r DateRange) Contains(date Date) bool {
	return r.start.BeforeOrEqual(date) &&
		r.end.AfterOrEqual(date)
}

// OverlapsWith checks if the DateRange instance overlaps with another DateRange instance.
func (r DateRange) OverlapsWith(target DateRange) bool {
	return r.end.AfterOrEqual(target.start) &&
		target.end.AfterOrEqual(r.start)
}

// LessThan checks if the DateRange instance is less than another DateRange instance.
// It returns true if the start date of the current range is before the start date of the target range,
// or if the start dates are equal and the end date of the current range is before the end date of the target range.
func (r DateRange) LessThan(target DateRange) bool {
	if r.StartsOnSameDate(target) {
		return r.EndsBefore(target)
	}

	return r.StartsBefore(target)
}

// LessThanOrEqual checks if the DateRange instance is less than or equal to another DateRange instance.
// It returns true if the current range is equal to or less than the target range.
func (r DateRange) LessThanOrEqual(target DateRange) bool {
	return r.Equal(target) || r.LessThan(target)
}

// GreaterThan checks if the DateRange instance is greater than another DateRange instance.
// It returns true if the start date of the current range is after the start date of the target range,
// or if the start dates are equal and the end date of the current range is after the end date of the target range.
func (r DateRange) GreaterThan(target DateRange) bool {
	return !r.LessThanOrEqual(target)
}

// GreaterThanOrEqual checks if the DateRange instance is greater than or equal to another DateRange instance.
// It returns true if the current range is equal to or greater than the target range.
func (r DateRange) GreaterThanOrEqual(target DateRange) bool {
	return !r.LessThan(target)
}

// Conversion methods
// --------------------------------------------------

// Start returns the start date of the DateRange instance.
func (r DateRange) Start() Date {
	return r.start
}

// End returns the end date of the DateRange instance.
func (r DateRange) End() Date {
	return r.end
}

// Location returns the time.Location of the DateRange instance.
func (r DateRange) Location() *time.Location {
	return r.start.Location()
}

// Days returns the number of days in the DateRange instance.
func (r DateRange) Days() int {
	return int(r.end.value.Sub(r.start.value).Hours()/24) + 1
}

// GetOverlapping returns the overlapping DateRange between the DateRange instance and another DateRange instance.
func (r DateRange) GetOverlapping(target DateRange) (DateRange, error) {
	if !r.OverlapsWith(target) {
		return ZeroDateRange(), fmt.Errorf("GetOverlapping: %w", ErrRangesDontOverlap)
	}

	return DateRange{
		Dates{r.start, target.start}.MustMax(),
		Dates{r.end, target.end}.MustMin(),
	}, nil
}

// Dates returns the Dates within the DateRange instance.
func (r DateRange) Dates() Dates {
	ds := make(Dates, 0, r.Days())

	for d := r.start; d.BeforeOrEqual(r.end); d = d.AddDay() {
		ds = append(ds, d)
	}

	return ds
}

// String returns the string representation of the DateRange instance in the format "start/end".
func (r DateRange) String() string {
	return r.start.String() + "/" + r.end.String()
}
