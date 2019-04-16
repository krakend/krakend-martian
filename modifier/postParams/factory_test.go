package postParams

import (
	"reflect"
	"testing"
)

type Case struct {
	modifierCaseType string
	modifierData     interface{}
	modifierError    string
	modifierType     interface{}
}

type RenameCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Rename
}

type SetCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Set
}

var (
	createModifierCases = []Case{
		{"rename", "", DataTransformError, nil},
		{"rename", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"rename", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Rename{}},
		{"set", "", DataTransformError, nil},
		{"set", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"set", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Set{}},
		{"test", "", "unknown modifier: test", nil},
	}
	createRenameCases = []RenameCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Rename{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Rename{"a": "4", "b": "5", "c": "6"}},
	}
	createSetCases = []SetCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Set{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Set{"a": "4", "b": "5", "c": "6"}},
	}
)

func TestCreateModifier(t *testing.T) {
	for _, c := range createModifierCases {
		m, err := CreateModifier(c.modifierCaseType, c.modifierData)

		if (err != nil && (c.modifierError == "" || err.Error() != c.modifierError)) || (err == nil && c.modifierError != "") {
			t.Errorf("except %s and %s", err, c.modifierError)
		}

		if reflect.TypeOf(m) != reflect.TypeOf(c.modifierType) {
			t.Errorf("except %s and %s", reflect.TypeOf(m), reflect.TypeOf(c.modifierType))
		}
	}
}

func TestCreateRename(t *testing.T) {
	for _, c := range createRenameCases {
		m := createRename(c.modifierData)

		if false == eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateSet(t *testing.T) {
	for _, c := range createSetCases {
		m := createRename(c.modifierData)

		if false == eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func eq(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}
