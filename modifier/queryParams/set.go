package queryParams

import "net/url"

type Set map[string]interface{}

func (s *Set) Execute(values url.Values) {
	if nil != s {
		for key, value := range *s {
			switch value.(type) {
			case []interface{}:
				for _, v := range value.([]interface{}) {
					values.Add(key, v.(string))
				}
			default:
				values.Add(key, value.(string))
			}
		}
	}
}
