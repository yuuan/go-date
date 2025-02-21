package date

import (
	"fmt"
	"sort"
)

var (
	ErrDateRangesAreEmpty = fmt.Errorf("This DataRanges are empty")
)

type DateRanges []DateRange

// AreUnique checks if all DateRange instances in the DateRanges slice are unique.
func (drs DateRanges) AreUnique() bool {
	length := len(drs)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if drs[i].Equal(drs[j]) {
				return false
			}
		}
	}

	return true
}

// AreOverlapping checks if any DateRange instances in the DateRanges slice overlap with each other.
func (drs DateRanges) AreOverlapping() bool {
	length := len(drs)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if drs[i].OverlapsWith(drs[j]) {
				return true
			}
		}
	}

	return false
}

// SortMutable sorts the DateRanges slice in place in ascending order.
func (drs DateRanges) SortMutable() DateRanges {
	sort.SliceStable(drs, func(i, j int) bool {
		return drs[i].LessThan(drs[j])
	})

	return drs
}

// SortReverseMutable sorts the DateRanges slice in place in descending order.
func (drs DateRanges) SortReverseMutable() DateRanges {
	sort.SliceStable(drs, func(i, j int) bool {
		return drs[i].GreaterThanOrEqual(drs[j])
	})

	return drs
}

// Sort returns a new sorted DateRanges slice in ascending order.
func (drs DateRanges) Sort() DateRanges {
	return drs.clone().SortMutable()
}

// SortReverse returns a new sorted DateRanges slice in descending order.
func (drs DateRanges) SortReverse() DateRanges {
	return drs.clone().SortReverseMutable()
}

// StartDates returns a Dates slice containing the start dates of all DateRange instances in the DateRanges slice.
func (drs DateRanges) StartDates() Dates {
	starts := make(Dates, len(drs))

	for i, d := range drs {
		starts[i] = d.start
	}

	return starts
}

// EndDates returns a Dates slice containing the end dates of all DateRange instances in the DateRanges slice.
func (drs DateRanges) EndDates() Dates {
	ends := make(Dates, len(drs))

	for i, d := range drs {
		ends[i] = d.end
	}

	return ends
}

// FirstStart returns the earliest start date among all DateRange instances in the DateRanges slice.
func (drs DateRanges) FirstStart() (Date, error) {
	if len(drs) == 0 {
		return ZeroDate(), ErrDateRangesAreEmpty
	}

	return drs.StartDates().Sort().MustMin(), nil
}

// LastEnd returns the latest end date among all DateRange instances in the DateRanges slice.
func (drs DateRanges) LastEnd() (Date, error) {
	if len(drs) == 0 {
		return ZeroDate(), ErrDateRangesAreEmpty
	}

	return drs.EndDates().SortReverse().MustMax(), nil
}

// Strings returns a slice of string representations of all DateRange instances in the DateRanges slice.
func (drs DateRanges) Strings() []string {
	ranges := make([]string, len(drs))

	for i, dr := range drs {
		ranges[i] = dr.String()
	}

	return ranges
}

// clone creates a copy of the DateRanges slice.
func (drs DateRanges) clone() DateRanges {
	ranges := make(DateRanges, len(drs))
	copy(ranges, drs[:])

	return ranges
}
