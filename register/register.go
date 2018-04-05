package register

import "sync"

const (
	// ScopeRequest modifies an HTTP request.
	ScopeRequest Scope = "request"
	// ScopeResponse modifies an HTTP response.
	ScopeResponse Scope = "response"
)

type Register map[string]Component

type Scope string

type Component struct {
	Scope       []Scope
	NewFromJSON func(b []byte) (interface{}, error)
}

var (
	register = Register{}
	mutex    = &sync.RWMutex{}
)

func Set(name string, scope []Scope, f func(b []byte) (interface{}, error)) {
	mutex.Lock()
	register[name] = Component{
		Scope:       scope,
		NewFromJSON: f,
	}
	mutex.Unlock()
}

func Get() Register {
	mutex.RLock()
	r := make(Register, len(register))
	for k, v := range register {
		r[k] = v
	}
	mutex.RUnlock()
	return r
}
