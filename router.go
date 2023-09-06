package router

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
	"github.com/EugeneGpil/router/app/modules/get"
	"github.com/EugeneGpil/router/app/ship/types"
)

func AddRoute(
	method string,
	url string,
	handler types.HttpHandler,
	name string,
) {
	add.AddRoute(method, url, handler, name)
}

func AddRouteWithMiddlewares(
	method string,
	url string,
	handler types.HttpHandler,
	name string,
	middlewares []types.Middleware,
) {
	add.AddRouteWithMiddlewares(method, url, handler, name, middlewares)
}

func ByName(routeName string) types.SimpleRoute {
	return get.ByName(routeName)
}

func DefineRoutes(mux *http.ServeMux) {
	define.DefineRoutes(mux)
}
