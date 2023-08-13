package methods

import (
	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars"
)

func GetGroupedByUrlRoutes() map[string][]types.Route {
	groupedByUrlRoutes := make(map[string][]types.Route)

	for _, currentRoute := range vars.Routes.GetAll() {
		if _, isset := groupedByUrlRoutes[currentRoute.Url]; !isset {
			groupedByUrlRoutes[currentRoute.Url] = []types.Route{}
		}
		appended := append(groupedByUrlRoutes[currentRoute.Url], currentRoute)
		groupedByUrlRoutes[currentRoute.Url] = appended
	}

	return groupedByUrlRoutes
}
