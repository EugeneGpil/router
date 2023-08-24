package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

func AddRoute(
	method string,
	url string,
	callback func(http.ResponseWriter, *http.Request),
) {
	middlewares := make([]func(http.ResponseWriter, *http.Request) bool, 0)

	AddRouteWitMiddlewares(method, url, callback, middlewares)
}

func AddRouteWitMiddlewares(
	method string,
	url string,
	callback func(http.ResponseWriter, *http.Request),
	middlewares []func(http.ResponseWriter, *http.Request) bool,
) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:      method,
		Url:         trimmedUrl,
		Callback:    callback,
		Middlewares: middlewares,
	}

	routes.Add(route)
}
