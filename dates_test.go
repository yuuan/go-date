package date

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatesSortMutable(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		sorted := dates.SortMutable()

		assert.Equal(t, "2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03", sorted[2].String())
	})

	t.Run("mutable", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-01"),
			MustParse("2024-06-02"),
			MustParse("2024-06-03"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := dates.SortMutable()

			sorted[0].value = time.Now()

			assert.Equal(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := dates.SortMutable()

			sorted[0] = MustParse("2024-01-01")

			assert.Equal(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})
	})

}

func TestDatesSortReverseMutable(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		sorted := dates.SortReverseMutable()

		assert.Equal(t, "2024-06-03", sorted[0].String())
		assert.Equal(t, "2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-01", sorted[2].String())
	})

	t.Run("mutable", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-03"),
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := dates.SortReverseMutable()

			sorted[0].value = time.Now()

			assert.Equal(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := dates.SortReverseMutable()

			sorted[0] = MustParse("2024-01-01")

			assert.Equal(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})
	})
}

func TestDatesSort(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		sorted := dates.Sort()

		assert.Equal(t, "2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03", sorted[2].String())
	})

	t.Run("immutable", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-01"),
			MustParse("2024-06-02"),
			MustParse("2024-06-03"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := dates.Sort()

			sorted[0].value = time.Now()

			assert.NotEqual(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := dates.Sort()

			sorted[0] = MustParse("2024-01-01")

			assert.NotEqual(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})
	})

}

func TestDatesSortReverse(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		sorted := dates.SortReverse()

		assert.Equal(t, "2024-06-03", sorted[0].String())
		assert.Equal(t, "2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-01", sorted[2].String())
	})

	t.Run("immutable", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-03"),
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
		}

		t.Run("change field of slice element", func(t *testing.T) {
			sorted := dates.SortReverse()

			sorted[0].value = time.Now()

			assert.NotEqual(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})

		t.Run("change slice element", func(t *testing.T) {
			sorted := dates.SortReverse()

			sorted[0] = MustParse("2024-01-01")

			assert.NotEqual(t, dates[0], sorted[0])
			assert.Equal(t, dates[1], sorted[1])
			assert.Equal(t, dates[2], sorted[2])
		})
	})
}

func TestDatesMin(t *testing.T) {
	t.Run("When Dates is not empty", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		min, err := dates.Min()

		assert.Equal(t, "2024-06-01", min.String())
		assert.Nil(t, err, "Expected no error, got %v", err)
	})

	t.Run("When Dates is empty", func(t *testing.T) {
		dates := Dates{}

		min, err := dates.Min()

		assert.ErrorIs(t, ErrDatesAreEmpty, err)
		assert.Equal(t, ZeroDate(), min)
	})
}

func TestDatesMustMin(t *testing.T) {
	t.Run("When Dates is not empty", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		assert.Equal(t, "2024-06-01", dates.MustMin().String())
	})

	t.Run("When Dates is empty", func(t *testing.T) {
		dates := Dates{}

		assert.Panics(t, func() { dates.MustMin() }, "Did not panic")
	})
}

func TestDatesMax(t *testing.T) {
	t.Run("When Dates is not empty", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		max, err := dates.Max()

		assert.Equal(t, "2024-06-03", max.String())
		assert.Nil(t, err, "Expected no error, got %v", err)
	})

	t.Run("When Dates is empty", func(t *testing.T) {
		dates := Dates{}

		max, err := dates.Max()

		assert.ErrorIs(t, ErrDatesAreEmpty, err)
		assert.Equal(t, ZeroDate(), max)
	})
}

func TestDatesMustMax(t *testing.T) {
	t.Run("When Dates is not empty", func(t *testing.T) {
		dates := Dates{
			MustParse("2024-06-02"),
			MustParse("2024-06-01"),
			MustParse("2024-06-03"),
		}

		assert.Equal(t, "2024-06-03", dates.MustMax().String())
	})

	t.Run("When Dates is empty", func(t *testing.T) {
		dates := Dates{}

		assert.Panics(t, func() { dates.MustMax() }, "Did not panic")
	})
}

func TestDatesEqual(t *testing.T) {
	tests := []struct {
		dsA  Dates
		dsB  Dates
		want bool
	}{
		{
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			true,
		},
		{
			Dates{MustParse("2024-06-05"), MustParse("2024-06-05"), MustParse("2024-06-04")},
			Dates{MustParse("2024-06-05"), MustParse("2024-06-04"), MustParse("2024-06-05")},
			true,
		},
		{
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			true,
		},
		{
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			Dates{MustParse("2024-06-05")},
			false,
		},
		{
			Dates{MustParse("2024-06-05"), MustParse("2024-06-05")},
			Dates{MustParse("2024-06-05")},
			false,
		},
		{
			Dates{MustParse("2024-06-04"), MustParse("2024-06-05"), MustParse("2024-06-06")},
			Dates{MustParse("2023-06-04"), MustParse("2023-06-05"), MustParse("2023-06-06")},
			false,
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(
			`Dates{%s}.Equal(Dates{%s})`,
			`"`+strings.Join(tt.dsA.Strings(), `","`)+`"`,
			`"`+strings.Join(tt.dsB.Strings(), `","`)+`"`,
		)

		t.Run(testcase, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dsA.Equal(tt.dsB))
		})
	}
}

func TestDatesStrings(t *testing.T) {
	tests := []struct {
		ds []string
	}{
		{
			[]string{"2024-06-05"},
		},
		{
			[]string{"2024-06-04", "2024-06-05", "2024-06-06"},
		},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf(`Dates{%s}.Strings()`, `"`+strings.Join(tt.ds, `","`)+`"`)

		t.Run(testcase, func(t *testing.T) {
			dates := make(Dates, len(tt.ds))

			for i, d := range tt.ds {
				dates[i] = MustParse(d)
			}

			assert.Equal(t, tt.ds, dates.Strings())
		})
	}
}
