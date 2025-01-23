package date

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Factory functions
// --------------------------------------------------

func TestNullDateFromDate(t *testing.T) {
	tests := []struct {
		date Date
		want NullDate
	}{
		{
			MustParse("2024-06-05"),
			NullDate{date: MustParse("2024-06-05"), isNotNull: true},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDateFromDate(Date{"%s"})`, tt.date)

		t.Run(testcase, func(t *testing.T) {
			nd := NullDateFromDate(tt.date)

			assert.Equal(t, nd, tt.want)
		})
	}
}

func TestNullDateFromDatePtr(t *testing.T) {
	tests := []struct {
		ptr  *Date
		want NullDate
	}{
		{
			func() *Date { d := MustParse("2024-06-05"); return &d }(),
			NullDateFromDate(MustParse("2024-06-05")),
		},
		{
			nil,
			NullDate{},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDateFromDatePtr(%s)`, func(date *Date) string {
			if date == nil {
				return "nil"
			}
			return fmt.Sprintf(`&Date{"%s"}`, date.String())
		}(tt.ptr))

		t.Run(testcase, func(t *testing.T) {
			nd := NullDateFromDatePtr(tt.ptr)

			assert.Equal(t, nd, tt.want)
		})
	}
}

func TestNullDateForNull(t *testing.T) {
	t.Run("NullDateForNull()", func(t *testing.T) {
		assert.Equal(t, NullDate{}, NullDateForNull())
	})
}

// Determination methods
// --------------------------------------------------

func TestNullDateIsNull(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want bool
	}{
		{NullDateFromDate(MustParse("2024-06-05")), false},
		{NullDate{}, true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.IsNull()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.nd.IsNull(), tt.want)
		})
	}
}

func TestNullDateIsNotNull(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want bool
	}{
		{NullDateFromDate(MustParse("2024-06-05")), true},
		{NullDate{}, false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.IsNotNull()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.nd.IsNotNull(), tt.want)
		})
	}
}

// Comparison methods
// --------------------------------------------------

func TestNullDateEqual(t *testing.T) {
	tests := []struct {
		ndA  NullDate
		ndB  NullDate
		want bool
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDateFromDate(MustParse("2024-06-05")),
			true,
		},
		{
			NullDate{},
			NullDate{},
			true,
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDateFromDate(MustParse("2024-06-06")),
			false,
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDate{},
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NullDate{date:"%s",isNotNull:%v}.Equal(NullDate{date:"%s",isNotNull:%v})`,
			tt.ndA.date,
			tt.ndA.isNotNull,
			tt.ndB.date,
			tt.ndB.isNotNull,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.ndA.Equal(tt.ndB), tt.want)
		})
	}
}

func TestNullDateNotEqual(t *testing.T) {
	tests := []struct {
		ndA  NullDate
		ndB  NullDate
		want bool
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDateFromDate(MustParse("2024-06-05")),
			false,
		},
		{
			NullDate{},
			NullDate{},
			false,
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDateFromDate(MustParse("2024-06-06")),
			true,
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			NullDate{},
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NullDate{date:"%s",isNotNull:%v}.NotEqual(NullDate{date:"%s",isNotNull:%v})`,
			tt.ndA.date,
			tt.ndA.isNotNull,
			tt.ndB.date,
			tt.ndB.isNotNull,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.ndA.NotEqual(tt.ndB), tt.want)
		})
	}
}

// Conversion methods
// --------------------------------------------------

func TestNullDatePtr(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want *Date
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			func() *Date {
				d := MustParse("2024-06-05")
				return &d
			}(),
		},
		{
			NullDateFromDate(ZeroDate()),
			func() *Date {
				d := ZeroDate()
				return &d
			}(),
		},
		{
			NullDate{},
			nil,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.Ptr()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.nd.Ptr())
		})
	}

}

func TestTake(t *testing.T) {
	tests := []struct {
		nd      NullDate
		want    Date
		wantErr bool
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			MustParse("2024-06-05"),
			false,
		},
		{
			NullDateFromDate(ZeroDate()),
			ZeroDate(),
			false,
		},
		{
			NullDate{},
			ZeroDate(),
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.Take()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			date, err := tt.nd.Take()

			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("Want err is %v", tt.wantErr))
			assert.Equal(t, date, tt.want)
		})
	}
}

func TestTakeOr(t *testing.T) {
	tests := []struct {
		nd       NullDate
		fallback Date
		want     Date
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			MustParse("2024-01-01"),
			MustParse("2024-06-05"),
		},
		{
			NullDateFromDate(ZeroDate()),
			MustParse("2024-01-01"),
			ZeroDate(),
		},
		{
			NullDate{},
			MustParse("2024-01-01"),
			MustParse("2024-01-01"),
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NullDate{date:"%s",isNotNull:%v}.Take(Date{"%s"})`,
			tt.nd.date,
			tt.nd.isNotNull,
			tt.fallback,
		)

		t.Run(testcase, func(t *testing.T) {
			date := tt.nd.TakeOr(tt.fallback)

			assert.Equal(t, date, tt.want)
		})
	}
}

func TestNullDateString(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want string
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			"2024-06-05",
		},
		{
			NullDate{},
			"null",
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NullDate{date:"%s",isNotNull:%v}.String()`,
			tt.nd.date,
			tt.nd.isNotNull,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.nd.String())
		})
	}
}

func TestNullDateStringPtr(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want *string
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			func() *string {
				d := "2024-06-05"
				return &d
			}(),
		},
		{
			NullDate{},
			nil,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`NullDate{date:"%s",isNotNull:%v}.StringPtr()`,
			tt.nd.date,
			tt.nd.isNotNull,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.nd.StringPtr())
		})
	}
}

// Marshalling methods
// --------------------------------------------------

func TestNullDateValue(t *testing.T) {
	tests := []struct {
		nd      NullDate
		want    driver.Value
		wantErr bool
	}{
		{
			NullDateFromDate(MustParse("2024-06-05")),
			"2024-06-05",
			false,
		},
		{
			NullDate{},
			nil,
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.Value()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			value, err := tt.nd.Value()
			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("Want err is %v", tt.wantErr))
			assert.Equal(t, value, tt.want)
		})
	}
}

func TestNullDateScan(t *testing.T) {
	tests := []struct {
		value   interface{}
		want    NullDate
		wantErr bool
	}{
		{
			"2024-06-05",
			NullDateFromDate(MustParse("2024-06-05")),
			false,
		},
		{
			"invalid",
			NullDate{},
			true,
		},
		{
			nil,
			NullDate{},
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{}.Scan(%s)`, func(value interface{}) string {
			if value == nil {
				return "nil"
			}
			if v, ok := value.(string); ok {
				return fmt.Sprintf(`"%s"`, v)
			}
			return fmt.Sprintf(`"%v"`, value)
		}(tt.value))

		t.Run(testcase, func(t *testing.T) {
			nd := NullDate{}
			err := nd.Scan(tt.value)

			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("Want err is %v", tt.wantErr))
			assert.Equal(t, nd, tt.want)
		})
	}
}

func TestNullDateMarshalText(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want string
	}{
		{
			NullDate{},
			"null",
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			"2024-06-05",
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.MarshalText()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			text, err := tt.nd.MarshalText()

			assert.Nil(t, err, "Expected no error, got %v", err)
			assert.Equal(t, string(text), tt.want)
		})
	}
}

func TestNullDateUnmarshalText(t *testing.T) {
	tests := []struct {
		text    string
		want    NullDate
		wantErr bool
	}{
		{
			"2024-06-05",
			NullDateFromDate(MustParse("2024-06-05")),
			false,
		},
		{
			"null",
			NullDate{},
			false,
		},
		{
			"",
			NullDate{},
			true,
		},
		{
			"invalid",
			NullDate{},
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{}.UnmarshalText("%s")`, tt.text)

		t.Run(testcase, func(t *testing.T) {
			nd := NullDate{}
			err := nd.UnmarshalText([]byte(tt.text))

			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("Want err is %v", tt.wantErr))
			assert.Equal(t, nd, tt.want)
		})
	}
}

func TestNullDateMarshalJSON(t *testing.T) {
	tests := []struct {
		nd   NullDate
		want string
	}{
		{
			NullDate{},
			"null",
		},
		{
			NullDateFromDate(MustParse("2024-06-05")),
			`"2024-06-05"`,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`NullDate{date:"%s",isNotNull:%v}.MarshalJSON()`, tt.nd.date, tt.nd.isNotNull)

		t.Run(testcase, func(t *testing.T) {
			text, err := tt.nd.MarshalJSON()

			assert.Nil(t, err, "Expected no error, got %v", err)
			assert.Equal(t, string(text), tt.want)
		})
	}
}

func TestNullDateUnmarshalJSON(t *testing.T) {
	tests := []struct {
		json    string
		want    NullDate
		wantErr bool
	}{
		{
			`"2024-06-05"`,
			NullDateFromDate(MustParse("2024-06-05")),
			false,
		},
		{
			"null",
			NullDate{},
			false,
		},
		{
			"",
			NullDate{},
			true,
		},
		{
			"invalid",
			NullDate{},
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("NullDate{}.UnmarshalJSON(`%s`)", tt.json)

		t.Run(testcase, func(t *testing.T) {
			nd := NullDate{}
			err := nd.UnmarshalJSON([]byte(tt.json))

			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("Want err is %v", tt.wantErr))
			assert.Equal(t, nd, tt.want)
		})
	}
}
