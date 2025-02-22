package date

import (
	"fmt"
	"sort"
)

var (
	ErrDatesAreEmpty = fmt.Errorf("this Dates are empty")
)

type Dates []Date

// AreUnique checks if all Date instances in the Dates slice are unique.
func (ds Dates) AreUnique() bool {
	length := len(ds)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if ds[i].Equal(ds[j]) {
				return false
			}
		}
	}

	return true
}

// SortMutable sorts the Dates slice in place in ascending order.
func (ds Dates) SortMutable() Dates {
	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i].Before(ds[j])
	})

	return ds
}

// SortReverseMutable sorts the Dates slice in place in descending order.
func (ds Dates) SortReverseMutable() Dates {
	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i].After(ds[j])
	})

	return ds
}

// Sort returns a new sorted Dates slice in ascending order.
func (ds Dates) Sort() Dates {
	return ds.clone().SortMutable()
}

// SortReverse returns a new sorted Dates slice in descending order.
func (ds Dates) SortReverse() Dates {
	return ds.clone().SortReverseMutable()
}

// Min returns the minimum Date in the Dates slice.
func (ds Dates) Min() (Date, error) {
	if len(ds) == 0 {
		return ZeroDate(), fmt.Errorf("Min: %w", ErrDatesAreEmpty)
	}

	dates := ds.Sort()

	return dates[0], nil
}

// MustMin returns the minimum Date in the Dates slice. It panics if the Dates slice is empty.
func (ds Dates) MustMin() Date {
	min, err := ds.Min()
	if err != nil {
		panic(err)
	}

	return min
}

// Max returns the maximum Date in the Dates slice.
func (ds Dates) Max() (Date, error) {
	if len(ds) == 0 {
		return ZeroDate(), fmt.Errorf("Max: %w", ErrDatesAreEmpty)
	}

	dates := ds.SortReverse()

	return dates[0], nil
}

// MustMax returns the maximum Date in the Dates slice. It panics if the Dates slice is empty.
func (ds Dates) MustMax() Date {
	max, err := ds.Max()
	if err != nil {
		panic(err)
	}

	return max
}

// Equal checks if the Dates slice is equal to another Dates slice.
func (ds Dates) Equal(targets Dates) bool {
	if len(ds) != len(targets) {
		return false
	}

	sorted := ds.Sort()

	for i, d := range targets.Sort() {
		if d.NotEqual(sorted[i]) {
			return false
		}
	}

	return true
}

// Strings returns a slice of string representations of all Date instances in the Dates slice.
func (ds Dates) Strings() []string {
	dates := make([]string, len(ds))

	for i, d := range ds {
		dates[i] = d.String()
	}

	return dates
}

// clone creates a copy of the Dates slice.
func (ds Dates) clone() Dates {
	dates := make(Dates, len(ds))
	copy(dates, ds[:])

	return dates
}
