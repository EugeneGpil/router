package add

import (
	"net/http"
	"strings"

	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/types"
)

func AddRoute(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	trimmedUrl := strings.Trim(url, "/")

	route := types.Route{
		Method:  method,
		Url:     trimmedUrl,
		Handler: handler,
	}

	globals.Routes = append(globals.Routes, route)
}
