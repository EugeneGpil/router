package get

import (
	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

func ByName(routeName string) types.SimpleRoute {
	route := routes.GetAll()[routeName]

	return types.SimpleRoute{
		Method: route.Method,
		Url:    "/" + route.Url,
	}
}
