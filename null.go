package date

import (
	"bytes"
	"database/sql/driver"
	"fmt"
)

type NullDate struct {
	date      Date
	isNotNull bool
}

var (
	ErrNullDateIsNull = fmt.Errorf("NullDate is null")
)

// Factory functions
// --------------------------------------------------

// NullDateFromDate creates a NullDate instance from a Date instance.
func NullDateFromDate(date Date) NullDate {
	return NullDate{
		date:      date,
		isNotNull: true,
	}
}

// NullDateFromDatePtr creates a NullDate instance from a pointer to a Date instance.
func NullDateFromDatePtr(date *Date) NullDate {
	if date == nil {
		return NullDate{}
	}

	return NullDateFromDate(*date)
}

// NullDateForNull returns a NullDate instance representing a null value.
func NullDateForNull() NullDate {
	return NullDate{}
}

// Determination methods
// --------------------------------------------------

// IsNull checks if the NullDate instance is null.
func (nd NullDate) IsNull() bool {
	return !nd.isNotNull
}

// IsNotNull checks if the NullDate instance is not null.
func (nd NullDate) IsNotNull() bool {
	return nd.isNotNull
}

// Comparison methods
// --------------------------------------------------

// Equal checks if the NullDate instance is equal to another NullDate instance.
func (nd NullDate) Equal(target NullDate) bool {
	if nd.IsNull() {
		return nd.isNotNull == target.isNotNull
	}

	return nd.date.Equal(target.date)
}

// NotEqual checks if the NullDate instance is not equal to another NullDate instance.
func (nd NullDate) NotEqual(target NullDate) bool {
	return !nd.Equal(target)
}

// Conversion methods
// --------------------------------------------------

// Ptr returns a pointer to the Date instance if the NullDate instance is not null, otherwise it returns nil.
func (nd NullDate) Ptr() *Date {
	if nd.IsNull() {
		return nil
	}

	return &nd.date
}

// Take returns the Date instance if the NullDate instance is not null, otherwise it returns an error.
func (nd NullDate) Take() (Date, error) {
	if nd.IsNull() {
		return nd.date, fmt.Errorf("Take: %w", ErrNullDateIsNull)
	}

	return nd.date, nil
}

// TakeOr returns the Date instance if the NullDate instance is not null, otherwise it returns the specified fallback Date.
func (nd NullDate) TakeOr(fallback Date) Date {
	if nd.IsNull() {
		return fallback
	}

	return nd.date
}

// MustTake returns the Date instance if the NullDate instance is not null, otherwise it panics.
func (nd NullDate) MustTake() Date {
	d, err := nd.Take()
	if err != nil {
		panic(err)
	}

	return d
}

// String returns the string representation of the NullDate instance.
func (nd NullDate) String() string {
	if nd.IsNull() {
		return "null"
	}

	return nd.date.String()
}

// StringPtr returns a pointer to the string representation of the NullDate instance if it is not null, otherwise it returns nil.
func (nd NullDate) StringPtr() *string {
	if nd.IsNull() {
		return nil
	}

	return nd.date.StringPtr()
}

// Conditional methods
// --------------------------------------------------

// IfSome executes the specified function if the NullDate instance is not null.
func (nd NullDate) IfSome(f func(Date)) {
	if nd.IsNotNull() {
		f(nd.date)
	}
}

// IfSomeWithError executes the specified function if the NullDate instance is not null and returns any error from the function.
func (nd NullDate) IfSomeWithError(f func(Date) error) error {
	if nd.IsNotNull() {
		return f(nd.date)
	}

	return nil
}

// IfNone executes the specified function if the NullDate instance is null.
func (nd NullDate) IfNone(f func()) {
	if nd.IsNull() {
		f()
	}
}

// IfNoneWithError executes the specified function if the NullDate instance is null and returns any error from the function.
func (nd NullDate) IfNoneWithError(f func() error) error {
	if nd.IsNull() {
		return f()
	}

	return nil
}

// Map applies the specified function to the Date instance if the NullDate instance is not null and returns a new NullDate instance.
func (nd NullDate) Map(f func(Date) Date) NullDate {
	if nd.IsNotNull() {
		return NullDateFromDate(f(nd.date))
	}

	return nd
}

// Marshalling methods
// --------------------------------------------------

// Value returns the driver.Value representation of the NullDate instance.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.isNotNull {
		return nil, nil
	}

	return nd.date.Value()
}

// Scan scans a value into the NullDate instance.
func (nd *NullDate) Scan(value interface{}) error {
	if value == nil {
		nd.date, nd.isNotNull = ZeroDate(), false

		return nil
	}

	if err := nd.date.Scan(value); err != nil {
		nd.isNotNull = false

		return fmt.Errorf("NullDate.Scan: %w", err)
	}

	nd.isNotNull = true

	return nil
}

// MarshalText marshals the NullDate instance to a text representation.
func (nd NullDate) MarshalText() ([]byte, error) {
	if nd.isNotNull {
		return nd.date.MarshalText()
	}

	return []byte("null"), nil
}

// UnmarshalText unmarshals a text representation into the NullDate instance.
func (nd *NullDate) UnmarshalText(text []byte) error {
	if string(text) == "null" {
		*nd = NullDate{}

		return nil
	}

	err := nd.date.UnmarshalText(text)
	nd.isNotNull = err == nil

	return fmt.Errorf("NullDate.UnmarshalText: %w", err)
}

// MarshalJSON marshals the NullDate instance to a JSON representation.
func (nd NullDate) MarshalJSON() ([]byte, error) {
	if nd.isNotNull {
		return nd.date.MarshalJSON()
	}

	return []byte("null"), nil
}

// UnmarshalJSON unmarshals a JSON representation into the NullDate instance.
func (nd *NullDate) UnmarshalJSON(json []byte) error {
	if bytes.Equal(json, []byte("null")) {
		*nd = NullDate{}

		return nil
	}

	err := nd.date.UnmarshalJSON(json)
	nd.isNotNull = err == nil

	return fmt.Errorf("NullDate.UnmarshalJSON: %w", err)
}
