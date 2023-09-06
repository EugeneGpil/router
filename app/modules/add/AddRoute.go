package add

import (
	"github.com/EugeneGpil/router/app/ship/types"
)

func AddRoute(
	method string,
	url string,
	callback types.HttpHandler,
	name string,
) {
	middlewares := make([]types.Middleware, 0)

	AddRouteWithMiddlewares(method, url, callback, name, middlewares)
}
