package methods

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
	"github.com/EugeneGpil/router/app/ship/types"
)

func DefineGroupedByUrlRoutes(groupedByUrlRoutes map[string][]types.Route, mux *http.ServeMux) {
	for url, routes := range groupedByUrlRoutes {
		url1, url2 := getFormattedUrls.Run(url)

		addRoutes(url1, routes, mux)
		addRoutes(url2, routes, mux)
	}
}
