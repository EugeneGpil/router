package add

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"
)

func AddRoute(
	method string,
	url string,
	callback func(http.ResponseWriter, *http.Request),
) {
	middlewares := make([]types.Middleware, 0)

	AddRouteWitMiddlewares(method, url, callback, middlewares)
}
