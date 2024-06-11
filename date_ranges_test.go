package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateRangesSortMutable(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		ranges := DateRanges{
			MustParseDateRange("2024-06-02", "2024-06-02"),
			MustParseDateRange("2024-06-01", "2024-06-01"),
			MustParseDateRange("2024-06-03", "2024-06-04"),
			MustParseDateRange("2024-06-03", "2024-06-03"),
		}

		sorted := ranges.SortMutable()

		assert.Equal(t, "2024-06-01:2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02:2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03:2024-06-03", sorted[2].String())
		assert.Equal(t, "2024-06-03:2024-06-04", sorted[3].String())
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

		assert.Equal(t, "2024-06-03:2024-06-04", sorted[0].String())
		assert.Equal(t, "2024-06-03:2024-06-03", sorted[1].String())
		assert.Equal(t, "2024-06-02:2024-06-02", sorted[2].String())
		assert.Equal(t, "2024-06-01:2024-06-01", sorted[3].String())
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

		assert.Equal(t, "2024-06-01:2024-06-01", sorted[0].String())
		assert.Equal(t, "2024-06-02:2024-06-02", sorted[1].String())
		assert.Equal(t, "2024-06-03:2024-06-03", sorted[2].String())
		assert.Equal(t, "2024-06-03:2024-06-04", sorted[3].String())
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

		assert.Equal(t, "2024-06-03:2024-06-04", sorted[0].String())
		assert.Equal(t, "2024-06-03:2024-06-03", sorted[1].String())
		assert.Equal(t, "2024-06-02:2024-06-02", sorted[2].String())
		assert.Equal(t, "2024-06-01:2024-06-01", sorted[3].String())
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
