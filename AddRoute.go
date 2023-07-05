package router

import (
	"net/http"
	"strings"
)

func AddRoute(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := Route{
		method: method,
		url: trimmedUrl,
		handler: handler,
	}

	routes = append(routes, route)
}
