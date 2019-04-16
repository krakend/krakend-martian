package collection

type ChangeValue map[string]interface{}

func (cv *ChangeValue) Execute(body []map[string]interface{}) {
	if nil != cv {
		for _, item := range body {
			for cvKey, cvValue := range *cv {
				if _, ok := item[cvKey]; ok {
					item[cvKey] = cvValue
				}
			}
		}
	}
}
