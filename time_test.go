package time_test

import (
	"log"
	"testing"
	gotime "time"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/time"
)

func TestProvider_Now(t *testing.T) {
	t.Parallel()

	before := time.New(time.UTC).Now()

	gotime.Sleep(gotime.Nanosecond)

	after := time.New().Now()

	assert.True(t, after.After(before))
}

func TestProvider_Date(t *testing.T) {
	t.Parallel()

	loc, err := gotime.LoadLocation("Europe/Moscow")
	assert.NoError(t, err)

	expected := gotime.Now().In(loc)
	year, month, day := expected.Date()
	hour, min, sec := expected.Clock()

	actual := time.New(time.WithLocation(loc)).
		Date(year, month, day, hour, min, sec, expected.Nanosecond())

	assert.Equal(t, expected.UnixNano(), actual.UnixNano())
}

func ExampleProvider_Now() {
	now := time.New(time.UTC).Now()
	log.Println(now)
}

func ExampleProvider_Date() {
	date := time.New(time.UTC).Date(2020, 1, 1, 0, 0, 0, 0)
	log.Println(date)
}
