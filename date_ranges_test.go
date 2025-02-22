package date

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateRangesAreUnique(t *testing.T) {
	tests := []struct {
		ranges DateRanges
		want   bool
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
			},
			true,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-01", "2024-06-03"),
				MustParseDateRange("2024-06-02", "2024-06-03"),
			},
			true,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-01", "2024-06-01"),
			},
			false,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-02", "2024-06-03"),
			},
			false,
		},
		{
			DateRanges{
				ZeroDateRange(),
				ZeroDateRange(),
			},
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.AreUnique()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ranges.AreUnique())
		})
	}
}

func TestDateRangesAreOverlapping(t *testing.T) {
	tests := []struct {
		ranges DateRanges
		want   bool
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
			},
			false,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-02", "2024-06-03"),
				MustParseDateRange("2024-06-04", "2024-06-05"),
			},
			false,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-01", "2024-06-01"),
			},
			true,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-10"),
				MustParseDateRange("2024-06-05", "2024-06-05"),
			},
			true,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-05", "2024-06-05"),
				MustParseDateRange("2024-06-01", "2024-06-10"),
			},
			true,
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-05"),
				MustParseDateRange("2024-06-05", "2024-06-10"),
			},
			true,
		},
		{
			DateRanges{
				ZeroDateRange(),
				ZeroDateRange(),
			},
			true,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.AreOverlapping()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ranges.AreOverlapping())
		})
	}
}

func TestDateRangesSortMutable(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
		}

		sorted := ranges.SortMutable()

		assert.Equal(t, "2024-06-01/2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02/2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03/2024-06-03", sorted[2].String())
		assert.Equal(t, "2024-06-03/2024-06-04", sorted[3].String())
	})

	t.Run("mutable", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := ranges.SortMutable()

			sorted[0].start = MustParse("2024-01-01")

			assert.Equal(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := ranges.SortMutable()

			sorted[0] = MustParseDateRange("2024-01-01", "2024-01-01")

			assert.Equal(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})
	})

}

func TestDateRangesSortReverseMutable(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
		}

		sorted := ranges.SortReverseMutable()

		assert.Equal(t, "2024-06-03/2024-06-04", sorted[0].String())
		assert.Equal(t, "2024-06-03/2024-06-03", sorted[1].String())
		assert.Equal(t, "2024-06-02/2024-06-02", sorted[2].String())
		assert.Equal(t, "2024-06-01/2024-06-01", sorted[3].String())
	})

	t.Run("mutable", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := ranges.SortReverseMutable()

			sorted[0].start = MustParse("2024-01-01")

			assert.Equal(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := ranges.SortReverseMutable()

			sorted[0] = MustParseDateRange("2024-01-01", "2024-01-01")

			assert.Equal(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})
	})
}

func TestDateRangesSort(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
		}

		sorted := ranges.Sort()

		assert.Equal(t, "2024-06-01/2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02/2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03/2024-06-03", sorted[2].String())
		assert.Equal(t, "2024-06-03/2024-06-04", sorted[3].String())
	})

	t.Run("immutable", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := ranges.Sort()

			sorted[0].start = MustParse("2024-01-01")

			assert.NotEqual(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := ranges.Sort()

			sorted[0] = MustParseDateRange("2024-01-01", "2024-01-01")

			assert.NotEqual(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})
	})

}

func TestDateRangesSortReverse(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
		}

		sorted := ranges.SortReverse()

		assert.Equal(t, "2024-06-03/2024-06-04", sorted[0].String())
		assert.Equal(t, "2024-06-03/2024-06-03", sorted[1].String())
		assert.Equal(t, "2024-06-02/2024-06-02", sorted[2].String())
		assert.Equal(t, "2024-06-01/2024-06-01", sorted[3].String())
	})

	t.Run("immutable", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := ranges.SortReverse()

			sorted[0].start = MustParse("2024-01-01")

			assert.NotEqual(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := ranges.SortReverse()

			sorted[0] = MustParseDateRange("2024-01-01", "2024-01-01")

			assert.NotEqual(t, ranges[0], sorted[0])
			assert.Equal(t, ranges[1], sorted[1])
			assert.Equal(t, ranges[2], sorted[2])
			assert.Equal(t, ranges[3], sorted[3])
		})
	})
}

func TestDateRangesStartDates(t *testing.T) {
	tests := []struct {
		ranges DateRanges
		want   Dates
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-05", "2024-06-05"),
				MustParseDateRange("2024-06-05", "2024-06-06"),
				MustParseDateRange("2024-06-06", "2024-06-07"),
			},
			Dates{
				MustParse("2024-06-05"),
				MustParse("2024-06-05"),
				MustParse("2024-06-06"),
			},
		},
		{
			DateRanges{},
			Dates{},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.StartDates()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ranges.StartDates())
		})
	}
}

func TestDateRangesEndDates(t *testing.T) {
	tests := []struct {
		ranges DateRanges
		want   Dates
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-05", "2024-06-06"),
				MustParseDateRange("2024-06-06", "2024-06-06"),
				MustParseDateRange("2024-06-07", "2024-06-07"),
			},
			Dates{
				MustParse("2024-06-06"),
				MustParse("2024-06-06"),
				MustParse("2024-06-07"),
			},
		},
		{
			DateRanges{},
			Dates{},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.EndDates()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ranges.EndDates())
		})
	}
}

func TestDateRangesFirstStart(t *testing.T) {
	tests := []struct {
		ranges  DateRanges
		want    Date
		wantErr error
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-05", "2024-06-05"),
				MustParseDateRange("2024-06-05", "2024-06-06"),
				MustParseDateRange("2024-06-06", "2024-06-07"),
			},
			MustParse("2024-06-05"),
			nil,
		},
		{
			DateRanges{},
			ZeroDate(),
			ErrDateRangesAreEmpty,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.FirstStart()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			date, err := tt.ranges.FirstStart()

			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
				assert.Equal(t, tt.want, date)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, date)
			}
		})
	}
}

func TestDateRangesLastEnd(t *testing.T) {
	tests := []struct {
		ranges  DateRanges
		want    Date
		wantErr error
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-05", "2024-06-06"),
				MustParseDateRange("2024-06-05", "2024-06-06"),
				MustParseDateRange("2024-06-06", "2024-06-07"),
			},
			MustParse("2024-06-07"),
			nil,
		},
		{
			DateRanges{},
			ZeroDate(),
			ErrDateRangesAreEmpty,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{"%s"}.LastEnd()`, strings.Join(tt.ranges.Strings(), `","`))

		t.Run(testcase, func(t *testing.T) {
			date, err := tt.ranges.LastEnd()

			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
				assert.Equal(t, tt.want, date)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, date)
			}
		})
	}
}

func TestDateRangesStrings(t *testing.T) {
	tests := []struct {
		ranges DateRanges
		want   []string
	}{
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
			},
			[]string{
				"2024-06-01/2024-06-01",
			},
		},
		{
			DateRanges{
				MustParseDateRange("2024-06-01", "2024-06-01"),
				MustParseDateRange("2024-06-01", "2024-06-02"),
				MustParseDateRange("2024-06-02", "2024-06-03"),
			},
			[]string{
				"2024-06-01/2024-06-01",
				"2024-06-01/2024-06-02",
				"2024-06-02/2024-06-03",
			},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`DateRanges{%s}.Strings()`, `"`+strings.Join(tt.ranges.Strings(), `","`)+`"`)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ranges.Strings())
		})
	}
}
