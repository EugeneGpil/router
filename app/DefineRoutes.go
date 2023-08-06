package app

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/types"
)

var mux *http.ServeMux

func DefineRoutes(m *http.ServeMux) {
	groupedByUrlRoutesWithDifferentHttpMethods := getGroupedByUrlRoutesWithDifferentHttpMethods()
	mux = m

	for url, routes := range groupedByUrlRoutesWithDifferentHttpMethods {
		url1, url2 := getFormattedUrls.Run(url)

		addHandlers(url1, routes)
		addHandlers(url2, routes)
	}
}

func getGroupedByUrlRoutesWithDifferentHttpMethods() map[string][]types.Route {
	groupedByUrlRoutesWithDifferentHttpMethods := make(map[string][]types.Route)

	for _, currentRoute := range globals.Routes {
		if _, isset := groupedByUrlRoutesWithDifferentHttpMethods[currentRoute.Url]; !isset {
			groupedByUrlRoutesWithDifferentHttpMethods[currentRoute.Url] = []types.Route{}
		}
		appended := append(groupedByUrlRoutesWithDifferentHttpMethods[currentRoute.Url], currentRoute)
		groupedByUrlRoutesWithDifferentHttpMethods[currentRoute.Url] = appended
	}

	return groupedByUrlRoutesWithDifferentHttpMethods
}

func defineRouteMap()

func addHandlers(url string, routes []types.Route) {
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		for _, route := range routes {
			if request.Method == route.Method {
				route.Handler(writer, request)

				return
			}
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
