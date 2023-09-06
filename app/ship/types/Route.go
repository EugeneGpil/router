package types

import (
	"net/http"
)

type Route struct {
	Method      string
	Url         string
	Callback    func(http.ResponseWriter, *http.Request)
	Name        string
	Middlewares []Middleware
}

type SimpleRoute struct {
	Method string
	Url    string
}
