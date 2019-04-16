package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type RenameTestCase struct {
	RenameFields map[string]string
	ExpectedJSON string
}

var RenameJSONStr = `[
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

var RenameCases = []RenameTestCase{
	{
		map[string]string{"price": "bigSale", "level1": "someStructure"},
		`[
   {
      "id": 1,
      "bigSale": 1,
      "name": 1,
      "name1": 1,
      "someStructure": {
         "inner": "qwerty",
         "level2": "test"
      }
   },
   {
      "id": 2,
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

func TestRenameExecute(t *testing.T) {
	for caseNum, c := range RenameCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(RenameJSONStr), &body)
		rename := Rename{}
		for renameKey, renameValue := range c.RenameFields {
			rename[renameKey] = renameValue
		}
		rename.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
