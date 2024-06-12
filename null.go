package date

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
)

type NullDate struct {
	date      Date
	isNotNull bool
}

var (
	ErrNullDateIsNull = errors.New("NullDate is null")
)

// Factory functions
// --------------------------------------------------

func NullDateFromDate(date Date) NullDate {
	return NullDate{
		date:      date,
		isNotNull: true,
	}
}

func NullDateFromDatePtr(date *Date) NullDate {
	if date == nil {
		return NullDate{}
	}

	return NullDateFromDate(*date)
}

func NullDateForNull() NullDate {
	return NullDate{}
}

// Determination methods
// --------------------------------------------------

func (nd *NullDate) IsNull() bool {
	return !nd.isNotNull
}

func (nd *NullDate) IsNotNull() bool {
	return nd.isNotNull
}

// Comparison methods
// --------------------------------------------------

func (nd *NullDate) Equal(target NullDate) bool {
	if nd.IsNull() {
		return nd.isNotNull == target.isNotNull
	}

	return nd.date.Equal(target.date)
}

func (nd *NullDate) NotEqual(target NullDate) bool {
	return !nd.Equal(target)
}

// Conversion methods
// --------------------------------------------------

func (nd *NullDate) Ptr() *Date {
	if nd.IsNull() {
		return nil
	}

	return &nd.date
}

func (nd *NullDate) Take() (Date, error) {
	if nd.IsNull() {
		return nd.date, fmt.Errorf("Take: %w", ErrNullDateIsNull)
	}

	return nd.date, nil
}

func (nd *NullDate) TakeOr(fallback Date) Date {
	if nd.IsNull() {
		return fallback
	}

	return nd.date
}

func (nd NullDate) StringPtr() *string {
	if nd.IsNull() {
		return nil
	}

	return nd.date.StringPtr()
}

// Marshalling methods
// --------------------------------------------------

func (nd NullDate) Value() (driver.Value, error) {
	if !nd.isNotNull {
		return nil, nil
	}

	return nd.date.Value()
}

func (nd *NullDate) Scan(value interface{}) error {
	if value == nil {
		nd.date, nd.isNotNull = ZeroDate(), false

		return nil
	}

	if err := nd.date.Scan(value); err != nil {
		nd.isNotNull = false

		return fmt.Errorf("Scan: %w", err)
	}

	nd.isNotNull = true

	return nil
}

func (nd NullDate) MarshalText() ([]byte, error) {
	if nd.isNotNull {
		return nd.date.MarshalText()
	}

	return []byte("null"), nil
}

func (nd *NullDate) UnmarshalText(text []byte) error {
	if string(text) == "null" {
		*nd = NullDate{}

		return nil
	}

	err := nd.date.UnmarshalText(text)
	nd.isNotNull = err == nil

	return err
}

func (nd NullDate) MarshalJSON() ([]byte, error) {
	if nd.isNotNull {
		return nd.date.MarshalJSON()
	}

	return []byte("null"), nil
}

func (nd *NullDate) UnmarshalJSON(json []byte) error {
	if bytes.Equal(json, []byte("null")) {
		*nd = NullDate{}

		return nil
	}

	err := nd.date.UnmarshalJSON(json)
	nd.isNotNull = err == nil

	return err
}
