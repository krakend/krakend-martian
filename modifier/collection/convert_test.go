package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type ConvertTestCase struct {
	ConvertFields map[string]string
	ExpectedJSON  string
}

var ConvertJSONStr = `[
   {
      "id": 1,
      "price": 1,
      "name": "1",
      "name1": 1,
      "array_int": [
         1,
         2,
         3
      ],
      "array_string": [
         "1",
         "2",
         "3"
      ]
   },
   {
      "id": 2,
      "price": 2,
      "name": "2",
      "name1": 2,
      "array_int": [
         4,
         5,
         6
      ],
      "array_string": [
         "4",
         "5",
         "6"
      ]
   }
]`

var ConvertCases = []ConvertTestCase{
	{
		map[string]string{"price": "string", "name": "int", "array_int": "string", "array_string": "int"},
		`[
   {
      "id": 1,
      "price": "1",
      "name": 1,
      "name1": 1,
      "array_int": [
         "1",
         "2",
         "3"
      ],
      "array_string": [
         1,
         2,
         3
      ]
   },
   {
      "id": 2,
      "price": "2",
      "name": 2,
      "name1": 2,
      "array_int": [
         "4",
         "5",
         "6"
      ],
      "array_string": [
         4,
         5,
         6
      ]
   }
]`},
	{
		map[string]string{"wecwcew": "vrveverver"},
		`[
   {
      "id": 1,
      "price": 1,
      "name": "1",
      "name1": 1,
      "array_int": [
         1,
         2,
         3
      ],
      "array_string": [
         "1",
         "2",
         "3"
      ]
   },
   {
      "id": 2,
      "price": 2,
      "name": "2",
      "name1": 2,
      "array_int": [
         4,
         5,
         6
      ],
      "array_string": [
         "4",
         "5",
         "6"
      ]
   }
]`},
}

func TestConvertExecute(t *testing.T) {
	for caseNum, c := range ConvertCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(ConvertJSONStr), &body)
		convertModifier := NewConvert()
		for convertKey, convertValue := range c.ConvertFields {
			convertModifier.List[convertKey] = convertValue
		}
		convertModifier.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
