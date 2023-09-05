package routes

import "github.com/EugeneGpil/router/app/ship/types"

var routes = make(map[string]*types.Route)

func Add(route types.Route) {
	routes[route.Name] = &route
}

func GetAll() map[string]*types.Route {
	return routes
}
