package register

const (
	// ScopeRequest modifies an HTTP request.
	ScopeRequest = "request"
	// ScopeResponse modifies an HTTP response.
	ScopeResponse = "response"

	// Namespace is the key to look for extra configuration details
	Namespace = "github.com/devopsfaith/krakend-martian"
)

// NewComponent returns component ready to be injected into the register
func NewComponent(scope []string, newFromJSON func(b []byte) (interface{}, error)) *Component {
	return &Component{
		scope:       scope,
		newFromJSON: newFromJSON,
	}
}

// Component contains the scope and the module factory
type Component struct {
	scope       []string
	newFromJSON func(b []byte) (interface{}, error)
}

// NewFromJSON implements the martian.Component interface
func (c *Component) NewFromJSON(b []byte) (interface{}, error) {
	return c.newFromJSON(b)
}

// Scopes implements the martian.Component interface
func (c *Component) Scopes() []string {
	return c.scope
}
