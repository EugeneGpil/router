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
			addRoutes(url, routes, mux)
		}
	}
}
