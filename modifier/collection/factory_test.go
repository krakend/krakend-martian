package collection

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

type BlackListCase struct {
	modifierData  []interface{}
	modifierError string
	ExpectedData  *BlackList
}

type WhiteListCase struct {
	modifierData  []interface{}
	modifierError string
	ExpectedData  *WhiteList
}

type RenameCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Rename
}

type CopyCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Copy
}
type ChangeCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *ChangeValue
}
type MoveCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Move
}
type ConvertCase struct {
	modifierData  map[string]interface{}
	modifierError string
	ExpectedData  *Convert
}

var (
	createModifierCases = []Case{
		{"blacklist", "", DataTransformError, nil},
		{"blacklist", []interface{}{"1", "2", "3"}, "", &BlackList{}},
		{"blacklist", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, DataTransformError, nil},
		{"whitelist", "", DataTransformError, nil},
		{"whitelist", []interface{}{"1", "2", "3"}, "", &WhiteList{}},
		{"whitelist", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, DataTransformError, nil},
		{"rename", "", DataTransformError, nil},
		{"rename", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"rename", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Rename{}},
		{"copy", "", DataTransformError, nil},
		{"copy", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"copy", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Copy{}},
		{"changevalue", "", DataTransformError, nil},
		{"changevalue", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"changevalue", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &ChangeValue{}},
		{"move", "", DataTransformError, nil},
		{"move", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"move", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Move{}},
		{"convert", "", DataTransformError, nil},
		{"convert", []interface{}{"1", "2", "3"}, DataTransformError, nil},
		{"convert", map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Convert{}},
		{"test", "", "unknown modifier: test", nil},
	}
	createBlackListCases = []BlackListCase{
		{[]interface{}{"1", "2", "3"}, "", &BlackList{"1", "2", "3"}},
		{[]interface{}{"4", "5", "6"}, "", &BlackList{"4", "5", "6"}},
	}
	createWhiteListCases = []WhiteListCase{
		{[]interface{}{"1", "2", "3"}, "", &WhiteList{"1", "2", "3"}},
		{[]interface{}{"4", "5", "6"}, "", &WhiteList{"4", "5", "6"}},
	}
	createRenameCases = []RenameCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Rename{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Rename{"a": "4", "b": "5", "c": "6"}},
	}
	createCopyCases = []CopyCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Copy{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Copy{"a": "4", "b": "5", "c": "6"}},
	}
	createChangeCases = []ChangeCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &ChangeValue{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &ChangeValue{"a": "4", "b": "5", "c": "6"}},
	}
	createMoveCases = []MoveCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Move{"a": "1", "b": "2", "c": "3"}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Move{"a": "4", "b": "5", "c": "6"}},
	}
	first              = map[string]interface{}{"a": "1", "b": "2", "c": "3"}
	second             = map[string]interface{}{"a": "4", "b": "5", "c": "6"}
	createConvertCases = []ConvertCase{
		{map[string]interface{}{"a": "1", "b": "2", "c": "3"}, "", &Convert{List: first}},
		{map[string]interface{}{"a": "4", "b": "5", "c": "6"}, "", &Convert{List: second}},
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

func TestCreateBlackList(t *testing.T) {
	for _, c := range createBlackListCases {
		m := createBlackList(c.modifierData)

		if !reflect.DeepEqual(m, c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateWhiteList(t *testing.T) {
	for _, c := range createWhiteListCases {
		m := createWhiteList(c.modifierData)

		if !reflect.DeepEqual(m, c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateRename(t *testing.T) {
	for _, c := range createRenameCases {
		m := createRename(c.modifierData)

		if !eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateCopy(t *testing.T) {
	for _, c := range createCopyCases {
		m := createRename(c.modifierData)

		if !eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateChangeValue(t *testing.T) {
	for _, c := range createChangeCases {
		m := createRename(c.modifierData)

		if !eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateMove(t *testing.T) {
	for _, c := range createMoveCases {
		m := createMove(c.modifierData)

		if !eq(*m, *c.ExpectedData) {
			t.Errorf("aren't equal %v and %v", c.ExpectedData, m)
		}
	}
}

func TestCreateConvert(t *testing.T) {
	for _, c := range createConvertCases {
		m := createConvert(c.modifierData)

		if !eq(m.List, c.ExpectedData.List) {
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
