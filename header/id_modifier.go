package header

import (
	"encoding/json"
	"net/http"

	"github.com/google/martian"
	"github.com/google/martian/parse"
	"github.com/google/uuid"
)

const defaultHeaderName string = "X-Krakend-Id"

type idModifier struct {
	header string
}

type idModifierJSON struct {
	Scope  []parse.ModifierType `json:"scope"`
	Header string               `json:"header"`
}

// NewIDModifier returns a request modifier that will set a header with the name
// X-Krakend-Id with a value that is a unique identifier for the request. In the case
// that the X-Krakend-Id header is already set, the header is unmodified.
func NewIDModifier(header string) martian.RequestModifier {
	if header == "" {
		header = defaultHeaderName
	}
	return &idModifier{header: header}
}

// ModifyRequest sets the X-Krakend-Id header with a unique identifier.  In the case
// that the X-Krakend-Id header is already set, the header is unmodified.
func (im *idModifier) ModifyRequest(req *http.Request) error {
	// Do not rewrite an ID if req already has one
	if req.Header.Get(im.header) != "" {
		return nil
	}

	id := uuid.New()
	req.Header.Set(im.header, id.String())

	return nil
}

func IdModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &idModifierJSON{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	modifier := NewIDModifier(msg.Header)

	return parse.NewResult(modifier, msg.Scope)
}
