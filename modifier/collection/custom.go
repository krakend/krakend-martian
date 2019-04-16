package collection

const (
	menuPrice = "menu_price"
)

type CustomValue map[string]interface{}

func (cv *CustomValue) Execute(body []map[string]interface{}) {
	if nil != cv {
		for _, item := range body {
			for key, method := range *cv {
				if _, ok := item[key]; ok {
					switch method {
					case menuPrice:
						item[key] = menuPriceMethod(item, item[key])
					}
				}
			}
		}
	}
}

func menuPriceMethod(item map[string]interface{}, price interface{}) interface{} {
	var (
		value    interface{}
		currency string
	)

	if template, ok := item["template"]; ok {
		if classID, ok := template.(float64); ok {
			switch classID {
			case 2:
				if _, ok := item["points"]; ok {
					value = item["points"]
					currency = "POINTS"
				}
			default:
				if _, ok := item["price"]; ok {
					value = item["price"]
					currency = "RUB"
				}
			}
		}

		price = map[string]interface{}{"value": value, "currency": currency}
	}

	return price
}
