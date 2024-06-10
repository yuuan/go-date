package date

import (
	"fmt"
	"time"
)

var (
	ErrDifferentTimeZone        = fmt.Errorf("The start date TimeZone and the end date TimeZone did not match")
	ErrOnlyOneSideIsZero        = fmt.Errorf("Only one side cannot be zero")
	ErrEndDateIsBeforeStartDate = fmt.Errorf("The end date is before the start date")
	ErrRangesDontOverlap        = fmt.Errorf("This range and the target range don't overlap")
)

type DateRange struct {
	start Date
	end   Date
}

// Factory functions
// --------------------------------------------------

func NewDateRange(start, end Date) (DateRange, error) {
	if start.Location() != end.Location() {
		return ZeroDateRange(), ErrDifferentTimeZone
	}

	if start.IsZero() != end.IsZero() {
		return ZeroDateRange(), ErrOnlyOneSideIsZero
	}

	if end.Before(start) {
		return ZeroDateRange(), ErrEndDateIsBeforeStartDate
	}

	return DateRange{
		start: start,
		end:   end,
	}, nil
}

func MustNewDateRange(start, end Date) DateRange {
	r, err := NewDateRange(start, end)
	if err != nil {
		panic(err)
	}

	return r
}

func ZeroDateRange() DateRange {
	return DateRange{}
}

func ParseDateRange(start, end string) (DateRange, error) {
	s, err := Parse(start)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("Unable to parse the first date: %w", err)
	}

	e, err := Parse(end)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("Unable to parse the end date: %w", err)
	}

	return NewDateRange(s, e)
}

func MustParseDateRange(start, end string) DateRange {
	r, err := ParseDateRange(start, end)
	if err != nil {
		panic(err)
	}

	return r
}

func CustomParseDateRange(layout, start, end string) (DateRange, error) {
	s, err := CustomParse(layout, start)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("Unable to parse the first date: %w", err)
	}

	e, err := CustomParse(layout, end)
	if err != nil {
		return ZeroDateRange(), fmt.Errorf("Unable to parse the end date: %w", err)
	}

	return NewDateRange(s, e)
}

func MustCustomParseDateRange(layout, start, end string) DateRange {
	r, err := CustomParseDateRange(layout, start, end)
	if err != nil {
		panic(err)
	}

	return r
}

// Determination methods
// --------------------------------------------------

func (r DateRange) IsZero() bool {
	return r.start.IsZero() && r.end.IsZero()
}

func (r DateRange) OnlyOneDay() bool {
	return r.start.Equal(r.end)
}

// Comparison methods
// --------------------------------------------------

func (r DateRange) Equal(target DateRange) bool {
	return r.start.Equal(target.start) &&
		r.end.Equal(target.end)
}

func (r DateRange) NotEqual(target DateRange) bool {
	return !r.Equal(target)
}

func (r DateRange) StartsOn(date Date) bool {
	return r.start.Equal(date)
}

func (r DateRange) EndsOn(date Date) bool {
	return r.end.Equal(date)
}

func (r DateRange) StartsOnSameDate(target DateRange) bool {
	return r.start.Equal(target.start)
}

func (r DateRange) StartsBefore(target DateRange) bool {
	return r.start.Before(target.start)
}

func (r DateRange) StartsBeforeOrEqual(target DateRange) bool {
	return r.start.BeforeOrEqual(target.start)
}

func (r DateRange) StartsAfter(target DateRange) bool {
	return r.start.After(target.start)
}

func (r DateRange) StartsAfterOrEqual(target DateRange) bool {
	return r.start.AfterOrEqual(target.start)
}

func (r DateRange) EndsOnSameDate(target DateRange) bool {
	return r.end.Equal(target.end)
}

func (r DateRange) EndsBefore(target DateRange) bool {
	return r.end.Before(target.end)
}

func (r DateRange) EndsBeforeOrEqual(target DateRange) bool {
	return r.end.BeforeOrEqual(target.end)
}

func (r DateRange) EndsAfter(target DateRange) bool {
	return r.end.After(target.end)
}

func (r DateRange) EndsAfterOrEqual(target DateRange) bool {
	return r.end.AfterOrEqual(target.end)
}

func (r DateRange) Contains(date Date) bool {
	return r.start.BeforeOrEqual(date) &&
		r.end.AfterOrEqual(date)
}

func (r DateRange) OverlapsWith(target DateRange) bool {
	return r.end.AfterOrEqual(target.start) &&
		target.end.AfterOrEqual(r.start)
}

func (r DateRange) LessThan(target DateRange) bool {
	if r.StartsOnSameDate(target) {
		return r.EndsBefore(target)
	}

	return r.StartsBefore(target)
}

func (r DateRange) LessThanOrEqual(target DateRange) bool {
	return r.Equal(target) || r.LessThan(target)
}

func (r DateRange) GreaterThan(target DateRange) bool {
	return !r.LessThanOrEqual(target)
}

func (r DateRange) GreaterThanOrEqual(target DateRange) bool {
	return !r.LessThan(target)
}

// Conversion methods
// --------------------------------------------------

func (r DateRange) Start() Date {
	return r.start
}

func (r DateRange) End() Date {
	return r.end
}

func (r DateRange) Location() *time.Location {
	return r.start.Location()
}

func (r DateRange) Days() int {
	return int(r.end.value.Sub(r.start.value).Hours() / 24) + 1
}

func (r DateRange) GetOverlapping(target DateRange) (DateRange, error) {
	if !r.OverlapsWith(target) {
		return ZeroDateRange(), ErrRangesDontOverlap
	}

	return DateRange{
		Dates{r.start, target.start}.Max(),
		Dates{r.end, target.end}.Min(),
	}, nil
}

func (r DateRange) Dates() Dates {
	ds := make(Dates, r.Days())

	for d := r.start; d.BeforeOrEqual(r.end); d = d.AddDay() {
		ds = append(ds, d)
	}

	return ds
}

func (r DateRange) String() string {
	return r.start.String() + ":" + r.end.String()
}
