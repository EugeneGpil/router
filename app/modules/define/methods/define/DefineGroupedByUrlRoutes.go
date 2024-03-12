package define

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/router/app/ship/types"

	requestPackage "github.com/EugeneGpil/request"
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
		response := response.New(writer)

		requestRequest := requestPackage.New(request)

		findRouteWithMatchingHttpMethodAndRun(routes, response, &requestRequest)
	})
}

func findRouteWithMatchingHttpMethodAndRun(routes []types.Route, response response.Response, request *requestPackage.Request) {
	for _, route := range routes {
		if request.GetMethod() == route.Method {
			run(route, response, request)

			return
		}
	}

	response.SetStatusCode(http.StatusNotFound)
}

func run(route types.Route, response response.Response, request *requestPackage.Request) {
	for _, middleware := range route.Middlewares {
		middlewareResult := middleware(response, request)

		if !middlewareResult {
			return
		}
	}

	route.Callback(response, request)
}
