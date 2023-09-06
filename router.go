package router

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
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
	middlewares []types.Middleware,
	name string,
) {
	add.AddRouteWithMiddlewares(method, url, handler, name, middlewares)
}

func DefineRoutes(mux *http.ServeMux) {
	define.DefineRoutes(mux)
}
