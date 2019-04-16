package modifier

import (
	"encoding/json"
	"github.com/google/martian"
	"github.com/google/martian/parse"
	"net/http"
)

func init() {
	parse.Register("header.modifier", headerModifierFromJSON)
}

type HeaderModifier struct {
	headers map[string]string
}

type HeaderModifierJSON struct {
	Scope   []parse.ModifierType `json:"scope"`
	Headers map[string]string    `json:"headers"`
}

func (hm *HeaderModifier) ModifyRequest(req *http.Request) error {
	setHeaders(hm.headers, &req.Header)
	return nil
}

func (hm *HeaderModifier) ModifyResponse(res *http.Response) error {
	setHeaders(hm.headers, &res.Header)
	return nil
}

func setHeaders(headers map[string]string, httpHeader *http.Header) {
	for header, value := range headers {
		origHeader := httpHeader.Get(header)
		if origHeader == "" {
			httpHeader.Set(header, value)
		}
	}
}

func HeaderNewModifier(headers map[string]string) martian.RequestResponseModifier {
	return &HeaderModifier{
		headers: headers,
	}
}

func headerModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &HeaderModifierJSON{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(HeaderNewModifier(msg.Headers), msg.Scope)
}
