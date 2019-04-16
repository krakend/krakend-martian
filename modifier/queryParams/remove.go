package queryParams

import "net/url"

type Remove []interface{}

func (r *Remove) Execute(values url.Values) {
	if nil != r {
		for _, key := range *r {
			if _, ok := values[key.(string)]; ok {
				values.Del(key.(string))
			}
		}
	}
}
