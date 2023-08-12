package methods

import (
	"github.com/EugeneGpil/router/app/ship"
	"github.com/EugeneGpil/router/app/ship/types"
)

func GetGroupedByUrlRoutes() map[string][]types.Route {
	groupedByUrlRoutes := make(map[string][]types.Route)

	for _, currentRoute := range ship.Routes {
		if _, isset := groupedByUrlRoutes[currentRoute.Url]; !isset {
			groupedByUrlRoutes[currentRoute.Url] = []types.Route{}
		}
		appended := append(groupedByUrlRoutes[currentRoute.Url], currentRoute)
		groupedByUrlRoutes[currentRoute.Url] = appended
	}

	return groupedByUrlRoutes
}
