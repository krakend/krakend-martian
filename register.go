package martian

import (
	"fmt"
	"reflect"

	"github.com/devopsfaith/krakend/register"
	"github.com/google/martian/parse"
)

// Register gets all the modifiers from the krakend-martian register and registers
// them into the martian parser
func Register(namespaced *register.Namespaced) {
	reg, ok := namespaced.Get(Namespace)
	if !ok {
		return
	}

	for k, tmp := range reg.Clone() {
		component, ok := tmp.(Component)
		if !ok {
			fmt.Println("the component", k, "is not a martian component:", reflect.TypeOf(tmp))
			continue
		}
		parse.Register(k, func(b []byte) (*parse.Result, error) {
			v, err := component.NewFromJSON(b)
			if err != nil {
				return nil, err
			}

			return parse.NewResult(v, toModifierType(component.Scopes()))
		})
	}
}

func toModifierType(scopes []string) []parse.ModifierType {
	modifierType := make([]parse.ModifierType, len(scopes))
	for k, s := range scopes {
		modifierType[k] = parse.ModifierType(s)
	}
	return modifierType
}

// Component defines the interface to be implemented by the registrable components
type Component interface {
	Scopes() []string
	NewFromJSON(b []byte) (interface{}, error)
}
