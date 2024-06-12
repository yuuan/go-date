package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Factory functions
// --------------------------------------------------

func TestNewDate(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		day   int
		want  string
	}{
		{2024, time.June, 5, "2024-06-05"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("TestNewDate(%d, %d, %d)", tt.year, int(tt.month), tt.day)

		t.Run(testcase, func(t *testing.T) {
			date := NewDate(tt.year, tt.month, tt.day)

			assert.Equal(t, tt.want, date.String())
		})
	}
}

func TestZeroDate(t *testing.T) {
	t.Run("ZeroDate()", func(t *testing.T) {
		date := ZeroDate()
		assert.True(t, date.IsZero(), "Date created was not zero value")
	})
}

func TestFromTime(t *testing.T) {
	tests := []struct {
		time time.Time
		want string
	}{
		{time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local), "2024-06-05"},
		{time.Date(2024, time.June, 5, 12, 30, 10, 10, time.Local), "2024-06-05"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`FromTime(Time{%s})`, tt.time.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			date := FromTime(tt.time)
			assert.Equal(t, tt.want, date.String(), "Expected %s, got %s", tt.want, date.String())
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"2024-06-05", "2024-06-05"},

		{"05/06/2024", "error"},
		{"Wed, 05 Jun 2024", "error"},
		{"invalid-date", "error"},
		{"2024/02/31", "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Parse("%s")`, tt.value)

		t.Run(testcase, func(t *testing.T) {
			date, err := Parse(tt.value)
			if tt.want == "error" {
				assert.ErrorContains(t, err, "Unable to parse")
			} else {
				assert.Nil(t, err, "An unexpected error was returned")
				assert.Equal(t, tt.want, date.String(), "The parsed date does not match the expected date")
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"2024-06-05", "2024-06-05"},

		{"05/06/2024", "panic"},
		{"Wed, 05 Jun 2024", "panic"},
		{"invalid-date", "panic"},
		{"2024/02/31", "panic"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`MustParse("%s")`, tt.value)

		t.Run(testcase, func(t *testing.T) {
			if tt.want == "panic" {
				assert.Panics(t, func() { MustParse(tt.value) }, "Did not panic")
			} else {
				date := MustParse(tt.value)
				assert.Equal(t, tt.want, date.String(), "The parsed date does not match the expected date")
			}
		})
	}
}

func TestCustomParse(t *testing.T) {
	tests := []struct {
		layout string
		value  string
		want   string
	}{
		{"2006-01-02", "2024-06-05", "2024-06-05"},
		{"02/01/2006", "05/06/2024", "2024-06-05"},
		{"Mon, 02 Jan 2006", "Wed, 05 Jun 2024", "2024-06-05"},

		{"2006-01-02", "invalid-date", "error"},
		{"02/01/2006", "31/02/2024", "error"},
		{"Mon, 02 Jan 2006", "Invalid, 05 Jun 2024", "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Parse("%s","%s")`, tt.layout, tt.value)

		t.Run(testcase, func(t *testing.T) {
			date, err := CustomParse(tt.layout, tt.value)
			if tt.want == "error" {
				assert.ErrorContains(t, err, "Unable to parse")
			} else {
				assert.Nil(t, err, "An unexpected error was returned")
				assert.Equal(t, tt.want, date.String(), "The parsed date does not match the expected date")
			}
		})
	}
}

func TestMustCustomParse(t *testing.T) {
	tests := []struct {
		layout string
		value  string
		want   string
	}{
		{"2006-01-02", "2024-06-05", "2024-06-05"},
		{"02/01/2006", "05/06/2024", "2024-06-05"},
		{"Mon, 02 Jan 2006", "Wed, 05 Jun 2024", "2024-06-05"},

		{"2006-01-02", "invalid-date", "panic"},
		{"02/01/2006", "31/02/2024", "panic"},
		{"Mon, 02 Jan 2006", "Invalid, 05 Jun 2024", "panic"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`MustParse("%s","%s")`, tt.layout, tt.value)

		t.Run(testcase, func(t *testing.T) {
			if tt.want == "panic" {
				assert.Panics(t, func() { MustCustomParse(tt.layout, tt.value) }, "Did not panic")
			} else {
				date := MustCustomParse(tt.layout, tt.value)
				assert.Equal(t, tt.want, date.String(), "The parsed date does not match the expected date")
			}
		})
	}
}

func TestToday(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local), "2024-06-05"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-06-05"},
		{time.Date(2024, time.June, 5, 23, 59, 59, 999, time.Local), "2024-06-05"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Today() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, Today().String())
		})
	}
}

func TestYesterday(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local), "2024-06-04"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-06-04"},
		{time.Date(2024, time.June, 5, 23, 59, 59, 999, time.Local), "2024-06-04"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Yesterday() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, Yesterday().String())
		})
	}
}

func TestTomorrow(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local), "2024-06-06"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-06-06"},
		{time.Date(2024, time.June, 5, 23, 59, 59, 999, time.Local), "2024-06-06"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Tomorrow() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, Tomorrow().String())
		})
	}
}

// Determination methods
// --------------------------------------------------

func TestIsZero(t *testing.T) {
	tests := []struct {
		testcase string
		date     func() Date
		want     bool
	}{
		{"ZeroDate().IsZero()", ZeroDate, true},
		{"Date{}.IsZero()", func() Date { return Date{} }, true},
		{"Today().IsZero()", Today, false},
	}

	for _, tt := range tests {
		t.Run(tt.testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date().IsZero())
		})
	}
}

func TestIsFirstOfMonth(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-01"), true},
		{MustParse("2024-06-02"), false},
		{MustParse("2024-06-30"), false},
		{MustParse("2024-02-29"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsFirstOfMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsFirstOfMonth())
		})
	}
}

func TestIsLastOfMonth(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-01"), false},
		{MustParse("2024-06-02"), false},
		{MustParse("2024-06-30"), true},
		{MustParse("2024-02-29"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsLastOfMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsLastOfMonth())
		})
	}
}

func TestIsMonday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), true},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsMonday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsMonday())
		})
	}
}

func TestIsTuesday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), true},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsTuesday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsTuesday())
		})
	}
}

func TestIsWednesday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), true},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsWednesday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsWednesday())
		})
	}
}

func TestIsThursday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), true},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsThursday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsThursday())
		})
	}
}

func TestIsFriday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), true},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsFriday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsFriday())
		})
	}
}

func TestIsSaturday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), true},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsSaturday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsSaturday())
		})
	}
}

func TestIsSunday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsSunday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsSunday())
		})
	}
}

func TestIsWeekday(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), true},
		{MustParse("2024-06-11"), true},
		{MustParse("2024-06-12"), true},
		{MustParse("2024-06-13"), true},
		{MustParse("2024-06-14"), true},
		{MustParse("2024-06-15"), false},
		{MustParse("2024-06-16"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsWeekday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsWeekday())
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-10"), false},
		{MustParse("2024-06-11"), false},
		{MustParse("2024-06-12"), false},
		{MustParse("2024-06-13"), false},
		{MustParse("2024-06-14"), false},
		{MustParse("2024-06-15"), true},
		{MustParse("2024-06-16"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsWeekend()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsWeekend())
		})
	}
}

func TestIsPast(t *testing.T) {
	tests := []struct {
		now  time.Time
		date func() Date
		want bool
	}{
		{
			time.Date(2024, time.June, 3, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 4, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 6, 23, 59, 59, 999999999, time.Local),
			func() Date { return MustParse("2024-06-06") },
			false,
		},
		{
			time.Date(2024, time.June, 6, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			true,
		},
		{
			time.Date(2024, time.June, 7, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			true,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return ZeroDate() },
			true,
		},
		{
			time.Date(2024, time.June, 6, 0, 0, 0, 0, time.Local),
			func() Date { return Today() },
			false,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return Yesterday() },
			true,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return Tomorrow() },
			false,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		date := tt.date()

		testcase := fmt.Sprintf(`Date{"%s"}.IsPast() at %s`, date, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, date.IsPast())
		})
	}
}

func TestIsFuture(t *testing.T) {
	tests := []struct {
		now  time.Time
		date func() Date
		want bool
	}{
		{
			time.Date(2024, time.June, 3, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			true,
		},
		{
			time.Date(2024, time.June, 4, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			true,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 6, 23, 59, 59, 999999999, time.Local),
			func() Date { return MustParse("2024-06-06") },
			false,
		},
		{
			time.Date(2024, time.June, 6, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 7, 0, 0, 0, 0, time.Local),
			func() Date { return MustParse("2024-06-05") },
			false,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return ZeroDate() },
			false,
		},
		{
			time.Date(2024, time.June, 6, 0, 0, 0, 0, time.Local),
			func() Date { return Today() },
			false,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return Yesterday() },
			false,
		},
		{
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
			func() Date { return Tomorrow() },
			true,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		date := tt.date()

		testcase := fmt.Sprintf(`Date{"%s"}.IsFuture() at %s`, date, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, date.IsFuture())
		})
	}
}

func TestIsToday(t *testing.T) {
	now := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	SetTestNow(func() time.Time { return now })
	defer ResetTestNow()

	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-03"), false},
		{MustParse("2024-06-04"), false},
		{MustParse("2024-06-05"), true},
		{MustParse("2024-06-06"), false},
		{MustParse("2024-06-07"), false},
		{ZeroDate(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsToday() at %s`, tt.date, now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsToday())
		})
	}
}

func TestIsTomorrow(t *testing.T) {
	now := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	SetTestNow(func() time.Time { return now })
	defer ResetTestNow()

	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-03"), false},
		{MustParse("2024-06-04"), false},
		{MustParse("2024-06-05"), false},
		{MustParse("2024-06-06"), true},
		{MustParse("2024-06-07"), false},
		{ZeroDate(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsTomorrow() at %s`, tt.date, now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsTomorrow())
		})
	}
}

func TestIsYesterday(t *testing.T) {
	now := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	SetTestNow(func() time.Time { return now })
	defer ResetTestNow()

	tests := []struct {
		date Date
		want bool
	}{
		{MustParse("2024-06-03"), false},
		{MustParse("2024-06-04"), true},
		{MustParse("2024-06-05"), false},
		{MustParse("2024-06-06"), false},
		{MustParse("2024-06-07"), false},
		{ZeroDate(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.IsYesterday() at %s`, tt.date, now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.IsYesterday())
		})
	}
}

// Comparison methods
// --------------------------------------------------

func TestCompare(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  int
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), -1},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), 0},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), 1},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Compare(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.Compare(tt.dateB))
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), false},
		{ZeroDate(), ZeroDate(), true},
		{ZeroDate(), MustParse("2024-06-05"), false},
		{MustParse("2024-06-05"), ZeroDate(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Equal(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.Equal(tt.dateB))
		})
	}
}

func TestNotEqual(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), true},
		{ZeroDate(), ZeroDate(), false},
		{ZeroDate(), MustParse("2024-06-05"), true},
		{MustParse("2024-06-05"), ZeroDate(), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.NotEqual(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.NotEqual(tt.dateB))
		})
	}
}

func TestAfter(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.After(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.After(tt.dateB))
		})
	}
}

func TestAfterOrEqual(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AfterOrEqual(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.AfterOrEqual(tt.dateB))
		})
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), false},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Before(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.Before(tt.dateB))
		})
	}
}

func TestBeforeOrEqual(t *testing.T) {
	tests := []struct {
		dateA Date
		dateB Date
		want  bool
	}{
		{MustParse("2024-06-04"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-05"), MustParse("2024-06-05"), true},
		{MustParse("2024-06-06"), MustParse("2024-06-05"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.BeforeOrEqual(Date{"%s"})`, tt.dateA, tt.dateB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dateA.BeforeOrEqual(tt.dateB))
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		date    Date
		start   Date
		end     Date
		want    bool
		wantErr error
	}{
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-04"),
			MustParse("2024-06-06"),
			true,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-04"),
			MustParse("2024-06-05"),
			true,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-05"),
			MustParse("2024-06-06"),
			true,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-05"),
			MustParse("2024-06-05"),
			true,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-03"),
			MustParse("2024-06-04"),
			false,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-06"),
			MustParse("2024-06-07"),
			false,
			nil,
		},
		{
			MustParse("2024-06-05"),
			MustParse("2024-06-06"),
			MustParse("2024-06-04"),
			false,
			ErrEndDateIsBeforeStartDate,
		},
		{
			MustParse("2024-06-05"),
			FromTime(time.Date(2004, time.June, 5, 0, 0, 0, 0, time.UTC)),
			FromTime(time.Date(2004, time.June, 5, 0, 0, 0, 0, time.Local)),
			false,
			ErrDifferentTimeZone,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`Date{"%s"}.Between(Date{"%s"},Date{"%s"})`,
			tt.date,
			tt.start,
			tt.end,
		)

		t.Run(testcase, func(t *testing.T) {
			b, err := tt.date.Between(tt.start, tt.end)
			assert.Equal(t, tt.want, b)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
			}
		})
	}
}

// Addition and Subtraction methods
// --------------------------------------------------

func TestAddDay(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-06-02"},
		{MustParse("2024-06-30"), "2024-07-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddDay()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddDay().String())
		})
	}
}

func TestAddDays(t *testing.T) {
	tests := []struct {
		date Date
		days int
		want string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-06-02"},
		{MustParse("2024-06-01"), -1, "2024-05-31"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddDays(%d)`, tt.date, tt.days)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddDays(tt.days).String())
		})
	}
}

func TestSubDay(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-05-31"},
		{MustParse("2024-06-02"), "2024-06-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubDay()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubDay().String())
		})
	}
}

func TestSubDays(t *testing.T) {
	tests := []struct {
		date Date
		days int
		want string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-05-31"},
		{MustParse("2024-06-01"), -1, "2024-06-02"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubDays(%d)`, tt.date, tt.days)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubDays(tt.days).String())
		})
	}
}

func TestAddWeek(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-06-08"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddWeek()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddWeek().String())
		})
	}
}

func TestAddWeeks(t *testing.T) {
	tests := []struct {
		date  Date
		weeks int
		want  string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-06-08"},
		{MustParse("2024-06-01"), -1, "2024-05-25"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddWeeks(%d)`, tt.date, tt.weeks)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddWeeks(tt.weeks).String())
		})
	}
}

func TestSubWeek(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-05-25"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubWeek()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubWeek().String())
		})
	}
}

func TestSubWeeks(t *testing.T) {
	tests := []struct {
		date  Date
		weeks int
		want  string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-05-25"},
		{MustParse("2024-06-01"), -1, "2024-06-08"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubWeeks(%d)`, tt.date, tt.weeks)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubWeeks(tt.weeks).String())
		})
	}
}

func TestAddMonth(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-07-01"},
		{MustParse("2024-05-31"), "2024-07-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddMonth().String())
		})
	}
}

func TestAddMonths(t *testing.T) {
	tests := []struct {
		date   Date
		months int
		want   string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-07-01"},
		{MustParse("2024-05-31"), 1, "2024-07-01"},
		{MustParse("2024-06-01"), -1, "2024-05-01"},
		{MustParse("2024-05-31"), -1, "2024-05-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddMonths(%d)`, tt.date, tt.months)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddMonths(tt.months).String())
		})
	}
}

func TestSubMonth(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-05-01"},
		{MustParse("2024-05-31"), "2024-05-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubMonth().String())
		})
	}
}

func TestSubMonths(t *testing.T) {
	tests := []struct {
		date   Date
		months int
		want   string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2024-05-01"},
		{MustParse("2024-05-31"), 1, "2024-05-01"},
		{MustParse("2024-06-01"), -1, "2024-07-01"},
		{MustParse("2024-05-31"), -1, "2024-07-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubMonths(%d)`, tt.date, tt.months)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubMonths(tt.months).String())
		})
	}
}

func TestAddYear(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-01-01"), "2025-01-01"},
		{MustParse("2024-02-29"), "2025-03-01"},
		{MustParse("2024-06-01"), "2025-06-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddYear()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddYear().String())
		})
	}
}

func TestAddYears(t *testing.T) {
	tests := []struct {
		date  Date
		years int
		want  string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2025-06-01"},
		{MustParse("2024-02-29"), 1, "2025-03-01"},
		{MustParse("2024-06-01"), -1, "2023-06-01"},
		{MustParse("2024-02-29"), -1, "2023-03-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.AddYears(%d)`, tt.date, tt.years)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.AddYears(tt.years).String())
		})
	}
}

func TestSubYear(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-01-01"), "2023-01-01"},
		{MustParse("2024-02-29"), "2023-03-01"},
		{MustParse("2024-06-01"), "2023-06-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubYear()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubYear().String())
		})
	}
}

func TestSubYears(t *testing.T) {
	tests := []struct {
		date  Date
		years int
		want  string
	}{
		{MustParse("2024-06-01"), 0, "2024-06-01"},
		{MustParse("2024-06-01"), 1, "2023-06-01"},
		{MustParse("2024-02-29"), 1, "2023-03-01"},
		{MustParse("2024-06-01"), -1, "2025-06-01"},
		{MustParse("2024-02-29"), -1, "2025-03-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.SubYears(%d)`, tt.date, tt.years)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.SubYears(tt.years).String())
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-06-01"},
		{MustParse("2024-06-05"), "2024-06-01"},
		{MustParse("2024-02-01"), "2024-02-01"},
		{MustParse("2024-02-29"), "2024-02-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.StartOfMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.StartOfMonth().String())
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-01"), "2024-06-30"},
		{MustParse("2024-06-05"), "2024-06-30"},
		{MustParse("2024-02-01"), "2024-02-29"},
		{MustParse("2024-02-29"), "2024-02-29"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.EndOfMonth()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.EndOfMonth().String())
		})
	}
}

func TestStartOfYear(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-01-01"), "2024-01-01"},
		{MustParse("2024-06-05"), "2024-01-01"},
		{MustParse("2024-12-31"), "2024-01-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.StartOfYear()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.StartOfYear().String())
		})
	}
}

func TestEndOfYear(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-01-01"), "2024-12-31"},
		{MustParse("2024-06-05"), "2024-12-31"},
		{MustParse("2024-12-31"), "2024-12-31"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.EndOfYear()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.EndOfYear().String())
		})
	}
}

// Conversion methods
// --------------------------------------------------

func TestTime(t *testing.T) {
	tests := []struct {
		date Date
		want time.Time
	}{
		{
			MustParse("2024-06-05"),
			time.Date(2024, time.June, 5, 0, 0, 0, 0, time.Local),
		},
		{
			FromTime(time.Date(2024, time.June, 6, 12, 34, 56, 789, time.Local)),
			time.Date(2024, time.June, 6, 0, 0, 0, 0, time.Local),
		},
		{
			ZeroDate(),
			time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Time()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Time())
		})
	}
}

func TestYear(t *testing.T) {
	tests := []struct {
		date Date
		want int
	}{
		{MustParse("2024-06-05"), 2024},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Year()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Year())
		})
	}
}

func TestMonth(t *testing.T) {
	tests := []struct {
		date Date
		want time.Month
	}{
		{MustParse("2024-06-05"), time.June},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Month()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Month())
		})
	}
}

func TestDay(t *testing.T) {
	tests := []struct {
		date Date
		want int
	}{
		{MustParse("2024-06-05"), 5},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Day()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Day())
		})
	}
}

func TestYearDay(t *testing.T) {
	tests := []struct {
		date Date
		want int
	}{
		{MustParse("2024-01-01"), 1},
		{MustParse("2024-06-05"), 157},
		{MustParse("2024-12-31"), 366},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.YearDay()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.YearDay())
		})
	}
}

func TestISOWeek(t *testing.T) {
	tests := []struct {
		date     Date
		wantYear int
		wantWeek int
	}{
		{MustParse("2024-01-01"), 2024, 1},
		{MustParse("2024-06-05"), 2024, 23},
		{MustParse("2024-12-31"), 2025, 1},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.ISOWeek()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			year, week := tt.date.ISOWeek()
			assert.Equal(t, tt.wantYear, year, "Year of ISOWeek is not match")
			assert.Equal(t, tt.wantWeek, week, "Week of ISOWeek is not match")
		})
	}
}

func TestWeekday(t *testing.T) {
	tests := []struct {
		date Date
		want time.Weekday
	}{
		{MustParse("2024-06-10"), time.Monday},
		{MustParse("2024-06-11"), time.Tuesday},
		{MustParse("2024-06-12"), time.Wednesday},
		{MustParse("2024-06-13"), time.Thursday},
		{MustParse("2024-06-14"), time.Friday},
		{MustParse("2024-06-15"), time.Saturday},
		{MustParse("2024-06-16"), time.Sunday},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Weekday()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Weekday())
		})
	}
}

func TestLocation(t *testing.T) {
	tests := []struct {
		location *time.Location
	}{
		{time.UTC},
		{time.Local},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{location:"%s"}.Location()`, tt.location)

		t.Run(testcase, func(t *testing.T) {
			SetTestLocation(func() *time.Location { return tt.location })
			defer ResetTestLocation()

			assert.Equal(t, tt.location, MustParse("2024-06-05").Location())
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		date   Date
		format string
		want   string
	}{
		{MustParse("2024-06-05"), "2006-01-02", "2024-06-05"},
		{MustParse("2024-06-05"), "02/01/2006", "05/06/2024"},
		{MustParse("2024-06-05"), "Mon, 02 Jan 2006", "Wed, 05 Jun 2024"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Format("%s")`, tt.date, tt.format)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.Format(tt.format))
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		date      Date
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{MustParse("2024-06-05"), 2024, time.June, 5},
		{ZeroDate(), 1, time.January, 1},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Split()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			year, month, day := tt.date.Split()
			assert.Equal(t, tt.wantYear, year, "Year is not match")
			assert.Equal(t, tt.wantMonth, month, "Month is not match")
			assert.Equal(t, tt.wantDay, day, "Day is not match")
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-05"), "2024-06-05"},
		{ZeroDate(), "0001-01-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.String()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.date.String())
		})
	}
}

func TestStringPtr(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-05"), "2024-06-05"},
		{ZeroDate(), "0001-01-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.StringPtr()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, *tt.date.StringPtr())
		})
	}
}

// Marshalling methods
// --------------------------------------------------

func TestValue(t *testing.T) {
	tests := []struct {
		date      Date
		wantValue string
		wantErr   error
	}{
		{MustParse("2024-06-05"), "2024-06-05", nil},
		{ZeroDate(), "0001-01-01", nil},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.Value()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			value, err := tt.date.Value()
			assert.Equal(t, tt.wantValue, value)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestScan(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"2024-06-05", "2024-06-05"},
		{"", "0001-01-01"},
		{"invalid", "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{}.Scan("%s")`, tt.value)

		t.Run(testcase, func(t *testing.T) {
			d := ZeroDate()
			err := d.Scan(tt.value)
			if tt.want == "error" {
				assert.Error(t, err, "Unable to parse")
				assert.True(t, d.IsZero(), "date is not zero")
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
				assert.Equal(t, tt.want, d.String())
			}
		})
	}
}

func TestMarshalText(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-05"), "2024-06-05"},
		{ZeroDate(), "0001-01-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.MarshalText()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			text, err := tt.date.MarshalText()
			assert.Nil(t, err, "Excepted no error, got %v", err)
			assert.Equal(t, tt.want, string(text))
		})
	}
}

func TestUnmarshalText(t *testing.T) {
	tests := []struct {
		text []byte
		want string
	}{
		{[]byte("2024-06-05"), "2024-06-05"},
		{[]byte{}, "error"},
		{[]byte("invalid"), "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{}.UnmarshalText("%s")`, tt.text)

		t.Run(testcase, func(t *testing.T) {
			d := ZeroDate()
			err := d.UnmarshalText(tt.text)
			if tt.want == "error" {
				assert.Error(t, err, "Unable to parse")
				assert.True(t, d.IsZero(), "date is not zero")
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
				assert.Equal(t, tt.want, d.String())
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{MustParse("2024-06-05"), `"2024-06-05"`},
		{ZeroDate(), `"0001-01-01"`},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Date{"%s"}.MarshalJSON()`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			json, err := tt.date.MarshalJSON()
			assert.Nil(t, err, "Excepted no error, got %v", err)
			assert.Equal(t, tt.want, string(json))
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		json []byte
		want string
	}{
		{[]byte(`"2024-06-05"`), "2024-06-05"},
		{[]byte{}, "error"},
		{[]byte("invalid"), "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("Date{}.UnmarshalJSON(`%s`)", tt.json)

		t.Run(testcase, func(t *testing.T) {
			d := ZeroDate()
			err := d.UnmarshalJSON(tt.json)
			if tt.want == "error" {
				assert.Error(t, err, "Unable to parse")
				assert.True(t, d.IsZero(), "date is not zero")
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
				assert.Equal(t, tt.want, d.String())
			}
		})
	}
}
