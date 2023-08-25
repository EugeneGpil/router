package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

func AddRouteWitMiddlewares(
	method string,
	url string,
	callback func(http.ResponseWriter, *http.Request),
	middlewares []types.Middleware,
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
