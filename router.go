package router

import (
	"net/http"

	"github.com/EugeneGpil/router/app"
)

func AddRoute(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	app.AddRoute(method, url, handler)
}

func DefineRoutes(mux *http.ServeMux) {
	app.DefineRoutes(mux)
}
