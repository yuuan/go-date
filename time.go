package date

import (
	"time"
)

const iso8601 = "2006-01-02T15:04:05.999999999+09:00"

var (
	now      = time.Now
	location = func() *time.Location { return time.Local }
)

// Now returns the current time, which can be mocked using SetTestNow.
func Now() time.Time {
	return now()
}

// SetTestNow sets a custom function to return the current time.
// This function is used to mock the current time in tests.
func SetTestNow(getNow func() time.Time) {
	now = getNow
}

// ResetTestNow resets the function to return the current time to the default.
// This function is used to reset the mocked current time to its original value.
func ResetTestNow() {
	now = time.Now
}

// SetTestLocation sets a custom function to return the current location.
// This function is used to mock the current location in tests.
func SetTestLocation(getLocation func() *time.Location) {
	location = getLocation
}

// ResetTestLocation resets the function to return the current location to the default.
// This function is used to reset the mocked current location to its original value.
func ResetTestLocation() {
	location = func() *time.Location { return time.Local }
}

// today returns the current date with the time set to the start of the day.
func today() time.Time {
	return startOfDay(now())
}

// startOfDay returns the given time with the time set to the start of the day.
func startOfDay(origin time.Time) time.Time {
	return time.Date(origin.Year(), origin.Month(), origin.Day(), 0, 0, 0, 0, origin.Location())
}
