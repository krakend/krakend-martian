package collection

type BlackList []interface{}

func (bl *BlackList) Execute(body []map[string]interface{}) {
	if nil != bl {
		for key, item := range body {
			for _, blField := range *bl {
				if _, ok := item[blField.(string)]; ok {
					delete(body[key], blField.(string))
				}
			}
		}
	}
}
