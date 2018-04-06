package martian

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/google/martian/parse"
)

// Register gets all the modifiers from the krakend-martian register and registers
// them into the martian parser
func Register() {
	for k, component := range register.Get() {
		parse.Register(k, func(b []byte) (*parse.Result, error) {
			v, err := component.NewFromJSON(b)
			if err != nil {
				return nil, err
			}

			return parse.NewResult(v, toModifierType(component.Scope))
		})
	}
}

func toModifierType(scopes []register.Scope) []parse.ModifierType {
	modifierType := make([]parse.ModifierType, len(scopes))
	for k, s := range scopes {
		modifierType[k] = parse.ModifierType(s)
	}
	return modifierType
}
