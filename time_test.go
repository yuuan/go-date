package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetTestNow(t *testing.T) {
	mocked := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	factory := func() time.Time {
		return mocked
	}

	SetTestNow(factory)

	assert.Equal(t, now(), mocked, "Current time factory was not replaced by the set function")
}

func TestResetTestNow(t *testing.T) {
	mocked := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	factory := func() time.Time {
		return mocked
	}
	SetTestNow(factory)

	ResetTestNow()

	assert.NotEqual(t, now(), mocked, "Current time factory was replaced")
}

func TestSetTestLocation(t *testing.T) {
	mocked, _ := time.LoadLocation("Asia/Tokyo")
	factory := func() *time.Location {
		return mocked
	}

	SetTestLocation(factory)

	assert.Equal(t, location(), mocked, "Location factory was not replaced by the set function")
}

func TestResetTestLocation(t *testing.T) {
	mocked, _ := time.LoadLocation("Asia/Tokyo")
	factory := func() *time.Location {
		return mocked
	}
	SetTestLocation(factory)

	ResetTestLocation()

	assert.NotEqual(t, location(), mocked, "Location factory was replaced")
}

func TestNow(t *testing.T) {
	mocked := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	factory := func() time.Time {
		return mocked
	}

	SetTestNow(factory)
	defer ResetTestNow()

	subject := Now()

	assert.Equal(t, mocked, subject, "Now() should return mocked time when SetTestNow is used")
}
