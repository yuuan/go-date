package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Addition and Subtraction methods
// --------------------------------------------------

func TestNewMonth(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		want  string
	}{
		{2024, time.June, "2024-06"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("NewMonth(%d, %d)", tt.year, int(tt.month))

		t.Run(testcase, func(t *testing.T) {
			date := NewMonth(tt.year, tt.month)

			assert.Equal(t, tt.want, date.String())
		})
	}
}

func TestZeroMonth(t *testing.T) {
	t.Run("ZeroMonth()", func(t *testing.T) {
		month := ZeroMonth()
		assert.True(t, month.IsZero(), "Month created was not zero value")
	})
}

func TestMonthFromDate(t *testing.T) {
	tests := []struct {
		date Date
		want string
	}{
		{FromTime(time.Date(2024, time.June, 5, 0, 0, 0, 0, time.UTC)), "2024-06"},
		{FromTime(time.Date(2024, time.June, 5, 12, 30, 10, 10, time.Local)), "2024-06"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`FromTime(Date{date:"%s",location:"%s"})`, tt.date, tt.date.Location())

		t.Run(testcase, func(t *testing.T) {
			month := MonthFromDate(tt.date)
			assert.Equal(t, tt.want, month.String(), "Expected %s, got %s", tt.want, month)
		})
	}
}

func TestParseMonth(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"2024-06", "2024-06"},
		{"2024-06-05", "error"},
		{"2024/02", "error"},
		{"invalid-month", "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`ParseMonth("%s")`, tt.value)

		t.Run(testcase, func(t *testing.T) {
			month, err := ParseMonth(tt.value)
			if tt.want == "error" {
				assert.ErrorContains(t, err, "Unable to parse")
			} else {
				assert.Nil(t, err, "An unexpected error was returned")
				assert.Equal(t, tt.want, month.String(), "The parsed month does not match the expected month")
			}
		})
	}
}

func TestMustParseMonth(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"2024-06", "2024-06"},
		{"2024-06-05", "panic"},
		{"2024/02", "panic"},
		{"invalid-month", "panic"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`MustParseMonth("%s")`, tt.value)

		t.Run(testcase, func(t *testing.T) {
			if tt.want == "panic" {
				assert.Panics(t, func() { MustParseMonth(tt.value) }, "Did not panic")
			} else {
				month := MustParseMonth(tt.value)
				assert.Equal(t, tt.want, month.String(), "The parsed month does not match the expected month")
			}
		})
	}
}

func TestCurrentMonth(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local), "2024-06"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-06"},
		{time.Date(2024, time.June, 30, 23, 59, 59, 999, time.Local), "2024-06"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`CurrentMonth() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, CurrentMonth().String())
		})
	}
}

func TestNextMonth(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local), "2024-07"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-07"},
		{time.Date(2024, time.June, 30, 23, 59, 59, 999, time.Local), "2024-07"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NextMonth() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, NextMonth().String())
		})
	}
}

func TestLastMonth(t *testing.T) {
	tests := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local), "2024-05"},
		{time.Date(2024, time.June, 5, 12, 34, 56, 789, time.Local), "2024-05"},
		{time.Date(2024, time.June, 30, 23, 59, 59, 999, time.Local), "2024-05"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`LastMonth() at %s`, tt.now.Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			SetTestNow(func() time.Time { return tt.now })
			defer ResetTestNow()

			assert.Equal(t, tt.want, LastMonth().String())
		})
	}
}

// Determination methods
// --------------------------------------------------

func TestMonthIsZero(t *testing.T) {
	tests := []struct {
		testcase string
		month    func() Month
		want     bool
	}{
		{"ZeroMonth().IsZero()", ZeroMonth, true},
		{"Month{}.IsZero()", func() Month { return Month{} }, true},
		{"CurrentMonth().IsZero()", CurrentMonth, false},
	}

	for _, tt := range tests {
		t.Run(tt.testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month().IsZero())
		})
	}
}

func TestMonthIsJanuary(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), true},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsJanuary()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsJanuary())
		})
	}
}

func TestMonthIsFebruary(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), true},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsFebruary()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsFebruary())
		})
	}
}

func TestMonthIsMarch(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), true},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsMarch()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsMarch())
		})
	}
}

func TestMonthIsApril(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), true},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsApril()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsApril())
		})
	}
}

func TestMonthIsMay(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), true},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsMay()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsMay())
		})
	}
}

func TestMonthIsJune(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsJune()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsJune())
		})
	}
}

func TestMonthIsJuly(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), true},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsJuly()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsJuly())
		})
	}
}

func TestMonthIsAugust(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), true},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsAugust()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsAugust())
		})
	}
}

func TestMonthIsSeptember(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), true},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsSeptember()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsSeptember())
		})
	}
}

func TestMonthIsOctober(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), true},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsOctober()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsOctober())
		})
	}
}

func TestMonthIsNovember(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), true},
		{MustParseMonth("2024-12"), false},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsNovember()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsNovember())
		})
	}
}

func TestMonthIsDecember(t *testing.T) {
	tests := []struct {
		month Month
		want  bool
	}{
		{MustParseMonth("2024-01"), false},
		{MustParseMonth("2024-02"), false},
		{MustParseMonth("2024-03"), false},
		{MustParseMonth("2024-04"), false},
		{MustParseMonth("2024-05"), false},
		{MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), false},
		{MustParseMonth("2024-08"), false},
		{MustParseMonth("2024-09"), false},
		{MustParseMonth("2024-10"), false},
		{MustParseMonth("2024-11"), false},
		{MustParseMonth("2024-12"), true},
		{ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.IsDecember()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsDecember())
		})
	}
}

func TestMonthIsPast(t *testing.T) {
	tests := []struct {
		now   time.Time
		month Month
		want  bool
	}{
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-05"),
			true,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-05"),
			true,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		testcase := fmt.Sprintf(`Month{"%s"}.IsPast() at %s`, tt.month, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsPast())
		})
	}
}

func TestMonthIsFuture(t *testing.T) {
	tests := []struct {
		now   time.Time
		month Month
		want  bool
	}{
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-07"),
			true,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-07"),
			true,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		testcase := fmt.Sprintf(`Month{"%s"}.IsFuture() at %s`, tt.month, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsFuture())
		})
	}
}

func TestMonthIsCurrentMonth(t *testing.T) {
	tests := []struct {
		now   time.Time
		month Month
		want  bool
	}{
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-06"),
			true,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-06"),
			true,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		testcase := fmt.Sprintf(`Month{"%s"}.IsCurrentMonth() at %s`, tt.month, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsCurrentMonth())
		})
	}
}

func TestMonthIsNextMonth(t *testing.T) {
	tests := []struct {
		now   time.Time
		month Month
		want  bool
	}{
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-05"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-07"),
			true,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-07"),
			true,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-08"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-08"),
			false,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		testcase := fmt.Sprintf(`Month{"%s"}.IsNextMonth() at %s`, tt.month, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsNextMonth())
		})
	}
}

func TestMonthIsLastMonth(t *testing.T) {
	tests := []struct {
		now   time.Time
		month Month
		want  bool
	}{
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-06"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-04"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-04"),
			false,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-05"),
			true,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-05"),
			true,
		},
		{
			time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
		{
			time.Date(2024, time.June, 30, 23, 59, 59, 999999999, time.Local),
			MustParseMonth("2024-07"),
			false,
		},
	}

	for _, tt := range tests {
		SetTestNow(func() time.Time { return tt.now })
		defer ResetTestNow()

		testcase := fmt.Sprintf(`Month{"%s"}.IsLastMonth() at %s`, tt.month, now().Format(iso8601))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.IsLastMonth())
		})
	}
}

// Comparison methods
// --------------------------------------------------

func TestMonthCompare(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   int
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), -1},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), 0},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), 1},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Compare(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.Compare(tt.monthB))
		})
	}
}

func TestMonthEqual(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), false},
		{ZeroMonth(), ZeroMonth(), true},
		{ZeroMonth(), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-06"), ZeroMonth(), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Equal(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.Equal(tt.monthB))
		})
	}
}

func TestMonthNotEqual(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), true},
		{ZeroMonth(), ZeroMonth(), false},
		{ZeroMonth(), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-06"), ZeroMonth(), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Equal(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.NotEqual(tt.monthB))
		})
	}
}

func TestMonthAfter(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.After(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.After(tt.monthB))
		})
	}
}

func TestMonthAfterOrEqual(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AfterOrEqual(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.AfterOrEqual(tt.monthB))
		})
	}
}

func TestMonthBefore(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), false},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Before(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.Before(tt.monthB))
		})
	}
}

func TestMonthBeforeOrEqual(t *testing.T) {
	tests := []struct {
		monthA Month
		monthB Month
		want   bool
	}{
		{MustParseMonth("2024-05"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-06"), MustParseMonth("2024-06"), true},
		{MustParseMonth("2024-07"), MustParseMonth("2024-06"), false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.BeforeOrEqual(Month{"%s"})`, tt.monthA, tt.monthB)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.monthA.BeforeOrEqual(tt.monthB))
		})
	}
}

func TestMonthBetween(t *testing.T) {
	tests := []struct {
		month   Month
		start   Month
		end     Month
		want    bool
		wantErr error
	}{
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-04"),
			MustParseMonth("2024-06"),
			true,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-04"),
			MustParseMonth("2024-05"),
			true,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-05"),
			MustParseMonth("2024-06"),
			true,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-05"),
			MustParseMonth("2024-05"),
			true,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-03"),
			MustParseMonth("2024-04"),
			false,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-06"),
			MustParseMonth("2024-07"),
			false,
			nil,
		},
		{
			MustParseMonth("2024-05"),
			MustParseMonth("2024-06"),
			MustParseMonth("2024-04"),
			false,
			ErrEndMonthIsBeforeStartMonth,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`Month{"%s"}.Between(Month{"%s"},Month{"%s"})`,
			tt.month,
			tt.start,
			tt.end,
		)

		t.Run(testcase, func(t *testing.T) {
			b, err := tt.month.Between(tt.start, tt.end)
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

func TestMonthAddMonth(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-06"), "2024-07"},
		{MustParseMonth("2024-12"), "2025-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AddMonth()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.AddMonth().String())
		})
	}
}

func TestMonthAddMonths(t *testing.T) {
	tests := []struct {
		month  Month
		months int
		want   Month
	}{
		{
			MustParseMonth("2024-01"),
			1,
			MustParseMonth("2024-02"),
		},
		{
			MustParseMonth("2024-02"),
			2,
			MustParseMonth("2024-04"),
		},
		{
			MustParseMonth("2024-10"),
			2,
			MustParseMonth("2024-12"),
		},
		{
			MustParseMonth("2024-11"),
			2,
			MustParseMonth("2025-01"),
		},
		{
			MustParseMonth("2024-12"),
			2,
			MustParseMonth("2025-02"),
		},
		{
			MustParseMonth("2024-01"),
			24,
			MustParseMonth("2026-01"),
		},
		{
			MustParseMonth("2024-02"),
			-1,
			MustParseMonth("2024-01"),
		},
		{
			MustParseMonth("2024-04"),
			-2,
			MustParseMonth("2024-02"),
		},
		{
			MustParseMonth("2024-03"),
			-2,
			MustParseMonth("2024-01"),
		},
		{
			MustParseMonth("2024-02"),
			-2,
			MustParseMonth("2023-12"),
		},
		{
			MustParseMonth("2024-01"),
			-2,
			MustParseMonth("2023-11"),
		},
		{
			MustParseMonth("2024-01"),
			-24,
			MustParseMonth("2022-01"),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AddMonths(%d)`, tt.month.String(), tt.months)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.AddMonths(tt.months))
		})
	}
}

func TestMonthSubMonth(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-06"), "2024-05"},
		{MustParseMonth("2024-01"), "2023-12"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.SubMonth()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.SubMonth().String())
		})
	}
}

func TestMonthSubMonths(t *testing.T) {
	tests := []struct {
		month  Month
		months int
		want   Month
	}{
		{
			MustParseMonth("2024-12"),
			1,
			MustParseMonth("2024-11"),
		},
		{
			MustParseMonth("2024-08"),
			2,
			MustParseMonth("2024-06"),
		},
		{
			MustParseMonth("2024-03"),
			2,
			MustParseMonth("2024-01"),
		},
		{
			MustParseMonth("2024-02"),
			2,
			MustParseMonth("2023-12"),
		},
		{
			MustParseMonth("2024-01"),
			2,
			MustParseMonth("2023-11"),
		},
		{
			MustParseMonth("2024-01"),
			24,
			MustParseMonth("2022-01"),
		},
		{
			MustParseMonth("2024-11"),
			-1,
			MustParseMonth("2024-12"),
		},
		{
			MustParseMonth("2024-06"),
			-2,
			MustParseMonth("2024-08"),
		},
		{
			MustParseMonth("2024-10"),
			-2,
			MustParseMonth("2024-12"),
		},
		{
			MustParseMonth("2024-11"),
			-2,
			MustParseMonth("2025-01"),
		},
		{
			MustParseMonth("2024-12"),
			-2,
			MustParseMonth("2025-02"),
		},
		{
			MustParseMonth("2024-01"),
			-24,
			MustParseMonth("2026-01"),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AddMonths(%d)`, tt.month.String(), tt.months)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.SubMonths(tt.months))
		})
	}
}

func TestMonthAddYear(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-01"), "2025-01"},
		{MustParseMonth("2024-12"), "2025-12"},
		{MustParseMonth("0000-01"), "0001-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AddYear()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.AddYear().String())
		})
	}
}

func TestMonthAddYears(t *testing.T) {
	tests := []struct {
		month Month
		years int
		want  string
	}{
		{MustParseMonth("2024-01"), 0, "2024-01"},
		{MustParseMonth("2024-01"), 1, "2025-01"},
		{MustParseMonth("2024-01"), -1, "2023-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.AddYears(%d)`, tt.month, tt.years)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.AddYears(tt.years).String())
		})
	}
}

func TestMonthSubYear(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-01"), "2023-01"},
		{MustParseMonth("2024-12"), "2023-12"},
		{MustParseMonth("0001-01"), "0000-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.SubYear()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.SubYear().String())
		})
	}
}

func TestMonthSubYears(t *testing.T) {
	tests := []struct {
		month Month
		years int
		want  string
	}{
		{MustParseMonth("2024-01"), 0, "2024-01"},
		{MustParseMonth("2024-01"), 1, "2023-01"},
		{MustParseMonth("2024-01"), -1, "2025-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.SubYears(%d)`, tt.month, tt.years)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.SubYears(tt.years).String())
		})
	}
}

// Conversion methods
// --------------------------------------------------

func TestMonthYear(t *testing.T) {
	tests := []struct {
		month Month
		want  int
	}{
		{MustParseMonth("2024-06"), 2024},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Year()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.Year())
		})
	}
}

func TestMonthMonth(t *testing.T) {
	tests := []struct {
		month Month
		want  time.Month
	}{
		{MustParseMonth("2024-06"), time.June},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Month()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.Month())
		})
	}
}

func TestMonthFirstDate(t *testing.T) {
	tests := []struct {
		month Month
		want  Date
	}{
		{MustParseMonth("2024-06"), MustParse("2024-06-01")},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.FirstDate()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.FirstDate())
		})
	}
}

func TestMonthLastDate(t *testing.T) {
	tests := []struct {
		month Month
		want  Date
	}{
		{MustParseMonth("2024-01"), MustParse("2024-01-31")},
		{MustParseMonth("2024-02"), MustParse("2024-02-29")},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.LastDate()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.LastDate())
		})
	}
}

func TestMonthToDateRange(t *testing.T) {
	tests := []struct {
		month Month
		want  DateRange
	}{
		{MustParseMonth("2024-01"), MustParseDateRange("2024-01-01", "2024-01-31")},
		{MustParseMonth("2024-02"), MustParseDateRange("2024-02-01", "2024-02-29")},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.ToDateRange()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.ToDateRange())
		})
	}
}

func TestMonthDays(t *testing.T) {
	tests := []struct {
		month Month
		want  int
	}{
		{MustParseMonth("2024-01"), 31},
		{MustParseMonth("2024-02"), 29},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Days()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.Days())
		})
	}
}

func TestMonthDates(t *testing.T) {
	tests := []struct {
		month Month
		want  Dates
	}{
		{MustParseMonth("2024-01"), MustParseDateRange("2024-01-01", "2024-01-31").Dates()},
		{MustParseMonth("2024-02"), MustParseDateRange("2024-02-01", "2024-02-29").Dates()},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Dates()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.Dates())
		})
	}
}

func TestMonthFormat(t *testing.T) {
	tests := []struct {
		month  Month
		format string
		want   string
	}{
		{MustParseMonth("2024-06"), "2006-01", "2024-06"},
		{MustParseMonth("2024-06"), "2006/1", "2024/6"},
		{MustParseMonth("2024-06"), "06 Jan", "24 Jun"},
		{MustParseMonth("2024-06"), "2006 January", "2024 June"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Format("%s")`, tt.month, tt.format)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.Format(tt.format))
		})
	}
}

func TestMonthSplit(t *testing.T) {
	tests := []struct {
		month     Month
		wantYear  int
		wantMonth time.Month
	}{
		{MustParseMonth("2024-06"), 2024, time.June},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.Split()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			year, month := tt.month.Split()
			assert.Equal(t, tt.wantYear, year)
			assert.Equal(t, tt.wantMonth, month)
		})
	}
}

func TestMonthString(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{NewMonth(2024, time.June), "2024-06"},
		{NewMonth(0, time.January), "0000-01"},
		{NewMonth(10000, time.January), "10000-01"},
		{NewMonth(-1, time.January), "-0001-01"},
		{NewMonth(-10000, time.January), "-10000-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.String()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.month.String())
		})
	}
}

// Marshalling methods
// --------------------------------------------------

func TestMonthMarshalText(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-06"), "2024-06"},
		{ZeroMonth(), "0001-01"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.MarshalText()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			text, err := tt.month.MarshalText()
			assert.Nil(t, err, "Excepted no error, got %v", err)
			assert.Equal(t, tt.want, string(text))
		})
	}
}

func TestMonthUnmarshalText(t *testing.T) {
	tests := []struct {
		text []byte
		want string
	}{
		{[]byte("2024-06"), "2024-06"},
		{[]byte{}, "error"},
		{[]byte("invalid"), "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{}.UnmarshalText("%s")`, tt.text)

		t.Run(testcase, func(t *testing.T) {
			d := ZeroMonth()
			err := d.UnmarshalText(tt.text)
			if tt.want == "error" {
				assert.Error(t, err, "Unable to parse")
				assert.True(t, d.IsZero(), "month is not zero")
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
				assert.Equal(t, tt.want, d.String())
			}
		})
	}
}

func TestMonthMarshalJSON(t *testing.T) {
	tests := []struct {
		month Month
		want  string
	}{
		{MustParseMonth("2024-06"), `"2024-06"`},
		{ZeroMonth(), `"0001-01"`},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Month{"%s"}.MarshalJSON()`, tt.month)

		t.Run(testcase, func(t *testing.T) {
			json, err := tt.month.MarshalJSON()
			assert.Nil(t, err, "Excepted no error, got %v", err)
			assert.Equal(t, tt.want, string(json))
		})
	}
}

func TestMonthUnmarshalJSON(t *testing.T) {
	tests := []struct {
		json []byte
		want string
	}{
		{[]byte(`"2024-06"`), "2024-06"},
		{[]byte{}, "error"},
		{[]byte("invalid"), "error"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("Month{}.UnmarshalJSON(`%s`)", tt.json)

		t.Run(testcase, func(t *testing.T) {
			d := ZeroMonth()
			err := d.UnmarshalJSON(tt.json)
			if tt.want == "error" {
				assert.Error(t, err, "Unable to parse")
				assert.True(t, d.IsZero(), "month is not zero")
			} else {
				assert.Nil(t, err, "Expected no error, got %v", err)
				assert.Equal(t, tt.want, d.String())
			}
		})
	}
}
