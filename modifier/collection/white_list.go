package collection

type WhiteList []interface{}

func (wl *WhiteList) Execute(body []map[string]interface{}) {
	wlMap := getMapFromSlice(*wl)
	if nil != wl {
		for key, item := range body {
			for field := range item {
				if _, ok := wlMap[field]; !ok {
					delete(body[key], field)
				}
			}
		}
	}
}

func getMapFromSlice(sl []interface{}) map[string]bool {
	mp := make(map[string]bool)
	for _, v := range sl {
		mp[v.(string)] = true
	}
	return mp
}
