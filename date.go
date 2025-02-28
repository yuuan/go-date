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

// NewDate creates a new Date instance with the specified year, month, and day.
func NewDate(year int, month time.Month, day int) Date {
	return FromTime(time.Date(year, month, day, 0, 0, 0, 0, location()))
}

// ZeroDate returns a zero value Date instance.
func ZeroDate() Date {
	return Date{}
}

// FromTime creates a new Date instance from a time.Time value.
func FromTime(source time.Time) Date {
	return Date{
		value: startOfDay(source),
	}
}

// Parse parses a date string in the format "2006-01-02" and returns a Date instance.
func Parse(value string) (Date, error) {
	d, err := CustomParse("2006-01-02", value)
	if err != nil {
		return ZeroDate(), fmt.Errorf("Parse: %w", err)
	}

	return d, nil
}

// MustParse parses a date string in the format "2006-01-02" and returns a Date instance.
// It panics if the parsing fails.
func MustParse(value string) Date {
	return MustCustomParse("2006-01-02", value)
}

// CustomParse parses a date string using the specified layout and returns a Date instance.
func CustomParse(layout, value string) (Date, error) {
	time, err := time.ParseInLocation(layout, value, location())
	if err != nil {
		return ZeroDate(), fmt.Errorf("failed to parse date %q with layout %q: %w", value, layout, err)
	}

	return FromTime(time), nil
}

// MustCustomParse parses a date string using the specified layout and returns a Date instance.
// It panics if the parsing fails.
func MustCustomParse(layout, value string) Date {
	date, err := CustomParse(layout, value)
	if err != nil {
		panic(err)
	}

	return date
}

// Today returns the current date.
func Today() Date {
	return FromTime(today())
}

// Yesterday returns the date of the previous day.
func Yesterday() Date {
	return Today().SubDay()
}

// Tomorrow returns the date of the next day.
func Tomorrow() Date {
	return Today().AddDay()
}

// Determination methods
// --------------------------------------------------

// IsZero checks if the Date instance is a zero value.
func (d Date) IsZero() bool {
	return d.value.IsZero()
}

// IsFirstOfMonth checks if the Date instance is the first day of the month.
func (d Date) IsFirstOfMonth() bool {
	return d.value.Day() == 1
}

// IsLastOfMonth checks if the Date instance is the last day of the month.
func (d Date) IsLastOfMonth() bool {
	return d.value.AddDate(0, 0, 1).Day() == 1
}

// IsMonday checks if the Date instance is a Monday.
func (d Date) IsMonday() bool {
	return d.Weekday() == time.Monday
}

// IsTuesday checks if the Date instance is a Tuesday.
func (d Date) IsTuesday() bool {
	return d.Weekday() == time.Tuesday
}

// IsWednesday checks if the Date instance is a Wednesday.
func (d Date) IsWednesday() bool {
	return d.Weekday() == time.Wednesday
}

// IsThursday checks if the Date instance is a Thursday.
func (d Date) IsThursday() bool {
	return d.Weekday() == time.Thursday
}

// IsFriday checks if the Date instance is a Friday.
func (d Date) IsFriday() bool {
	return d.Weekday() == time.Friday
}

// IsSaturday checks if the Date instance is a Saturday.
func (d Date) IsSaturday() bool {
	return d.Weekday() == time.Saturday
}

// IsSunday checks if the Date instance is a Sunday.
func (d Date) IsSunday() bool {
	return d.Weekday() == time.Sunday
}

// IsWeekday checks if the Date instance is a weekday (Monday to Friday).
func (d Date) IsWeekday() bool {
	return !d.IsWeekend()
}

// IsWeekend checks if the Date instance is a weekend (Saturday or Sunday).
func (d Date) IsWeekend() bool {
	return d.IsSaturday() || d.IsSunday()
}

// IsPast checks if the Date instance is in the past.
func (d Date) IsPast() bool {
	return d.Before(Today())
}

// IsPastOrToday checks if the Date instance is in the past or today.
func (d Date) IsPastOrToday() bool {
	return d.BeforeOrEqual(Today())
}

// IsFuture checks if the Date instance is in the future.
func (d Date) IsFuture() bool {
	return d.After(Today())
}

// IsFutureOrToday checks if the Date instance is in the future or today.
func (d Date) IsFutureOrToday() bool {
	return d.AfterOrEqual(Today())
}

// IsToday checks if the Date instance is today.
func (d Date) IsToday() bool {
	return d.Equal(Today())
}

// IsYesterday checks if the Date instance is yesterday.
func (d Date) IsYesterday() bool {
	return d.Equal(Yesterday())
}

// IsTomorrow checks if the Date instance is tomorrow.
func (d Date) IsTomorrow() bool {
	return d.Equal(Tomorrow())
}

// Comparison methods
// --------------------------------------------------

// Compare compares the Date instance with another Date instance.
// It returns -1 if the Date instance is before the other Date, 0 if they are equal, and 1 if it is after.
func (d Date) Compare(date Date) int {
	return d.value.Compare(date.value)
}

// Equal checks if the Date instance is equal to another Date instance.
func (d Date) Equal(date Date) bool {
	return d.value.Equal(date.value)
}

// NotEqual checks if the Date instance is not equal to another Date instance.
func (d Date) NotEqual(date Date) bool {
	return !d.Equal(date)
}

// After checks if the Date instance is after another Date instance.
func (d Date) After(date Date) bool {
	return d.value.After(date.value)
}

// AfterOrEqual checks if the Date instance is after or equal to another Date instance.
func (d Date) AfterOrEqual(date Date) bool {
	return d.Equal(date) || d.After(date)
}

// Before checks if the Date instance is before another Date instance.
func (d Date) Before(date Date) bool {
	return d.value.Before(date.value)
}

// BeforeOrEqual checks if the Date instance is before or equal to another Date instance.
func (d Date) BeforeOrEqual(date Date) bool {
	return d.Equal(date) || d.Before(date)
}

// Between checks if the Date instance is between two other Date instances.
func (d Date) Between(start, end Date) (bool, error) {
	r, err := NewDateRange(start, end)
	if err != nil {
		return false, fmt.Errorf("Between: cannot determine if %v is between %v and %v: %w", d, start, end, err)
	}

	return r.Contains(d), nil
}

// Addition and Subtraction methods
// --------------------------------------------------

// AddDate adds the specified number of years, months, and days to the Date instance.
func (d Date) AddDate(years, months, days int) Date {
	return FromTime(d.value.AddDate(years, months, days))
}

// AddDay adds one day to the Date instance.
func (d Date) AddDay() Date {
	return d.AddDays(1)
}

// AddDays adds the specified number of days to the Date instance.
func (d Date) AddDays(days int) Date {
	return d.AddDate(0, 0, days)
}

// SubDay subtracts one day from the Date instance.
func (d Date) SubDay() Date {
	return d.SubDays(1)
}

// SubDays subtracts the specified number of days from the Date instance.
func (d Date) SubDays(days int) Date {
	return d.AddDays(days * -1)
}

// AddWeek adds one week to the Date instance.
func (d Date) AddWeek() Date {
	return d.AddWeeks(1)
}

// AddWeeks adds the specified number of weeks to the Date instance.
func (d Date) AddWeeks(weeks int) Date {
	return d.AddDate(0, 0, weeks*7)
}

// SubWeek subtracts one week from the Date instance.
func (d Date) SubWeek() Date {
	return d.SubWeeks(1)
}

// SubWeeks subtracts the specified number of weeks from the Date instance.
func (d Date) SubWeeks(weeks int) Date {
	return d.AddWeeks(weeks * -1)
}

// AddMonth adds one month to the Date instance.
func (d Date) AddMonth() Date {
	return d.AddMonths(1)
}

// AddMonths adds the specified number of months to the Date instance.
func (d Date) AddMonths(months int) Date {
	m := d.ToMonth().AddMonths(months)
	date := NewDate(m.Year(), m.Month(), d.Day())

	if date.ToMonth().Equal(m) {
		return date
	}

	return date.SubDays(date.Day())
}

// SubMonth subtracts one month from the Date instance.
func (d Date) SubMonth() Date {
	return d.SubMonths(1)
}

// SubMonths subtracts the specified number of months from the Date instance.
func (d Date) SubMonths(months int) Date {
	return d.AddMonths(months * -1)
}

// AddYear adds one year to the Date instance.
func (d Date) AddYear() Date {
	return d.AddYears(1)
}

// AddYears adds the specified number of years to the Date instance.
func (d Date) AddYears(years int) Date {
	return d.AddDate(years, 0, 0)
}

// SubYear subtracts one year from the Date instance.
func (d Date) SubYear() Date {
	return d.SubYears(1)
}

// SubYears subtracts the specified number of years from the Date instance.
func (d Date) SubYears(years int) Date {
	return d.AddYears(years * -1)
}

// StartOfMonth returns the first day of the month for the Date instance.
func (d Date) StartOfMonth() Date {
	return NewDate(d.Year(), d.Month(), 1)
}

// EndOfMonth returns the last day of the month for the Date instance.
func (d Date) EndOfMonth() Date {
	return d.StartOfMonth().AddMonth().SubDay()
}

// StartOfYear returns the first day of the year for the Date instance.
func (d Date) StartOfYear() Date {
	return NewDate(d.Year(), 1, 1)
}

// EndOfYear returns the last day of the year for the Date instance.
func (d Date) EndOfYear() Date {
	return d.StartOfYear().AddYear().SubDay()
}

// Conversion methods
// --------------------------------------------------

// ToMonth converts the Date instance to a Month instance.
func (d Date) ToMonth() Month {
	return NewMonth(d.Year(), d.Month())
}

// Nullable converts the Date instance to a NullDate instance.
func (d Date) Nullable() NullDate {
	return NullDateFromDate(d)
}

// Time returns the time.Time value of the Date instance.
func (d Date) Time() time.Time {
	return d.value
}

// At returns a time.Time with the specified time, keeping the same date.
func (d Date) At(hour, min, sec, nsec int) time.Time {
	return time.Date(
		d.value.Year(),
		d.value.Month(),
		d.value.Day(),
		hour,
		min,
		sec,
		nsec,
		d.value.Location(),
	)
}

// Year returns the year of the Date instance.
func (d Date) Year() int {
	return d.value.Year()
}

// Month returns the month of the Date instance.
func (d Date) Month() time.Month {
	return d.value.Month()
}

// Day returns the day of the Date instance.
func (d Date) Day() int {
	return d.value.Day()
}

// YearDay returns the day of the year of the Date instance.
func (d Date) YearDay() int {
	return d.value.YearDay()
}

// ISOWeek returns the ISO 8601 year and week number of the Date instance.
func (d Date) ISOWeek() (int, int) {
	return d.value.ISOWeek()
}

// Weekday returns the day of the week of the Date instance.
func (d Date) Weekday() time.Weekday {
	return d.value.Weekday()
}

// Location returns the time.Location of the Date instance.
func (d Date) Location() *time.Location {
	return d.value.Location()
}

// Format formats the Date instance using the specified layout.
func (d Date) Format(layout string) string {
	return d.value.Format(layout)
}

// Split splits the Date instance into year, month, and day components.
func (d Date) Split() (int, time.Month, int) {
	return d.value.Date()
}

// String returns the string representation of the Date instance in the format "2006-01-02".
func (d Date) String() string {
	return d.Format("2006-01-02")
}

// StringPtr returns a pointer to the string representation of the Date instance.
func (d Date) StringPtr() *string {
	date := d.String()

	return &date
}

// Marshalling methods
// --------------------------------------------------

// Value returns the driver.Value representation of the Date instance.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan scans a value into the Date instance.
func (d *Date) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		return nil

	case time.Time:
		date := FromTime(value)
		if date.IsZero() {
			return fmt.Errorf("Scan: value is zero value of time.Time")
		}

		*d = date

	case string:
		if value == "" {
			return nil
		}

		date, err := Parse(value)
		if err != nil {
			return fmt.Errorf("Scan: %w", err)
		}

		*d = date

	case []byte:
		if len(value) == 0 {
			return nil
		}

		return d.Scan(string(value))

	default:
		return fmt.Errorf("Scan: unable to scan type %T into Date", value)
	}

	return nil
}

// MarshalText marshals the Date instance to a text representation.
func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText unmarshals a text representation into the Date instance.
func (d *Date) UnmarshalText(text []byte) error {
	date, err := Parse(string(text))
	if err != nil {
		return fmt.Errorf("Date.UnmarshalText: %w", err)
	}

	d.value = date.value

	return nil
}

// MarshalJSON marshals the Date instance to a JSON representation.
func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON unmarshals a JSON representation into the Date instance.
func (d *Date) UnmarshalJSON(json []byte) error {
	value := strings.Trim(string(json), `"`)

	date, err := Parse(value)
	if err != nil {
		return fmt.Errorf("Date.UnmarshalJSON: %w", err)
	}

	d.value = date.value

	return nil
}
