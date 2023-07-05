package router

import "net/http"

type Route struct {
	method string
	url string
	handler func(http.ResponseWriter, *http.Request)
}
