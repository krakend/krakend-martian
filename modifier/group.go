package modifier

import (
	"encoding/json"
	"fmt"
	"github.com/devopsfaith/krakend/logging"
	"github.com/google/martian"
	"github.com/google/martian/parse"
	"net/http"
	"sync"
)

func init() {
	parse.Register("fifo.LGroup", groupFromJSON)
}

type Group struct {
	reqMu   sync.RWMutex
	reqMods []martian.RequestModifier

	resMu   sync.RWMutex
	resMods []martian.ResponseModifier

	Logger logging.Logger

	aggregateErrors bool
}

func NewGroup() *Group {
	return &Group{}
}

type groupJSON struct {
	Modifiers       []json.RawMessage    `json:"modifiers"`
	Scope           []parse.ModifierType `json:"scope"`
	AggregateErrors bool                 `json:"aggregateErrors"`
}

func (g *Group) SetLogger(l logging.Logger) {
	g.Logger = l

	for _, mm := range g.reqMods {
		if v, ok := mm.(LoggerAdder); ok {
			v.SetLogger(g.Logger)
		}
	}

	for _, mm := range g.resMods {
		if v, ok := mm.(LoggerAdder); ok {
			v.SetLogger(g.Logger)
		}
	}
}

func (g *Group) SetAggregateErrors(aggerr bool) {
	g.aggregateErrors = aggerr
}

func (g *Group) AddRequestModifier(reqmod martian.RequestModifier) {
	g.reqMu.Lock()
	defer g.reqMu.Unlock()

	if v, ok := reqmod.(LoggerAdder); ok {
		v.SetLogger(g.Logger)
	}

	g.reqMods = append(g.reqMods, reqmod)
}

func (g *Group) AddResponseModifier(resmod martian.ResponseModifier) {
	g.resMu.Lock()
	defer g.resMu.Unlock()

	if v, ok := resmod.(LoggerAdder); ok {
		v.SetLogger(g.Logger)
	}

	g.resMods = append(g.resMods, resmod)
}

func (g *Group) ModifyRequest(req *http.Request) error {
	g.reqMu.RLock()
	defer g.reqMu.RUnlock()

	merr := martian.NewMultiError()

	for _, reqmod := range g.reqMods {
		if err := reqmod.ModifyRequest(req); err != nil {
			if g.aggregateErrors {
				merr.Add(err)
				continue
			}

			return err
		}
	}
	g.Logger.Debug(fmt.Sprintf("fifo.LGroup: %v", req))

	if merr.Empty() {
		return nil
	}

	return merr
}

func (g *Group) ModifyResponse(res *http.Response) error {
	g.resMu.RLock()
	defer g.resMu.RUnlock()

	merr := martian.NewMultiError()

	for _, resmod := range g.resMods {
		if err := resmod.ModifyResponse(res); err != nil {
			if g.aggregateErrors {
				merr.Add(err)
				continue
			}

			return err
		}
	}

	if merr.Empty() {
		return nil
	}

	return merr
}

// groupFromJSON builds a fifo.Group from JSON.
//
// Example JSON:
// {
//   "fifo.LGroup" : {
//     "scope": ["request", "result"],
//     "modifiers": [
//       { ... },
//       { ... },
//     ]
//   }
// }
func groupFromJSON(b []byte) (*parse.Result, error) {
	msg := &groupJSON{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	g := NewGroup()
	if msg.AggregateErrors {
		g.SetAggregateErrors(true)
	}

	for _, m := range msg.Modifiers {
		r, err := parse.FromJSON(m)
		if err != nil {
			return nil, err
		}

		reqmod := r.RequestModifier()
		if reqmod != nil {
			g.AddRequestModifier(reqmod)
		}

		resmod := r.ResponseModifier()
		if resmod != nil {
			g.AddResponseModifier(resmod)
		}
	}

	return parse.NewResult(g, msg.Scope)
}
