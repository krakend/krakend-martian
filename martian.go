package martian

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	_ "github.com/google/martian/body"
	_ "github.com/google/martian/fifo"
	_ "github.com/google/martian/header"
	"github.com/google/martian/parse"
)

// NewBackendFactory creates a proxy.BackendFactory with the martian request executor wrapping the injected one.
// If there is any problem parsing the extra config data, it just uses the injected request executor.
func NewBackendFactory(logger logging.Logger, re proxy.HTTPRequestExecutor) proxy.BackendFactory {
	return func(remote *config.Backend) proxy.Proxy {
		result, ok := ConfigGetter(remote.ExtraConfig).(MartianResult)
		if !ok {
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, re, remote.Decoder)
		}
		switch result.Err {
		case nil:
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, HTTPRequestExecutor(result.Result, re), remote.Decoder)
		case ErrEmptyValue:
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, re, remote.Decoder)
		default:
			logger.Error(result, remote.ExtraConfig)
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, re, remote.Decoder)
		}
	}
}

// HTTPRequestExecutor creates a wrapper over the received request executor, so the martian modifiers can be
// executed before and after the execution of the request
func HTTPRequestExecutor(result *parse.Result, re proxy.HTTPRequestExecutor) proxy.HTTPRequestExecutor {
	return func(ctx context.Context, req *http.Request) (*http.Response, error) {
		result.RequestModifier().ModifyRequest(req)
		resp, err := re(ctx, req)
		result.ResponseModifier().ModifyResponse(resp)
		return resp, err
	}
}

// Namespace is the key to look for extra configuration details
const Namespace = "github.com/devopsfaith/krakend-martian"

// MartianResult is a simple wrapper over the parse.FromJSON response tuple
type MartianResult struct {
	Result *parse.Result
	Err    error
}

// ConfigGetter implements the config.ConfigGetter interface. It parses the extra config for the
// martian adapter and returns a MartianResult wrapping the results.
func ConfigGetter(e config.ExtraConfig) interface{} {
	cfg, ok := e[Namespace]
	if !ok {
		return MartianResult{nil, ErrEmptyValue}
	}

	data, ok := cfg.(map[string]interface{})
	if !ok {
		return MartianResult{nil, ErrBadValue}
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return MartianResult{nil, ErrMarshallingValue}
	}

	r, err := parse.FromJSON(raw)

	return MartianResult{r, err}
}

var (
	// ErrEmptyValue is the error returned when there is no config under the namespace
	ErrEmptyValue = fmt.Errorf("getting the extra config for the martian module")
	// ErrBadValue is the error returned when the config is not a map
	ErrBadValue = fmt.Errorf("casting the extra config for the martian module")
	// ErrMarshallingValue is the error returned when the config map can not be marshalled again
	ErrMarshallingValue = fmt.Errorf("marshalling the extra config for the martian module")
)
