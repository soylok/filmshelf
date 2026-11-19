package server

import (
	"fmt"
	"net/http"
)

type Version string

var (
	Version1 = Version("v1")
)

type Router struct {
	mux *http.ServeMux
	v   Version
}

func NewRouter(v Version) *Router {
	return &Router{
		mux: http.NewServeMux(),
		v:   v,
	}
}

func (r *Router) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Endpoint)
		r.mux.Handle(pattern, route.Handler)
	}
}
