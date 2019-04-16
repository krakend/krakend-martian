package modifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devopsfaith/krakend/logging"
	"github.com/google/martian/parse"
	"io/ioutil"
	"github.com/devopsfaith/krakend-martian/modifier/postParams"
	"mime"
	"mime/multipart"
	"net/http"
)

func init() {
	parse.Register("post.modifier", postParamsModifierFromJSON)
}

// обязательно слайс, чтобы выполнялся по порядку
type postParamsModificationList struct {
	modifiers []postParams.ModifierInterface
	Logger    logging.Logger
}

type postModifierJSON struct {
	Scope            []parse.ModifierType `json:"scope"`
	ModificationList []postModifier       `json:"modifications"`
}

type postModifier struct {
	Type string
	Data interface{}
}

func (m *postParamsModificationList) SetLogger(l logging.Logger) {
	m.Logger = l

	for _, mm := range m.modifiers {
		if v, ok := mm.(LoggerAdder); ok {
			v.SetLogger(m.Logger)
		}
	}
}

func (m postParamsModificationList) ModifyRequest(req *http.Request) error {
	if req.Method == http.MethodPost {
		isMultipartForm := false
		parametersBuffer := new(bytes.Buffer)
		parametersWriter := multipart.NewWriter(parametersBuffer)

		_, params, err := mime.ParseMediaType(req.Header.Get(ContentHeader))
		if err != nil {
			m.Logger.Error(err)
		}
		if _, ok := params["boundary"]; ok {
			isMultipartForm = true
			if err = parametersWriter.SetBoundary(params["boundary"]); err != nil {
				m.Logger.Error(err)
			}
		}

		if err = req.ParseMultipartForm(32 << 20); err != nil {
			m.Logger.Error(err)
		}
		values := req.PostForm

		for _, modifier := range m.modifiers {
			modifier.Execute(values)
		}

		if len(values) > 0 {
			if isMultipartForm {
				for fieldName, valueArr := range values {
					for _, value := range valueArr {
						if err = parametersWriter.WriteField(fieldName, value); err != nil {
							m.Logger.Error(err)
						}
					}
				}
				if err = parametersWriter.Close(); err != nil {
					m.Logger.Error(err)
				}
			} else {
				parametersBuffer = bytes.NewBufferString(values.Encode())
			}

			if err := req.Body.Close(); err != nil {
				m.Logger.Error(err)
			}
			req.Body = ioutil.NopCloser(parametersBuffer)
		}
	}
	return nil
}

func postParamsModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &postModifierJSON{}
	ml := &postParamsModificationList{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	for _, mod := range msg.ModificationList {
		modifier, err := postParams.CreateModifier(mod.Type, mod.Data)
		if err != nil {
			return nil, fmt.Errorf(PostParamsModifierWarning, err, mod.Type)
		} else {
			if v, ok := modifier.(LoggerAdder); ok {
				v.SetLogger(ml.Logger)
			}
			ml.modifiers = append(ml.modifiers, modifier)
		}
	}

	r, err := parse.NewResult(ml, msg.Scope)
	if err != nil {
		return nil, err
	}

	return r, nil
}
