package martian

import (
	"encoding/json"
	"net/http"

	"github.com/google/martian/parse"
	"github.com/google/martian/static"
)

// StaticModifier is a martian.RequestResponseModifier that routes reqeusts to rootPath
// and serves the assets there, while skipping the HTTP roundtrip.
type StaticModifier struct {
	*static.Modifier
}

type staticJSON struct {
	ExplicitPaths map[string]string    `json:"explicitPaths"`
	RootPath      string               `json:"rootPath"`
	Scope         []parse.ModifierType `json:"scope"`
}

// NewStaticModifier constructs a static.Modifier that takes a path to serve files from, as well as an optional mapping of request paths to local
// file paths (still rooted at rootPath).
func NewStaticModifier(rootPath string) *StaticModifier {
	return &StaticModifier{
		Modifier: static.NewModifier(rootPath),
	}
}

// ModifyRequest marks the context to skip the roundtrip and downgrades any https requests
// to http.
func (s *StaticModifier) ModifyRequest(req *http.Request) error {
	ctx := NewContext(req.Context())
	ctx.SkipRoundTrip()

	if req.URL.Scheme == "https" {
		req.URL.Scheme = "http"
	}

	*req = *req.WithContext(ctx)

	return nil
}

func staticModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &staticJSON{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	mod := NewStaticModifier(msg.RootPath)
	mod.SetExplicitPathMappings(msg.ExplicitPaths)
	return parse.NewResult(mod, msg.Scope)
}
