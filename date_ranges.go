package date

import "sort"

type DateRanges []DateRange

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

func (drs DateRanges) clone() DateRanges {
	ranges := make(DateRanges, len(drs))
	copy(ranges, drs[:])

	return ranges
}
