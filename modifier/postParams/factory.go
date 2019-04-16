package postParams

import (
	"fmt"
	"strings"
)

const DataTransformError = "can't transform modifier data"

func CreateModifier(modifierType string, modifierData interface{}) (ModifierInterface, error) {
	switch strings.ToLower(modifierType) {
	case "rename":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createRename(rangedModifier), nil
		}
	case "set":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createSet(rangedModifier), nil
		}
	default:
		return nil, fmt.Errorf("unknown modifier: %s", modifierType)
	}
	return nil, fmt.Errorf(DataTransformError)
}

func createRename(rangedModifier map[string]interface{}) *Rename {
	r := Rename{}
	for key, val := range rangedModifier {
		r[key] = val
	}
	return &r
}

func createSet(rangedModifier map[string]interface{}) *Set {
	s := Set{}
	for key, val := range rangedModifier {
		s[key] = val
	}
	return &s
}
