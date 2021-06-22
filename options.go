package time

import "time"

// Options struct.
type Options struct {
	Locaton *time.Location
}

// Option type.
type Option func(*Options)

// WithLocation option to set custom location.
func WithLocation(l *time.Location) Option {
	return func(o *Options) {
		o.Locaton = l
	}
}

// UTC sets location to UTC timezone.
func UTC(o *Options) {
	WithLocation(time.UTC)(o)
}
