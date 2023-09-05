package router

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
)

func AddRoute(
	method string,
	url string,
	handler func(http.ResponseWriter, *http.Request),
	name string,
) {
	add.AddRoute(method, url, handler, name)
}

func DefineRoutes(mux *http.ServeMux) {
	define.DefineRoutes(mux)
}
