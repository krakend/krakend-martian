package martian

import "context"

func NewContext(parent context.Context) *Context {
	return &Context{
		Context: parent,
	}
}

type Context struct {
	context.Context
	skipRoundTrip bool
}

func (c *Context) SkipRoundTrip() {
	c.skipRoundTrip = true
}

func (c *Context) SkippingRoundTrip() bool {
	return c.skipRoundTrip
}

var _ context.Context = &Context{Context: context.Background()}
