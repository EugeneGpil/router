package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/types"
)

func AddRoute(method string, url string, callback func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:   method,
		Url:      trimmedUrl,
		Callback: callback,
	}

	globals.Routes = append(globals.Routes, route)
}
