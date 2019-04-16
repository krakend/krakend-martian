package queryParams

import (
	"net/url"
	"reflect"
	"testing"
)

type RemoveTestCase struct {
	RemoveFields      []string
	ExpectedValues url.Values
}

var RemoveCases = []RemoveTestCase{
	{
		[]string{"limit", "newKey"},
		url.Values{
			"price":  []string{"1000"},
			"level1": []string{"qwerty"},
		},
	},
}

func TestRemoveExecute(t *testing.T) {
	for caseNum, c := range RemoveCases {
		remove := Remove{}
		for _, value := range c.RemoveFields {
			remove = append(remove, value)
		}

		testValues := url.Values{}
		for k, v := range originalQueryValues {
			testValues[k] = v
		}

		remove.Execute(testValues)

		if false == reflect.DeepEqual(testValues, c.ExpectedValues) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
