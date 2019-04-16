package postParams

import (
	"net/url"
	"reflect"
	"testing"
)

type RenameTestCase struct {
	RenameFields   map[string]string
	ExpectedValues url.Values
}

var originalQueryValues = url.Values{
	"price":  []string{"1000"},
	"limit":  []string{"50"},
	"level1": []string{"qwerty"},
}

var RenameCases = []RenameTestCase{
	{
		map[string]string{"price": "bigSale", "level1": "someParam"},
		url.Values{
			"bigSale":   []string{"1000"},
			"limit":     []string{"50"},
			"someParam": []string{"qwerty"},
		},
	},
	{
		map[string]string{"abrakadabra": "12345"},
		url.Values{
			"price":  []string{"1000"},
			"limit":  []string{"50"},
			"level1": []string{"qwerty"},
		},
	},
}

func TestRenameExecute(t *testing.T) {
	for caseNum, c := range RenameCases {
		rename := Rename{}
		for key, value := range c.RenameFields {
			rename[key] = value
		}

		testValues := url.Values{}
		for k, v := range originalQueryValues {
			testValues[k] = v
		}

		rename.Execute(testValues)

		if false == reflect.DeepEqual(testValues, c.ExpectedValues) {
			t.Errorf("[%d] jsons aren't equal", caseNum)
		}
	}
}
