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

func NewMonth(year int, month time.Month) Month {
	return Month{year - 1, int(month) - 1}
}

func ZeroMonth() Month {
	return Month{}
}

func MonthFromDate(date Date) Month {
	return NewMonth(date.Year(), date.Month())
}

func MonthFromTime(time time.Time) Month {
	return NewMonth(time.Year(), time.Month())
}

func ParseMonth(value string) (Month, error) {
	d, err := CustomParse("2006-01", value)
	if err != nil {
		return ZeroMonth(), fmt.Errorf("Unable to parse the year-month: %v", err)
	}

	return MonthFromDate(d), nil
}

func MustParseMonth(value string) Month {
	m, err := ParseMonth(value)
	if err != nil {
		panic(err)
	}

	return m
}

func CurrentMonth() Month {
	return MonthFromDate(Today())
}

func NextMonth() Month {
	return CurrentMonth().AddMonth()
}

func LastMonth() Month {
	return CurrentMonth().SubMonth()
}

// Determination methods
// --------------------------------------------------

func (m Month) IsZero() bool {
	zm := ZeroMonth()

	return m.y == zm.y && m.m == zm.m
}

func (m Month) IsJanuary() bool {
	return m.Month() == time.January
}

func (m Month) IsFebruary() bool {
	return m.Month() == time.February
}

func (m Month) IsMarch() bool {
	return m.Month() == time.March
}

func (m Month) IsApril() bool {
	return m.Month() == time.April
}

func (m Month) IsMay() bool {
	return m.Month() == time.May
}

func (m Month) IsJune() bool {
	return m.Month() == time.June
}

func (m Month) IsJuly() bool {
	return m.Month() == time.July
}

func (m Month) IsAugust() bool {
	return m.Month() == time.August
}

func (m Month) IsSeptember() bool {
	return m.Month() == time.September
}

func (m Month) IsOctober() bool {
	return m.Month() == time.October
}

func (m Month) IsNovember() bool {
	return m.Month() == time.November
}

func (m Month) IsDecember() bool {
	return m.Month() == time.December
}

func (m Month) IsPast() bool {
	return m.Before(CurrentMonth())
}

func (m Month) IsFuture() bool {
	return m.After(CurrentMonth())
}

func (m Month) IsCurrentMonth() bool {
	tm := CurrentMonth()

	return m.y == tm.y && m.m == tm.m
}

func (m Month) IsNextMonth() bool {
	nm := NextMonth()

	return m.y == nm.y && m.m == nm.m
}

func (m Month) IsLastMonth() bool {
	lm := LastMonth()

	return m.y == lm.y && m.m == lm.m
}

// Comparison methods
// --------------------------------------------------

func (m Month) Compare(month Month) int {
	if m.After(month) {
		return 1
	} else if m.Equal(month) {
		return 0
	}

	return -1
}

func (m Month) Equal(month Month) bool {
	return m.y == month.y && m.m == month.m
}

func (m Month) NotEqual(month Month) bool {
	return !m.Equal(month)
}

func (m Month) After(month Month) bool {
	return m.y > month.y ||
		m.y == month.y && m.m > month.m
}

func (m Month) AfterOrEqual(month Month) bool {
	return m.Equal(month) || m.After(month)
}

func (m Month) Before(month Month) bool {
	return m.y < month.y ||
		m.y == month.y && m.m < month.m
}

func (m Month) BeforeOrEqual(month Month) bool {
	return m.Equal(month) || m.Before(month)
}

func (m Month) Between(start, end Month) (bool, error) {
	if start.After(end) {
		return false, ErrEndMonthIsBeforeStartMonth
	}

	return start.BeforeOrEqual(m) && end.AfterOrEqual(m), nil
}

// Addition and Subtraction methods
// --------------------------------------------------

func (m Month) AddMonth() Month {
	return m.AddMonths(1)
}

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

func (m Month) SubMonth() Month {
	return m.SubMonths(1)
}

func (m Month) SubMonths(months int) Month {
	return m.AddMonths(months * -1)
}

func (m Month) AddYear() Month {
	return m.AddYears(1)
}

func (m Month) AddYears(years int) Month {
	return Month{
		m.y + years,
		m.m,
	}
}

func (m Month) SubYear() Month {
	return m.SubYears(1)
}

func (m Month) SubYears(years int) Month {
	return m.AddYears(years * -1)
}

// Conversion methods
// --------------------------------------------------

func (m Month) Year() int {
	return m.y + 1
}

func (m Month) Month() time.Month {
	return time.Month(m.m + 1)
}

func (m Month) FirstDate() Date {
	return NewDate(m.Year(), m.Month(), 1)
}

func (m Month) LastDate() Date {
	return m.AddMonth().FirstDate().SubDay()
}

func (m Month) ToDateRange() DateRange {
	r, _ := NewDateRange(m.FirstDate(), m.LastDate())

	return r
}

func (m Month) Days() int {
	return m.ToDateRange().Days()
}

func (m Month) Dates() Dates {
	return m.ToDateRange().Dates()
}

func (m Month) Format(layout string) string {
	return m.FirstDate().Format(layout)
}

func (m Month) Split() (int, time.Month) {
	return m.Year(), m.Month()
}

func (m Month) String() string {
	f := "%04d-%02d"
	if m.Year() < 0 {
		f = "%05d-%02d"
	}

	return fmt.Sprintf(f, m.Year(), m.Month())
}

// Marshalling methods
// --------------------------------------------------

func (m *Month) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *Month) UnmarshalText(text []byte) error {
	month, err := ParseMonth(string(text))
	if err != nil {
		return err
	}

	m.y, m.m = month.y, month.m

	return nil
}

func (m *Month) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

func (m *Month) UnmarshalJSON(json []byte) error {
	value := strings.Trim(string(json), `"`)

	month, err := ParseMonth(value)
	if err != nil {
		return err
	}

	m.y, m.m = month.y, month.m

	return nil
}
