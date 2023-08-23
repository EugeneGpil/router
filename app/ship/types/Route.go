package types

import "net/http"

type Route struct {
	Method      string
	Url         string
	Callback    func(http.ResponseWriter, *http.Request)
	Middlewares []func(http.ResponseWriter, *http.Request) bool
}

func (route *Route) SetMiddlewares(middlewares []func(http.ResponseWriter, *http.Request) bool) {
	route.Middlewares = middlewares
}

func (route *Route) AddMiddleware(middleware func(http.ResponseWriter, *http.Request) bool) {
	route.Middlewares = append(route.Middlewares, middleware)
}
