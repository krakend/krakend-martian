package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type WlTestCase struct {
	WhiteListFields []string
	ExpectedJSON    string
}

var wlJSONStr = `[
   {
      "id": 1,
      "price": 1,
      "name": 1,
      "name1": 1
   },
   {
      "id": 2,
      "price": 2,
      "name": 2,
      "name1": 2
   }
]`

var wlTestCases = []WlTestCase{
	{
		[]string{"price", "id"},
		`[
   {
      "id": 1,
      "price": 1
   },
   {
      "id": 2,
      "price": 2
   }
]`},
	{
		[]string{"vlerkvelrkvneklr"},
		`[{},{}]`},
}

func TestWlExecute(t *testing.T) {
	for caseNum, c := range wlTestCases {
		var body, expectedJSON []map[string]interface{}
		json.Unmarshal([]byte(wlJSONStr), &body)
		wl := WhiteList{}
		for _, wlField := range c.WhiteListFields {
			wl = append(wl, wlField)
		}
		wl.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJSON)

		if !reflect.DeepEqual(body, expectedJSON) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
