// Package time provides time wrapper that mockable.
package time

import "time"

// Provider struct.
type Provider struct {
	location *time.Location
}

// Now returns current time in specified location.
func (p *Provider) Now() time.Time {
	return time.Now().In(p.location)
}

// Date returns time corresponding to date.
func (p *Provider) Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, p.location)
}

// New returns new provider instance.
func New(options ...Option) *Provider {
	o := Options{
		Locaton: time.Local,
	}

	for _, option := range options {
		option(&o)
	}

	return &Provider{
		location: o.Locaton,
	}
}
