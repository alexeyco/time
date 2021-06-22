package time_test

import (
	"testing"
	gotime "time"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/time"
)

func TestUTC(t *testing.T) {
	t.Parallel()

	var o time.Options

	time.WithLocation(gotime.UTC)(&o)

	assert.Equal(t, gotime.UTC, o.Locaton)
}
