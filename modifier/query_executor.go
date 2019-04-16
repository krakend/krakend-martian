package modifier

import (
	"encoding/json"
	"fmt"
	"github.com/devopsfaith/krakend/logging"
	"github.com/google/martian/parse"
	"github.com/devopsfaith/krakend-martian/modifier/queryParams"
	"net/http"
)

func init() {
	parse.Register("query.modifier", queryParamsModifierFromJSON)
}

// обязательно слайс, чтобы выполнялся по порядку
type queryParamsModificationList struct {
	modifiers []queryParams.ModifierInterface
	Logger    logging.Logger
}

type queryModifierJSON struct {
	Scope            []parse.ModifierType `json:"scope"`
	ModificationList []queryModifier      `json:"modifications"`
}

type queryModifier struct {
	Type string
	Data interface{}
}

func (m *queryParamsModificationList) SetLogger(l logging.Logger) {
	m.Logger = l

	for _, mm := range m.modifiers {
		if v, ok := mm.(LoggerAdder); ok {
			v.SetLogger(m.Logger)
		}
	}
}

func (m queryParamsModificationList) ModifyRequest(req *http.Request) error {
	if req.Method == http.MethodGet {
		values := req.URL.Query()
		for _, modifier := range m.modifiers {
			modifier.Execute(values)
		}
		req.URL.RawQuery = values.Encode()
	}
	return nil
}

func queryParamsModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &queryModifierJSON{}
	ml := &queryParamsModificationList{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	for _, mod := range msg.ModificationList {
		modifier, err := queryParams.CreateModifier(mod.Type, mod.Data)
		if err != nil {
			return nil, fmt.Errorf(QueryParamsModifierWarning, err, mod.Type)
		}

		if v, ok := modifier.(LoggerAdder); ok {
			v.SetLogger(ml.Logger)
		}
		ml.modifiers = append(ml.modifiers, modifier)
	}

	r, err := parse.NewResult(ml, msg.Scope)
	if err != nil {
		return nil, err
	}

	return r, nil
}
