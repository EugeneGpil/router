package add

import (
	"strings"

	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

func AddRouteWithMiddlewares(
	method string,
	url string,
	callback types.HttpHandler,
	name string,
	middlewares []types.Middleware,
) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:      method,
		Url:         trimmedUrl,
		Callback:    callback,
		Name:        name,
		Middlewares: middlewares,
	}

	routes.Add(route)
}
