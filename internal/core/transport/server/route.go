package server

import "net/http"

type Route struct {
	Method   string
	Endpoint string
	Handler  http.HandlerFunc
}
