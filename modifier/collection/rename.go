package collection

type Rename map[string]interface{}

func (r *Rename) Execute(body []map[string]interface{}) {
	if nil != r {
		for key, item := range body {
			for from, to := range *r {
				v, ok := item[from]
				if ok {
					body[key][to.(string)] = v
					delete(body[key], from)
				}
			}
		}
	}
}
