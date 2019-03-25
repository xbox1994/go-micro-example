package router

import (
	"net/http"

	"github.com/micro/go-api/resolver"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type defaultResolver struct {
	handler   string
	namespace string
}

func (r *defaultResolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	var name string

	switch r.handler {
	case "meta", "api", "rpc":
		name, _ = apiRoute(r.namespace, req.URL.Path)
	default:
		name = proxyRoute(r.namespace, req.URL.Path)
	}

	return &resolver.Endpoint{
		Name: name,
	}, nil
}

func (r *defaultResolver) String() string {
	return "default"
}
