package date

import "sort"

type Dates []Date

func (ds Dates) SortMutable() Dates {
	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i].Before(ds[j])
	})

	return ds
}

func (ds Dates) SortReverseMutable() Dates {
	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i].After(ds[j])
	})

	return ds
}

func (ds Dates) Sort() Dates {
	return ds.clone().SortMutable()
}

func (ds Dates) SortReverse() Dates {
	return ds.clone().SortReverseMutable()
}

func (ds Dates) Min() Date {
	dates := ds.Sort()

	return dates[0]
}

func (ds Dates) Max() Date {
	dates := ds.SortReverse()

	return dates[0]
}

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

func (ds Dates) Strings() []string {
	dates := make([]string, len(ds))

	for i, d := range ds {
		dates[i] = d.String()
	}

	return dates
}

func (ds Dates) clone() Dates {
	dates := make(Dates, len(ds))
	copy(dates, ds[:])

	return dates
}
