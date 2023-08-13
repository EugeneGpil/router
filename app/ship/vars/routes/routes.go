package routes

import "github.com/EugeneGpil/router/app/ship/types"

var routes []types.Route

func Add(route types.Route) {
	routes = append(routes, route)
}

func GetAll() []types.Route {
	return routes
}
