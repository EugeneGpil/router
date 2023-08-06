package app

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/types"
)

var mux *http.ServeMux

func DefineRoutes(m *http.ServeMux) {
	urlRoutesMap := make(map[string][]types.Route)
	mux = m

	for _, currentRoute := range globals.Routes {
		if _, isset := urlRoutesMap[currentRoute.Url]; !isset {
			urlRoutesMap[currentRoute.Url] = []types.Route{}
		}
		urlRoutesMap[currentRoute.Url] = append(urlRoutesMap[currentRoute.Url], currentRoute)
	}

	for url, routes := range urlRoutesMap {
		url1, url2 := getFormattedUrls.Run(url)

		addHandlers(url1, routes)
		addHandlers(url2, routes)
	}
}

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
