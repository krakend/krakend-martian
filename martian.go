package martian

import (
	"encoding/json"
	"fmt"

	"github.com/devopsfaith/krakend/config"
	_ "github.com/google/martian/body"
	_ "github.com/google/martian/fifo"
	_ "github.com/google/martian/header"
	"github.com/google/martian/parse"
)

// Parse transforms the received config into a JSON string and calls the parse.FromJSON method.
// It replaces the usual ConfigGetter
func Parse(e config.ExtraConfig, namespace string) (*parse.Result, error) {
	cfg, ok := e[namespace]
	if !ok {
		return nil, ErrEmptyValue
	}

	data, ok := cfg.(map[string]interface{})
	if !ok {
		return nil, ErrBadValue
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return nil, ErrMarshallingValue
	}

	return parse.FromJSON(raw)
}

var (
	// ErrEmptyValue is the error returned when there is no config under the namespace
	ErrEmptyValue = fmt.Errorf("getting the extra config for the martian module")
	// ErrBadValue is the error returned when the config is not a map
	ErrBadValue = fmt.Errorf("casting the extra config for the martian module")
	// ErrMarshallingValue is the error returned when the config map can not be marshalled again
	ErrMarshallingValue = fmt.Errorf("marshalling the extra config for the martian module")
)
