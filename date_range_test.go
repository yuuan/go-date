package date

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Factory functions
// --------------------------------------------------

func TestNewDateRange(t *testing.T) {
	tests := []struct {
		start   Date
		end     Date
		want    DateRange
		wantErr error
	}{
		{
			MustParse("2024-06-01"),
			MustParse("2024-06-02"),
			MustNewDateRange(MustParse("2024-06-01"), MustParse("2024-06-02")),
			nil,
		},
		{
			ZeroDate(),
			ZeroDate(),
			ZeroDateRange(),
			nil,
		},
		{
			FromTime(time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)),
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.Local)),
			ZeroDateRange(),
			ErrDifferentTimeZone,
		},
		{
			ZeroDate(),
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC)),
			ZeroDateRange(),
			ErrOnlyOneSideIsZero,
		},
		{
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC)),
			ZeroDate(),
			ZeroDateRange(),
			ErrOnlyOneSideIsZero,
		},
		{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			ZeroDateRange(),
			ErrEndDateIsBeforeStartDate,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NewDateRange(Date{value:"%s",location:"%s"},Date{value:"%s",location:"%s"})`,
			tt.start,
			tt.start.Location(),
			tt.end,
			tt.end.Location(),
		)

		t.Run(testcase, func(t *testing.T) {
			r, err := NewDateRange(tt.start, tt.end)

			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
				assert.Equal(t, tt.want.String(), r.String())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			}
		})
	}
}

func TestMustNewDateRange(t *testing.T) {
	tests := []struct {
		start     Date
		end       Date
		want      DateRange
		wantPanic bool
	}{
		{
			MustParse("2024-06-01"),
			MustParse("2024-06-02"),
			MustNewDateRange(MustParse("2024-06-01"), MustParse("2024-06-02")),
			false,
		},
		{
			ZeroDate(),
			ZeroDate(),
			ZeroDateRange(),
			false,
		},
		{
			FromTime(time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)),
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.Local)),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDate(),
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC)),
			ZeroDateRange(),
			true,
		},
		{
			FromTime(time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC)),
			ZeroDate(),
			ZeroDateRange(),
			true,
		},
		{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NewDateRange(Date{value:"%s",location:"%s"},Date{value:"%s",location:"%s"})`,
			tt.start,
			tt.start.Location(),
			tt.end,
			tt.end.Location(),
		)

		t.Run(testcase, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { MustNewDateRange(tt.start, tt.end) }, "Did not panic")
			} else {
				assert.Equal(t, tt.want.String(), MustNewDateRange(tt.start, tt.end).String())
			}
		})
	}
}

func TestZeroDateRange(t *testing.T) {
	t.Run("ZeroDateRange()", func(t *testing.T) {
		r := ZeroDateRange()
		assert.True(t, r.IsZero(), "DateRange created was not zero value")
	})
}

func TestParseDateRange(t *testing.T) {
	tests := []struct {
		start   string
		end     string
		want    DateRange
		wantErr bool
	}{
		{
			"2024-06-05",
			"2024-06-05",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"invalid",
			"2024-06-05",
			ZeroDateRange(),
			true,
		},
		{
			"2024-06-05",
			"invalid",
			ZeroDateRange(),
			true,
		},
		{
			"invalid",
			"invaiid",
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`ParseDateRange("%s","%s")`, tt.start, tt.end)

		t.Run(testcase, func(t *testing.T) {
			r, err := ParseDateRange(tt.start, tt.end)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			}
		})
	}
}

func TestMustParseDateRange(t *testing.T) {
	tests := []struct {
		start     string
		end       string
		want      DateRange
		wantPanic bool
	}{
		{
			"2024-06-05",
			"2024-06-05",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"invalid",
			"2024-06-05",
			ZeroDateRange(),
			true,
		},
		{
			"2024-06-05",
			"invalid",
			ZeroDateRange(),
			true,
		},
		{
			"invalid",
			"invaiid",
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`ParseDateRange("%s","%s")`, tt.start, tt.end)

		t.Run(testcase, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { MustParseDateRange(tt.start, tt.end) }, "Did not panic")
			} else {
				assert.Equal(t, tt.want.String(), MustParseDateRange(tt.start, tt.end).String())
			}
		})
	}
}

func TestCustomParseDateRange(t *testing.T) {
	tests := []struct {
		layout  string
		start   string
		end     string
		want    DateRange
		wantErr bool
	}{
		{
			"2006-01-02",
			"2024-06-05",
			"2024-06-05",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"02/01/2006",
			"05/06/2024",
			"05/06/2024",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"02/01/2006",
			"invalid",
			"05/06/2024",
			ZeroDateRange(),
			true,
		},
		{
			"02/01/2006",
			"05/06/2024",
			"invalid",
			ZeroDateRange(),
			true,
		},
		{
			"02/01/2006",
			"invalid",
			"invalid",
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`ParseDateRange("%s","%s")`, tt.start, tt.end)

		t.Run(testcase, func(t *testing.T) {
			r, err := CustomParseDateRange(tt.layout, tt.start, tt.end)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			}
		})
	}
}

func TestMustCustomParseDateRange(t *testing.T) {
	tests := []struct {
		layout    string
		start     string
		end       string
		want      DateRange
		wantPanic bool
	}{
		{
			"2006-01-02",
			"2024-06-05",
			"2024-06-05",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"02/01/2006",
			"05/06/2024",
			"05/06/2024",
			MustNewDateRange(MustParse("2024-06-05"), MustParse("2024-06-05")),
			false,
		},
		{
			"02/01/2006",
			"invalid",
			"05/06/2024",
			ZeroDateRange(),
			true,
		},
		{
			"02/01/2006",
			"05/06/2024",
			"invalid",
			ZeroDateRange(),
			true,
		},
		{
			"02/01/2006",
			"invalid",
			"invalid",
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`ParseDateRange("%s","%s")`, tt.start, tt.end)

		t.Run(testcase, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { MustCustomParseDateRange(tt.layout, tt.start, tt.end) }, "Did not panic")
			} else {
				assert.Equal(t, tt.want.String(), MustCustomParseDateRange(tt.layout, tt.start, tt.end).String())
			}
		})
	}
}

// Determination methods
// --------------------------------------------------

func TestDateRangeIsZero(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want bool
	}{
		{MustParseDateRange("2024-06-05", "2024-06-05"), false},
		{ZeroDateRange(), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRange{"%s","%s"}.IsZero()`, tt.dr.start, tt.dr.end)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.IsZero())
		})
	}
}

func TestDateRangeOnlyOneDay(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want bool
	}{
		{MustParseDateRange("2024-06-05", "2024-06-05"), true},
		{MustParseDateRange("2024-06-05", "2024-06-06"), false},
		{ZeroDateRange(), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRange{"%s","%s"}.OnlyoneDay()`, tt.dr.start, tt.dr.end)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.OnlyOneDay())
		})
	}
}

// Comparison methods
// --------------------------------------------------

func TestDateRangeEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-06", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.Equal(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.Equal(tt.drB))
		})
	}
}

func TestDateRangeNotEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-06", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.NotEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.NotEqual(tt.drB))
		})
	}
}

func TestDateRangeStartsOn(t *testing.T) {
	tests := []struct {
		dr   DateRange
		date Date
		want bool
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-01"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-06"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDate(),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDate(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsOn(Date{"%s"})`,
			tt.dr.start,
			tt.dr.end,
			tt.date,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.StartsOn(tt.date))
		})
	}
}

func TestDateRangeEndsOn(t *testing.T) {
	tests := []struct {
		dr   DateRange
		date Date
		want bool
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-01"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-06"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDate(),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDate(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsOn(Date{"%s"})`,
			tt.dr.start,
			tt.dr.end,
			tt.date,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.EndsOn(tt.date))
		})
	}
}

func TestDateRangeStartsOnSameDate(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsOnSameDate(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.StartsOnSameDate(tt.drB))
		})
	}
}

func TestDateRangeStartsBefore(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsBefore(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.StartsBefore(tt.drB))
		})
	}
}

func TestDateRangeStartsBeforeOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsBeforeOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.StartsBeforeOrEqual(tt.drB))
		})
	}
}

func TestDateRangeStartsAfter(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsAfter(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.StartsAfter(tt.drB))
		})
	}
}

func TestDateRangeStartsAfterOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.StartsAfterOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.StartsAfterOrEqual(tt.drB))
		})
	}
}

func TestDateRangeEndsOnSameDate(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsOnSameDate(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.EndsOnSameDate(tt.drB))
		})
	}
}

func TestDateRangeEndsBefore(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsBefore(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.EndsBefore(tt.drB))
		})
	}
}

func TestDateRangeEndsBeforeOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsBeforeOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.EndsBeforeOrEqual(tt.drB))
		})
	}
}

func TestDateRangeEndsAfter(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsAfter(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.EndsAfter(tt.drB))
		})
	}
}

func TestDateRangeEndsAfterOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.EndsAfterOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.EndsAfterOrEqual(tt.drB))
		})
	}
}

func TestDateRangeOverlapsWith(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.OverlapsWith(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.OverlapsWith(tt.drB))
		})
	}
}

func TestDateRangeContains(t *testing.T) {
	tests := []struct {
		dr   DateRange
		date Date
		want bool
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-01"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-04"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-06"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDate(),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDate(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.Contains(Date{"%s"})`,
			tt.dr.start,
			tt.dr.end,
			tt.date,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.Contains(tt.date))
		})
	}
}

func TestDateRangeLessThan(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.LessThan(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.LessThan(tt.drB))
		})
	}
}

func TestDateRangeLessThanOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			false,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.LessThanOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.LessThanOrEqual(tt.drB))
		})
	}
}

func TestDateRangeGreaterThan(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.GreaterThan(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.GreaterThan(tt.drB))
		})
	}
}

func TestDateRangeGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		drA  DateRange
		drB  DateRange
		want bool
	}{
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-06", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-01", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			false,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			true,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			true,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			false,
		},
		{
			ZeroDateRange(),
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.GreaterThanOrEqual(DateRange{"%s","%s"})`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.drA.GreaterThanOrEqual(tt.drB))
		})
	}
}

// Conversion methods
// --------------------------------------------------

func TestDateRangeStart(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want Date
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-05"),
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-01"),
		},
		{
			ZeroDateRange(),
			ZeroDate(),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.Start())`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.Start())
		})
	}
}

func TestDateRangeEnd(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want Date
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParse("2024-06-05"),
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParse("2024-06-10"),
		},
		{
			ZeroDateRange(),
			ZeroDate(),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.End())`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.End())
		})
	}
}

func TestDateRangeLocation(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want *time.Location
	}{
		{
			MustParseDateRange("2024-06-01", "2024-06-01"),
			time.Local,
		},
		{
			MustNewDateRange(
				FromTime(time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local)),
				FromTime(time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local)),
			),
			time.Local,
		},
		{
			MustNewDateRange(
				FromTime(time.Date(2024, time.June, 5, 0, 0, 0, 0, time.UTC)),
				FromTime(time.Date(2024, time.June, 5, 0, 0, 0, 0, time.UTC)),
			),
			time.UTC,
		},
		{
			ZeroDateRange(),
			time.UTC,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{Date{value:"%s",location:"%s"},Date{value:"%s",location:"%s"}}.Location())`,
			tt.dr.start,
			tt.dr.start.Location(),
			tt.dr.end,
			tt.dr.end.Location(),
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.Location())
		})
	}
}

func TestDateRangeDays(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want int
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			1,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			10,
		},
		{
			ZeroDateRange(),
			1,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.Days())`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.Days())
		})
	}
}

func TestDateRangeGetOverlapping(t *testing.T) {
	tests := []struct {
		drA     DateRange
		drB     DateRange
		want    DateRange
		wantErr error
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			nil,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			nil,
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			nil,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-06"),
			MustParseDateRange("2024-06-01", "2024-06-10"),
			MustParseDateRange("2024-06-05", "2024-06-06"),
			nil,
		},
		{
			MustParseDateRange("2024-06-04", "2024-06-04"),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			ErrRangesDontOverlap,
		},
		{
			ZeroDateRange(),
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			ErrRangesDontOverlap,
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			ZeroDateRange(),
			ZeroDateRange(),
			ErrRangesDontOverlap,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange("%s","%s"}.GetOverlapping(DateRange{"%s","%s"}))`,
			tt.drA.start,
			tt.drA.end,
			tt.drB.start,
			tt.drB.end,
		)

		t.Run(testcase, func(t *testing.T) {
			r, err := tt.drA.GetOverlapping(tt.drB)

			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
				assert.Equal(t, tt.want.String(), r.String())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.String(), r.String())
			}
		})
	}
}

func TestDateRangeDates(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want Dates
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			Dates{MustParse("2024-06-05")},
		},
		{
			MustParseDateRange("2024-06-05", "2024-06-07"),
			Dates{MustParse("2024-06-05"), MustParse("2024-06-06"), MustParse("2024-06-07")},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.Dates()`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			for i, d := range tt.dr.Dates() {
				assert.Equal(t, d, tt.want[i])
			}
		})
	}
}

func TestDateRangeString(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want string
	}{
		{
			MustParseDateRange("2024-06-05", "2024-06-05"),
			"2024-06-05/2024-06-05",
		},
		{
			MustParseDateRange("2024-06-01", "2024-06-10"),
			"2024-06-01/2024-06-10",
		},
		{
			ZeroDateRange(),
			"0001-01-01/0001-01-01",
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.End())`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dr.String())
		})
	}
}

func TestDateRangeMarshalJSON(t *testing.T) {
	tests := []struct {
		dr   DateRange
		want string
	}{
		{
			MustParseDateRange("2024-01-15", "2024-02-20"),
			`{"start":"2024-01-15","end":"2024-02-20"}`,
		},
		{
			MustParseDateRange("2024-03-01", "2024-03-01"),
			`{"start":"2024-03-01","end":"2024-03-01"}`,
		},
		{
			MustParseDateRange("2023-12-31", "2024-01-01"),
			`{"start":"2023-12-31","end":"2024-01-01"}`,
		},
		{
			ZeroDateRange(),
			`{"start":"0001-01-01","end":"0001-01-01"}`,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"}.MarshalJSON()`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			result, err := json.Marshal(tt.dr)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.want, string(result))
		})
	}
}

func TestDateRangeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		json    string
		want    DateRange
		wantErr bool
	}{
		{
			`{"start":"2024-01-15","end":"2024-02-20"}`,
			MustParseDateRange("2024-01-15", "2024-02-20"),
			false,
		},
		{
			`{"start":"2024-03-01","end":"2024-03-01"}`,
			MustParseDateRange("2024-03-01", "2024-03-01"),
			false,
		},
		{
			`{"start":"2023-12-31","end":"2024-01-01"}`,
			MustParseDateRange("2023-12-31", "2024-01-01"),
			false,
		},
		{
			`{"start":"0001-01-01","end":"0001-01-01"}`,
			ZeroDateRange(),
			false,
		},
		{
			`{"start":"invalid","end":"2024-02-20"}`,
			ZeroDateRange(),
			true,
		},
		{
			`{"start":"2024-01-15","end":"invalid"}`,
			ZeroDateRange(),
			true,
		},
		{
			`{"start":"2024-02-20","end":"2024-01-15"}`,
			ZeroDateRange(),
			true,
		},
		{
			`{"start":"2024-01-15"`,
			ZeroDateRange(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`UnmarshalJSON(%s)`, tt.json)

		t.Run(testcase, func(t *testing.T) {
			var dr DateRange
			err := json.Unmarshal([]byte(tt.json), &dr)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, dr)
			}
		})
	}
}

func TestDateRangeJSONRoundTrip(t *testing.T) {
	tests := []struct {
		dr DateRange
	}{
		{MustParseDateRange("2024-01-15", "2024-02-20")},
		{MustParseDateRange("2024-05-05", "2024-05-05")},
		{MustParseDateRange("2023-12-31", "2024-01-01")},
		{MustParseDateRange("2020-01-01", "2025-12-31")},
		{ZeroDateRange()},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`DateRange{"%s","%s"} JSON round trip`,
			tt.dr.start,
			tt.dr.end,
		)

		t.Run(testcase, func(t *testing.T) {
			jsonData, err := json.Marshal(tt.dr)
			assert.NoError(t, err)

			var result DateRange
			err = json.Unmarshal(jsonData, &result)
			assert.NoError(t, err)

			assert.Equal(t, tt.dr, result)
		})
	}
}
