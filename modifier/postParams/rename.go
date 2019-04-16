package postParams

import "net/url"

type Rename map[string]interface{}

func (r *Rename) Execute(values url.Values) {
	if nil != r {
		for from, to := range *r {
			if _, ok := values[from]; ok {
				value := values.Get(from)
				values.Del(from)
				values.Set(to.(string), value)
			}
		}
	}
}
