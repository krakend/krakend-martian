package collection

import (
	"encoding/json"
	"reflect"
	"testing"
)

type MoveTestCase struct {
	MoveFields   map[string]string
	ExpectedJSON string
}

var MoveJSONStr = `[
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

var MoveCases = []MoveTestCase{
	{
		map[string]string{"price": "bigSale", "level1.level2": "someStructure", "id": "identificator.primary.id"},
		`[
   {
      "bigSale": 1,
      "name": 1,
      "name1": 1,
      "level1": {
         "inner": "qwerty"
      },
      "someStructure": "test",
      "identificator": {
         "primary": {
            "id": 1
         }
      }
   },
   {
      "bigSale": 2,
      "name": 2,
      "name1": 2,
      "identificator": {
         "primary": {
            "id": 2
         }
      }
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

func TestMoveExecute(t *testing.T) {
	for caseNum, c := range MoveCases {
		var body, expectedJs []map[string]interface{}
		json.Unmarshal([]byte(MoveJSONStr), &body)
		move := Move{}
		for originalItemPath, newItemPath := range c.MoveFields {
			move[originalItemPath] = newItemPath
		}
		move.Execute(body)
		json.Unmarshal([]byte(c.ExpectedJSON), &expectedJs)

		if !reflect.DeepEqual(body, expectedJs) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
