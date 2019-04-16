package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type BlTestCase struct {
	BlackListFields []string
	ExpectedJSON    string
}

var BlJSONStr = `[
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

var BlCases = []BlTestCase{
	{
		[]string{"price", "level1"},
		`[
   {
      "id": 1,
      "name": 1,
      "name1": 1
   },
   {
      "id": 2,
      "name": 2,
      "name1": 2
   }
]`},
	{
		[]string{"vlerkvelrkvneklr"},
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
	{
		[]string{"id", "price", "name", "name1", "level1"},
		`[{},{}]`},
}

func TestBlExecute(t *testing.T) {
	for caseNum, c := range BlCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(BlJSONStr), &body)
		bl := BlackList{}
		for _, blField := range c.BlackListFields {
			bl = append(bl, blField)
		}
		bl.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
