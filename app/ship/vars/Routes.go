package vars

import "github.com/EugeneGpil/router/app/ship/types"

var Routes routes

type routes struct {
	routes []types.Route
}

func (routes *routes) Add(route types.Route) *routes {
	routes.routes = append(routes.routes, route)

	return routes
}

func (routes *routes) GetAll() []types.Route {
	return routes.routes
}
