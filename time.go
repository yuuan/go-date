package date

import (
	"time"
)

// StrictLocationComparison determines how time.Location values are compared.
//
// If true, Location comparison is strict and checks whether the pointer
// addresses of the Locations are identical. This ensures that only the exact
// same Location instances are considered equal.
//
// If false, Location comparison is relaxed and checks whether the names of
// the Locations are equal. This allows Locations with the same name but
// different underlying instances to be considered equal.
//
// Default: false
var StrictLocationComparison = false

const iso8601 = "2006-01-02T15:04:05.999999999+09:00"

var (
	now      = time.Now
	location = func() *time.Location { return time.Local }
)

func SetTestNow(getNow func() time.Time) {
	now = getNow
}

func ResetTestNow() {
	now = time.Now
}

func SetTestLocation(getLocation func() *time.Location) {
	location = getLocation
}

func ResetTestLocation() {
	location = func() *time.Location { return time.Local }
}

func today() time.Time {
	return startOfDay(now())
}

func startOfDay(origin time.Time) time.Time {
	return time.Date(origin.Year(), origin.Month(), origin.Day(), 0, 0, 0, 0, origin.Location())
}
