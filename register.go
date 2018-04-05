package martian

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/google/martian/parse"
)

func Register() {
	for k, component := range register.Get() {
		parse.Register(k, func(b []byte) (*parse.Result, error) {
			v, err := component.NewFromJSON(b)
			if err != nil {
				return nil, err
			}

			modifierType := make([]parse.ModifierType, len(component.Scope))
			for k, s := range component.Scope {
				modifierType[k] = parse.ModifierType(s)
			}

			return parse.NewResult(v, modifierType)
		})
	}
}
