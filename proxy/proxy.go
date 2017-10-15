package proxy

import (
	"context"
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/google/martian/parse"

	"github.com/devopsfaith/krakend-martian"
)

// NewBackendFactory creates a proxy.BackendFactory with the martian request executor wrapping the injected one.
// If there is any problem parsing the extra config data, it just uses the injected request executor.
func NewBackendFactory(logger logging.Logger, re proxy.HTTPRequestExecutor) proxy.BackendFactory {
	return func(remote *config.Backend) proxy.Proxy {
		result, err := martian.Parse(remote.ExtraConfig, Namespace)
		switch err {
		case nil:
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, HTTPRequestExecutor(result, re), remote.Decoder)
		case martian.ErrEmptyValue:
			return proxy.NewHTTPProxyWithHTTPExecutor(remote, re, remote.Decoder)
		default:
			logger.Error(err, remote.ExtraConfig)
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
const Namespace = "github.com/krakend-martian/proxy"
