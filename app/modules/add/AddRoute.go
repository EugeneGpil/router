package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/ship"
	"github.com/EugeneGpil/router/app/ship/types"
)

func AddRoute(method string, url string, callback func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:   method,
		Url:      trimmedUrl,
		Callback: callback,
	}

	ship.Routes = append(ship.Routes, route)
}
