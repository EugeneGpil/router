package app

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/types"
)

func DefineRoutes(mux *http.ServeMux) {
	groupedByUrlRoutes := getGroupedByUrlRoutes()

	defineRoutes(groupedByUrlRoutes, mux)
}

func getGroupedByUrlRoutes() map[string][]types.Route {
	groupedByUrlRoutes := make(map[string][]types.Route)

	for _, currentRoute := range globals.Routes {
		if _, isset := groupedByUrlRoutes[currentRoute.Url]; !isset {
			groupedByUrlRoutes[currentRoute.Url] = []types.Route{}
		}
		appended := append(groupedByUrlRoutes[currentRoute.Url], currentRoute)
		groupedByUrlRoutes[currentRoute.Url] = appended
	}

	return groupedByUrlRoutes
}

func defineRoutes(groupedByUrlRoutes map[string][]types.Route, mux *http.ServeMux) {
	for url, routes := range groupedByUrlRoutes {
		url1, url2 := getFormattedUrls.Run(url)

		addRoutes(url1, routes, mux)
		addRoutes(url2, routes, mux)
	}
}

func addRoutes(url string, routes []types.Route, mux *http.ServeMux) {
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
