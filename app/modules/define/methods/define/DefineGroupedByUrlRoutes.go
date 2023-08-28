package define

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/router/app/ship/types"
)

func DefineGroupedByUrlRoutes(groupedByUrlRoutes map[string][]types.Route, mux *http.ServeMux) {
	for url, routes := range groupedByUrlRoutes {
		urls := getFormattedUrls.Run(url)

		for _, url := range urls {
			defineRoutesWithDifferentHttpMethodsOnOneUrl(url, routes, mux)
		}
	}
}

func defineRoutesWithDifferentHttpMethodsOnOneUrl(url string, routes []types.Route, mux *http.ServeMux) {
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		findRouteWithMatchingHttpMethodAndRun(routes, writer, request)
	})
}

func findRouteWithMatchingHttpMethodAndRun(routes []types.Route, writer http.ResponseWriter, request *http.Request) {
	for _, route := range routes {
		if request.Method == route.Method {
			run(route, writer, request)

			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func run(route types.Route, writer http.ResponseWriter, request *http.Request) {
	for _, callback := range route.Middlewares {
		callbackReslut := callback(writer, request)

		if callbackReslut == false {
			return
		}
	}

	route.Callback(writer, request)
}
