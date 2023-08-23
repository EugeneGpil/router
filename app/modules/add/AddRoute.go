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
) types.Route {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:      method,
		Url:         trimmedUrl,
		Callback:    callback,
		Middlewares: []func(http.ResponseWriter, *http.Request) bool{},
	}

	routes.Add(route)

	return route
}
