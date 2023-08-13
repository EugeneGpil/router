package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars"
)

func AddRoute(method string, url string, callback func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:   method,
		Url:      trimmedUrl,
		Callback: callback,
	}

	vars.Routes.Add(route)
}
