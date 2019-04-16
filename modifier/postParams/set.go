package postParams

import "net/url"

type Set map[string]interface{}

func (s *Set) Execute(values url.Values) {
	if nil != s {
		for key, value := range *s {
			values.Set(key, value.(string))
		}
	}
}
