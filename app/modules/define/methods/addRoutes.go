package methods

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"
)

func addRoutes(url string, routes []types.Route, mux *http.ServeMux) {
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		for _, route := range routes {
			if request.Method == route.Method {
				route.Callback(writer, request)

				return
			}
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
