package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type CvTestCase struct {
	ChangeValueFields map[string]interface{}
	ExpectedJSON      string
}

var CvJSONStr = `[
   {
      "id": 1,
      "price": 1,
      "name": 1,
      "name1": 1,
      "level1": {
         "inner": "qwerty",
         "level2": "test"
      }
   },
   {
      "id": 2,
      "price": 2,
      "name": 2,
      "name1": 2
   }
]`

var CvCases = []CvTestCase{
	{
		map[string]interface{}{"price": "changedPriceValue", "level1": "changedLevel1", "name": true},
		`[
   {
      "id": 1,
      "price": "changedPriceValue",
      "name": true,
      "name1": 1,
      "level1": "changedLevel1"
   },
   {
      "id": 2,
      "price": "changedPriceValue",
      "name": true,
      "name1": 2
   }
]`},
	{
		map[string]interface{}{"wecwcew": "vrveverver"},
		`[
   {
      "id": 1,
      "price": 1,
      "name": 1,
      "name1": 1,
      "level1": {
         "inner": "qwerty",
         "level2": "test"
      }
   },
   {
      "id": 2,
      "price": 2,
      "name": 2,
      "name1": 2
   }
]`},
}

func TestCvExecute(t *testing.T) {
	for caseNum, c := range CvCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(CvJSONStr), &body)
		cv := ChangeValue{}
		for cvKey, cvValue := range c.ChangeValueFields {
			cv[cvKey] = cvValue
		}
		cv.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
