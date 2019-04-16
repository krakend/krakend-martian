package collection

type Copy map[string]interface{}

func (c *Copy) Execute(body []map[string]interface{}) {
	if nil != c {
		for key, item := range body {
			for from, to := range *c {
				v, ok := item[from]
				if ok {
					body[key][to.(string)] = v
				}
			}
		}
	}
}
