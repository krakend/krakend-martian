package collection

import (
	"fmt"
	"strings"
)

const DataTransformError = "can't transform modifier data"

func CreateModifier(modifierType string, modifierData interface{}) (ModifierInterface, error) {
	switch strings.ToLower(modifierType) {
	case "blacklist":
		if rangedModifier, ok := modifierData.([]interface{}); ok {
			return createBlackList(rangedModifier), nil
		}
	case "whitelist":
		if rangedModifier, ok := modifierData.([]interface{}); ok {
			return createWhiteList(rangedModifier), nil
		}
	case "rename":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createRename(rangedModifier), nil
		}
	case "copy":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createCopy(rangedModifier), nil
		}
	case "changevalue":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createChangeValue(rangedModifier), nil
		}
	case "move":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createMove(rangedModifier), nil
		}
	case "convert":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createConvert(rangedModifier), nil
		}
	case "custom":
		if rangedModifier, ok := modifierData.(map[string]interface{}); ok {
			return createCustom(rangedModifier), nil
		}
	default:
		return nil, fmt.Errorf("unknown modifier: %s", modifierType)
	}
	return nil, fmt.Errorf(DataTransformError)
}

func createBlackList(rangedModifier []interface{}) *BlackList {
	bl := BlackList{}
	for _, val := range rangedModifier {
		bl = append(bl, val)
	}
	return &bl
}

func createWhiteList(rangedModifier []interface{}) *WhiteList {
	wl := WhiteList{}
	for _, val := range rangedModifier {
		wl = append(wl, val)
	}
	return &wl
}

func createRename(rangedModifier map[string]interface{}) *Rename {
	r := Rename{}
	for key, val := range rangedModifier {
		r[key] = val
	}
	return &r
}

func createCopy(rangedModifier map[string]interface{}) *Copy {
	c := Copy{}
	for key, val := range rangedModifier {
		c[key] = val
	}
	return &c
}

func createChangeValue(rangedModifier map[string]interface{}) *ChangeValue {
	cv := ChangeValue{}
	for key, val := range rangedModifier {
		cv[key] = val
	}
	return &cv
}

func createMove(rangedModifier map[string]interface{}) *Move {
	m := Move{}
	for key, val := range rangedModifier {
		m[key] = val
	}
	return &m
}

func createConvert(rangedModifier map[string]interface{}) *Convert {
	c := NewConvert()

	for key, val := range rangedModifier {
		c.List[key] = val
	}
	return c
}

func createCustom(rangedModifier map[string]interface{}) *CustomValue {
	cv := CustomValue{}
	for key, val := range rangedModifier {
		cv[key] = val
	}
	return &cv
}
