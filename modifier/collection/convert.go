package collection

import (
	"fmt"
	"github.com/devopsfaith/krakend/logging"
	"strconv"
)

type Convert struct {
	List   map[string]interface{}
	Logger logging.Logger
}

func (m *Convert) SetLogger(l logging.Logger) {
	m.Logger = l
}

func NewConvert() *Convert {
	list := make(map[string]interface{})
	return &Convert{List: list}
}

func (m *Convert) Execute(body []map[string]interface{}) {
	if nil != m {
		for _, item := range body {
			for field, toType := range m.List {
				if _, ok := item[field]; ok {
					switch toType {
					case "string":
						item[field] = m.convertInterfaceToString(item[field])
					case "int":
						item[field] = m.convertInterfaceToInt(item[field])
					default:
						m.Logger.Warning(fmt.Sprintf(ConvertModifierTypeWarning, toType))
					}
				}
			}
		}
	}
}

func (m Convert) convertInterfaceToInt(item interface{}) interface{} {
	switch item.(type) {
	case []interface{}:
		for key, val := range item.([]interface{}) {
			converted, err := convertToInt(val)
			if err != nil {
				m.Logger.Warning(fmt.Sprintf(ConvertModifierSliceWarning, "int", err))
			} else {
				item.([]interface{})[key] = converted
			}
		}
		return item
	case interface{}:
		converted, err := convertToInt(item)
		if err != nil {
			m.Logger.Warning(fmt.Sprintf(ConvertModifierValueWarning, "int", err))
			return item
		}
		return converted
	default:
		m.Logger.Warning(fmt.Sprintf(ConvertModifierOtherWarning, item, "int"))
		return item
	}
}

func (m Convert) convertInterfaceToString(item interface{}) interface{} {
	switch item.(type) {
	case []interface{}:
		for key, val := range item.([]interface{}) {
			converted, err := convertToString(val)
			if err != nil {
				m.Logger.Warning(fmt.Sprintf(ConvertModifierSliceWarning, "string", err))
			} else {
				item.([]interface{})[key] = converted
			}
		}
		return item
	case interface{}:
		converted, err := convertToString(item)
		if err != nil {
			m.Logger.Warning(fmt.Sprintf(ConvertModifierValueWarning, "string", err))
			return item
		}
		return converted
	default:
		m.Logger.Warning(fmt.Sprintf(ConvertModifierOtherWarning, item, "string"))
		return item
	}
}

func convertToInt(item interface{}) (float64, error) {
	switch item.(type) {
	case float64:
		return item.(float64), nil
	case string:
		return strconv.ParseFloat(item.(string), 64)
	default:
		return 0, fmt.Errorf(ConvertModifierOtherWarning, item, "int")
	}
}

func convertToString(item interface{}) (string, error) {
	switch item.(type) {
	case float64:
		return fmt.Sprintf("%v", int(item.(float64))), nil
	case string:
		return item.(string), nil
	default:
		return "", fmt.Errorf(ConvertModifierOtherWarning, item, "string")
	}
}
