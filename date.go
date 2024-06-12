package date

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	value time.Time
}

// Factory functions
// --------------------------------------------------

func NewDate(year int, month time.Month, day int) Date {
	return FromTime(time.Date(year, month, day, 0, 0, 0, 0, location()))
}

func ZeroDate() Date {
	return Date{}
}

func FromTime(source time.Time) Date {
	return Date{
		value: startOfDay(source),
	}
}

func Parse(value string) (Date, error) {
	return CustomParse("2006-01-02", value)
}

func MustParse(value string) Date {
	return MustCustomParse("2006-01-02", value)
}

func CustomParse(layout, value string) (Date, error) {
	time, err := time.ParseInLocation(layout, value, location())
	if err != nil {
		return ZeroDate(), fmt.Errorf("Unable to parse the date: %w", err)
	}

	return FromTime(time), nil
}

func MustCustomParse(layout, value string) Date {
	date, err := CustomParse(layout, value)
	if err != nil {
		panic(err)
	}

	return date
}

func Today() Date {
	return FromTime(today())
}

func Yesterday() Date {
	return Today().SubDay()
}

func Tomorrow() Date {
	return Today().AddDay()
}

// Determination methods
// --------------------------------------------------

func (d Date) IsZero() bool {
	return d.value.IsZero()
}

func (d Date) IsFirstOfMonth() bool {
	return d.value.Day() == 1
}

func (d Date) IsLastOfMonth() bool {
	return d.value.AddDate(0, 0, 1).Day() == 1
}

func (d Date) IsMonday() bool {
	return d.Weekday() == time.Monday
}

func (d Date) IsTuesday() bool {
	return d.Weekday() == time.Tuesday
}

func (d Date) IsWednesday() bool {
	return d.Weekday() == time.Wednesday
}

func (d Date) IsThursday() bool {
	return d.Weekday() == time.Thursday
}

func (d Date) IsFriday() bool {
	return d.Weekday() == time.Friday
}

func (d Date) IsSaturday() bool {
	return d.Weekday() == time.Saturday
}

func (d Date) IsSunday() bool {
	return d.Weekday() == time.Sunday
}

func (d Date) IsWeekday() bool {
	return !d.IsWeekend()
}

func (d Date) IsWeekend() bool {
	return d.IsSaturday() || d.IsSunday()
}

func (d Date) IsPast() bool {
	return d.Before(Today())
}

func (d Date) IsFuture() bool {
	return d.After(Today())
}

func (d Date) IsToday() bool {
	return d.Equal(Today())
}

func (d Date) IsYesterday() bool {
	return d.Equal(Yesterday())
}

func (d Date) IsTomorrow() bool {
	return d.Equal(Tomorrow())
}

// Comparison methods
// --------------------------------------------------

func (d Date) Compare(date Date) int {
	return d.value.Compare(date.value)
}

func (d Date) Equal(date Date) bool {
	return d.value.Equal(date.value)
}

func (d Date) NotEqual(date Date) bool {
	return !d.Equal(date)
}

func (d Date) After(date Date) bool {
	return d.value.After(date.value)
}

func (d Date) AfterOrEqual(date Date) bool {
	return d.Equal(date) || d.After(date)
}

func (d Date) Before(date Date) bool {
	return d.value.Before(date.value)
}

func (d Date) BeforeOrEqual(date Date) bool {
	return d.Equal(date) || d.Before(date)
}

func (d Date) Between(start, end Date) (bool, error) {
	r, err := NewDateRange(start, end)
	if err != nil {
		return false, err
	}

	return r.Contains(d), nil
}

// Addition and Subtraction methods
// --------------------------------------------------

func (d Date) AddDate(years, months, days int) Date {
	return FromTime(d.value.AddDate(years, months, days))
}

func (d Date) AddDay() Date {
	return d.AddDays(1)
}

func (d Date) AddDays(days int) Date {
	return d.AddDate(0, 0, days)
}

func (d Date) SubDay() Date {
	return d.SubDays(1)
}

func (d Date) SubDays(days int) Date {
	return d.AddDays(days * -1)
}

func (d Date) AddWeek() Date {
	return d.AddWeeks(1)
}

func (d Date) AddWeeks(weeks int) Date {
	return d.AddDate(0, 0, weeks*7)
}

func (d Date) SubWeek() Date {
	return d.SubWeeks(1)
}

func (d Date) SubWeeks(weeks int) Date {
	return d.AddWeeks(weeks * -1)
}

func (d Date) AddMonth() Date {
	return d.AddMonths(1)
}

func (d Date) AddMonths(months int) Date {
	return d.AddDate(0, months, 0)
}

func (d Date) SubMonth() Date {
	return d.SubMonths(1)
}

func (d Date) SubMonths(months int) Date {
	return d.AddMonths(months * -1)
}

func (d Date) AddYear() Date {
	return d.AddYears(1)
}

func (d Date) AddYears(years int) Date {
	return d.AddDate(years, 0, 0)
}

func (d Date) SubYear() Date {
	return d.SubYears(1)
}

func (d Date) SubYears(years int) Date {
	return d.AddYears(years * -1)
}

func (d Date) StartOfMonth() Date {
	return NewDate(d.Year(), d.Month(), 1)
}

func (d Date) EndOfMonth() Date {
	return d.StartOfMonth().AddMonth().SubDay()
}

func (d Date) StartOfYear() Date {
	return NewDate(d.Year(), 1, 1)
}

func (d Date) EndOfYear() Date {
	return d.StartOfYear().AddYear().SubDay()
}

// Conversion methods
// --------------------------------------------------

func (d Date) Time() time.Time {
	return d.value
}

func (d Date) Year() int {
	return d.value.Year()
}

func (d Date) Month() time.Month {
	return d.value.Month()
}

func (d Date) Day() int {
	return d.value.Day()
}

func (d Date) YearDay() int {
	return d.value.YearDay()
}

func (d Date) ISOWeek() (int, int) {
	return d.value.ISOWeek()
}

func (d Date) Weekday() time.Weekday {
	return d.value.Weekday()
}

func (d Date) Location() *time.Location {
	return d.value.Location()
}

func (d Date) Format(layout string) string {
	return d.value.Format(layout)
}

func (d Date) Split() (int, time.Month, int) {
	return d.value.Date()
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func (d Date) StringPtr() *string {
	date := d.String()

	return &date
}

// Marshalling methods
// --------------------------------------------------

func (d *Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *Date) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		return nil

	case string:
		if value == "" {
			return nil
		}

		date, err := Parse(value)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*d = date

	case []byte:
		if len(value) == 0 {
			return nil
		}

		return d.Scan(string(value))

	default:
		return fmt.Errorf("Scan: Unable to scan type %T into Date", value)
	}

	return nil
}

func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	date, err := Parse(string(text))
	if err != nil {
		return err
	}

	d.value = date.value

	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) UnmarshalJSON(json []byte) error {
	value := strings.Trim(string(json), `"`)

	date, err := Parse(value)
	if err != nil {
		return err
	}

	d.value = date.value

	return nil
}
