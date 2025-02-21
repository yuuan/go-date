package date

import (
	"fmt"
	"strings"
	"time"
)

var (
	ErrEndMonthIsBeforeStartMonth = fmt.Errorf("The end month is before the start month")
)

type Month struct {
	y int // year -1
	m int // month -1
}

// Factory functions
// --------------------------------------------------

// NewMonth creates a new Month instance with the specified year and month.
func NewMonth(year int, month time.Month) Month {
	return Month{year - 1, int(month) - 1}
}

// ZeroMonth returns a zero value Month instance.
func ZeroMonth() Month {
	return Month{}
}

// MonthFromDate creates a new Month instance from a Date instance.
func MonthFromDate(date Date) Month {
	return NewMonth(date.Year(), date.Month())
}

// MonthFromTime creates a new Month instance from a time.Time value.
func MonthFromTime(time time.Time) Month {
	return NewMonth(time.Year(), time.Month())
}

// ParseMonth parses a year-month string in the format "2006-01" and returns a Month instance.
func ParseMonth(value string) (Month, error) {
	d, err := CustomParse("2006-01", value)
	if err != nil {
		return ZeroMonth(), fmt.Errorf("Unable to parse the year-month: %v", err)
	}

	return MonthFromDate(d), nil
}

// MustParseMonth parses a year-month string in the format "2006-01" and returns a Month instance.
// It panics if the parsing fails.
func MustParseMonth(value string) Month {
	m, err := ParseMonth(value)
	if err != nil {
		panic(err)
	}

	return m
}

// CurrentMonth returns the current month.
func CurrentMonth() Month {
	return MonthFromDate(Today())
}

// NextMonth returns the next month.
func NextMonth() Month {
	return CurrentMonth().AddMonth()
}

// LastMonth returns the previous month.
func LastMonth() Month {
	return CurrentMonth().SubMonth()
}

// Determination methods
// --------------------------------------------------

// IsZero checks if the Month instance is a zero value.
func (m Month) IsZero() bool {
	zm := ZeroMonth()

	return m.y == zm.y && m.m == zm.m
}

// IsJanuary checks if the Month instance is January.
func (m Month) IsJanuary() bool {
	return m.Month() == time.January
}

// IsFebruary checks if the Month instance is February.
func (m Month) IsFebruary() bool {
	return m.Month() == time.February
}

// IsMarch checks if the Month instance is March.
func (m Month) IsMarch() bool {
	return m.Month() == time.March
}

// IsApril checks if the Month instance is April.
func (m Month) IsApril() bool {
	return m.Month() == time.April
}

// IsMay checks if the Month instance is May.
func (m Month) IsMay() bool {
	return m.Month() == time.May
}

// IsJune checks if the Month instance is June.
func (m Month) IsJune() bool {
	return m.Month() == time.June
}

// IsJuly checks if the Month instance is July.
func (m Month) IsJuly() bool {
	return m.Month() == time.July
}

// IsAugust checks if the Month instance is August.
func (m Month) IsAugust() bool {
	return m.Month() == time.August
}

// IsSeptember checks if the Month instance is September.
func (m Month) IsSeptember() bool {
	return m.Month() == time.September
}

// IsOctober checks if the Month instance is October.
func (m Month) IsOctober() bool {
	return m.Month() == time.October
}

// IsNovember checks if the Month instance is November.
func (m Month) IsNovember() bool {
	return m.Month() == time.November
}

// IsDecember checks if the Month instance is December.
func (m Month) IsDecember() bool {
	return m.Month() == time.December
}

// IsPast checks if the Month instance is in the past.
func (m Month) IsPast() bool {
	return m.Before(CurrentMonth())
}

// IsFuture checks if the Month instance is in the future.
func (m Month) IsFuture() bool {
	return m.After(CurrentMonth())
}

// IsCurrentMonth checks if the Month instance is the current month.
func (m Month) IsCurrentMonth() bool {
	tm := CurrentMonth()

	return m.y == tm.y && m.m == tm.m
}

// IsNextMonth checks if the Month instance is the next month.
func (m Month) IsNextMonth() bool {
	nm := NextMonth()

	return m.y == nm.y && m.m == nm.m
}

// IsLastMonth checks if the Month instance is the previous month.
func (m Month) IsLastMonth() bool {
	lm := LastMonth()

	return m.y == lm.y && m.m == lm.m
}

// Comparison methods
// --------------------------------------------------

// Compare compares the Month instance with another Month instance.
// It returns 1 if the Month instance is after the other Month, 0 if they are equal, and -1 if it is before.
func (m Month) Compare(month Month) int {
	if m.After(month) {
		return 1
	} else if m.Equal(month) {
		return 0
	}

	return -1
}

// Equal checks if the Month instance is equal to another Month instance.
func (m Month) Equal(month Month) bool {
	return m.y == month.y && m.m == month.m
}

// NotEqual checks if the Month instance is not equal to another Month instance.
func (m Month) NotEqual(month Month) bool {
	return !m.Equal(month)
}

// After checks if the Month instance is after another Month instance.
func (m Month) After(month Month) bool {
	return m.y > month.y ||
		m.y == month.y && m.m > month.m
}

// AfterOrEqual checks if the Month instance is after or equal to another Month instance.
func (m Month) AfterOrEqual(month Month) bool {
	return m.Equal(month) || m.After(month)
}

// Before checks if the Month instance is before another Month instance.
func (m Month) Before(month Month) bool {
	return m.y < month.y ||
		m.y == month.y && m.m < month.m
}

// BeforeOrEqual checks if the Month instance is before or equal to another Month instance.
func (m Month) BeforeOrEqual(month Month) bool {
	return m.Equal(month) || m.Before(month)
}

// Between checks if the Month instance is between two other Month instances.
func (m Month) Between(start, end Month) (bool, error) {
	if start.After(end) {
		return false, ErrEndMonthIsBeforeStartMonth
	}

	return start.BeforeOrEqual(m) && end.AfterOrEqual(m), nil
}

// Addition and Subtraction methods
// --------------------------------------------------

// AddMonth adds one month to the Month instance.
func (m Month) AddMonth() Month {
	return m.AddMonths(1)
}

// AddMonths adds the specified number of months to the Month instance.
func (m Month) AddMonths(months int) Month {
	ms := m.m + months
	ty := m.y + ms/12
	tm := ms % 12
	if tm < 0 {
		ty--
		tm += 12
	}

	return Month{ty, tm}
}

// SubMonth subtracts one month from the Month instance.
func (m Month) SubMonth() Month {
	return m.SubMonths(1)
}

// SubMonths subtracts the specified number of months from the Month instance.
func (m Month) SubMonths(months int) Month {
	return m.AddMonths(months * -1)
}

// AddYear adds one year to the Month instance.
func (m Month) AddYear() Month {
	return m.AddYears(1)
}

// AddYears adds the specified number of years to the Month instance.
func (m Month) AddYears(years int) Month {
	return Month{
		m.y + years,
		m.m,
	}
}

// SubYear subtracts one year from the Month instance.
func (m Month) SubYear() Month {
	return m.SubYears(1)
}

// SubYears subtracts the specified number of years from the Month instance.
func (m Month) SubYears(years int) Month {
	return m.AddYears(years * -1)
}

// Conversion methods
// --------------------------------------------------

// Year returns the year of the Month instance.
func (m Month) Year() int {
	return m.y + 1
}

// Month returns the month of the Month instance.
func (m Month) Month() time.Month {
	return time.Month(m.m + 1)
}

// FirstDate returns the first date of the Month instance.
func (m Month) FirstDate() Date {
	return NewDate(m.Year(), m.Month(), 1)
}

// LastDate returns the last date of the Month instance.
func (m Month) LastDate() Date {
	return m.AddMonth().FirstDate().SubDay()
}

// ToDateRange converts the Month instance to a DateRange instance.
func (m Month) ToDateRange() DateRange {
	r, _ := NewDateRange(m.FirstDate(), m.LastDate())

	return r
}

// Days returns the number of days in the Month instance.
func (m Month) Days() int {
	return m.ToDateRange().Days()
}

// Dates returns the Dates within the Month instance.
func (m Month) Dates() Dates {
	return m.ToDateRange().Dates()
}

// Format formats the Month instance using the specified layout.
func (m Month) Format(layout string) string {
	return m.FirstDate().Format(layout)
}

// Split splits the Month instance into year and month components.
func (m Month) Split() (int, time.Month) {
	return m.Year(), m.Month()
}

// String returns the string representation of the Month instance in the format "2006-01".
func (m Month) String() string {
	f := "%04d-%02d"
	if m.Year() < 0 {
		f = "%05d-%02d"
	}

	return fmt.Sprintf(f, m.Year(), m.Month())
}

// Marshalling methods
// --------------------------------------------------

// MarshalText marshals the Month instance to a text representation.
func (m *Month) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText unmarshals a text representation into the Month instance.
func (m *Month) UnmarshalText(text []byte) error {
	month, err := ParseMonth(string(text))
	if err != nil {
		return err
	}

	m.y, m.m = month.y, month.m

	return nil
}

// MarshalJSON marshals the Month instance to a JSON representation.
func (m *Month) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

// UnmarshalJSON unmarshals a JSON representation into the Month instance.
func (m *Month) UnmarshalJSON(json []byte) error {
	value := strings.Trim(string(json), `"`)

	month, err := ParseMonth(value)
	if err != nil {
		return err
	}

	m.y, m.m = month.y, month.m

	return nil
}
