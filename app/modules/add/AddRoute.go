package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/ship/types"
)

func AddRoute(method string, url string, callback func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:   method,
		Url:      trimmedUrl,
		Callback: callback,
	}

	types.Routes = append(types.Routes, route)
}
