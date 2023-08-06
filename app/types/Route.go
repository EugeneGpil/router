package types

import "net/http"

type Route struct {
	Method  string
	Url     string
	Handler func(http.ResponseWriter, *http.Request)
}
