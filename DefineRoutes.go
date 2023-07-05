package router

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
)

var mux *http.ServeMux
var routes []Route

func DefineRoutes (m *http.ServeMux) {
	urlRoutesMap := make(map[string][]Route)
	mux = m

	for _, route := range routes {
		if _, isset := urlRoutesMap[route.url]; !isset {
			urlRoutesMap[route.url] = []Route{}
		}
		urlRoutesMap[route.url] = append(urlRoutesMap[route.url], route)
	}

	for url, routes := range urlRoutesMap {
		url1, url2 := getFormattedUrls.Run(url)

		addHandlers(url1, routes)
		addHandlers(url2, routes)
	}
}

func addHandlers(url string, routes []Route) {
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		for _, route := range routes {
			if request.Method == route.method {
				route.handler(writer, request)

				return
			}
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
