package martian

import "context"

// NewContext returns a context wrapping the received parent
func NewContext(parent context.Context) *Context {
	return &Context{
		Context: parent,
	}
}

// Context provides information for a single request/response pair.
type Context struct {
	context.Context
	skipRoundTrip bool
}

// SkipRoundTrip flags the context to skip the round trip
func (c *Context) SkipRoundTrip() {
	c.skipRoundTrip = true
}

// SkippingRoundTrip returns the flag for skipping the round trip
func (c *Context) SkippingRoundTrip() bool {
	return c.skipRoundTrip
}

var _ context.Context = &Context{Context: context.Background()}
