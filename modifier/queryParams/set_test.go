package queryParams

import (
	"net/url"
	"reflect"
	"testing"
)

type SetTestCase struct {
	SetFields      map[string]string
	ExpectedValues url.Values
}

var SetCases = []SetTestCase{
	{
		map[string]string{"limit": "11111", "newKey": "someParam"},
		url.Values{
			"price":  []string{"1000"},
			"limit":  []string{"50","11111"},
			"level1": []string{"qwerty"},
			"newKey": []string{"someParam"},
		},
	},
}

func TestSetExecute(t *testing.T) {
	for caseNum, c := range SetCases {
		set := Set{}
		for key, value := range c.SetFields {
			set[key] = value
		}

		testValues := url.Values{}
		for k, v := range originalQueryValues {
			testValues[k] = v
		}

		set.Execute(testValues)

		if false == reflect.DeepEqual(testValues, c.ExpectedValues) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
