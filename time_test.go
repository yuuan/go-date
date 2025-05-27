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

	subject := now()

	assert.Equal(t, subject, mocked, "Current time factory was not replaced by the set function")
}

func TestResetTestNow(t *testing.T) {
	mocked := time.Date(2024, time.June, 5, 12, 0, 0, 0, time.Local)
	factory := func() time.Time {
		return mocked
	}
	SetTestNow(factory)

	ResetTestNow()

	subject := now()

	assert.NotEqual(t, subject, mocked, "Current time factory was replaced")
}

func TestSetTestLocation(t *testing.T) {
	mocked, _ := time.LoadLocation("Asia/Tokyo")
	factory := func() *time.Location {
		return mocked
	}

	SetTestLocation(factory)

	subject := location()

	assert.Equal(t, subject, mocked, "Location factory was not replaced by the set function")
}

func TestResetTestLocation(t *testing.T) {
	mocked, _ := time.LoadLocation("Asia/Tokyo")
	factory := func() *time.Location {
		return mocked
	}
	SetTestLocation(factory)

	ResetTestLocation()

	subject := location()

	assert.NotEqual(t, subject, mocked, "Location factory was replaced")
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
