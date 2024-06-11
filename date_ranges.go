package date

import (
	"fmt"
	"sort"
)

var (
	ErrDateRangesAreEmpty = fmt.Errorf("This DataRanges are empty")
)

type DateRanges []DateRange

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

func (drs DateRanges) SortMutable() DateRanges {
	sort.SliceStable(drs, func(i, j int) bool {
		return drs[i].LessThan(drs[j])
	})

	return drs
}

func (drs DateRanges) SortReverseMutable() DateRanges {
	sort.SliceStable(drs, func(i, j int) bool {
		return drs[i].GreaterThanOrEqual(drs[j])
	})

	return drs
}

func (drs DateRanges) Sort() DateRanges {
	return drs.clone().SortMutable()
}

func (drs DateRanges) SortReverse() DateRanges {
	return drs.clone().SortReverseMutable()
}

func (drs DateRanges) StartDates() Dates {
	starts := make(Dates, len(drs))

	for i, d := range drs {
		starts[i] = d.start
	}

	return starts
}

func (drs DateRanges) EndDates() Dates {
	ends := make(Dates, len(drs))

	for i, d := range drs {
		ends[i] = d.end
	}

	return ends
}

func (drs DateRanges) FirstStart() (Date, error) {
	if len(drs) == 0 {
		return ZeroDate(), ErrDateRangesAreEmpty
	}

	return drs.StartDates().Sort().MustMin(), nil
}

func (drs DateRanges) LastEnd() (Date, error) {
	if len(drs) == 0 {
		return ZeroDate(), ErrDateRangesAreEmpty
	}

	return drs.EndDates().SortReverse().MustMax(), nil
}

func (drs DateRanges) Strings() []string {
	ranges := make([]string, len(drs))

	for i, dr := range drs {
		ranges[i] = dr.String()
	}

	return ranges
}

func (drs DateRanges) clone() DateRanges {
	ranges := make(DateRanges, len(drs))
	copy(ranges, drs[:])

	return ranges
}
