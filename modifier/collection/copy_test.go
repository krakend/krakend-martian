package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type CopyTestCase struct {
	CopyFields   map[string]string
	ExpectedJSON string
}

var CopyJSONStr = `[
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

var CopyCases = []CopyTestCase{
	{
		map[string]string{"price": "bigSale", "level1": "someStructure"},
		`[
   {
      "id": 1,
      "price": 1,
      "bigSale": 1,
      "name": 1,
      "name1": 1,
      "level1": {
         "inner": "qwerty",
         "level2": "test"
      },
      "someStructure": {
         "inner": "qwerty",
         "level2": "test"
      }
   },
   {
      "id": 2,
      "price": 2,
      "bigSale": 2,
      "name": 2,
      "name1": 2
   }
]`},
	{
		map[string]string{"wecwcew": "vrveverver"},
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

func TestCopyExecute(t *testing.T) {
	for caseNum, c := range CopyCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(CopyJSONStr), &body)
		copyModifier := Copy{}
		for copyKey, copyValue := range c.CopyFields {
			copyModifier[copyKey] = copyValue
		}
		copyModifier.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
