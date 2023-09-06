package tests

import (
	"net/http"

	"github.com/EugeneGpil/router"
)

func AddRouteWithMiddlewares(handlerMessage []byte, middlewareMessages [][]byte) {
	callback := GetHttpHandler(handlerMessage)

	middlewares := GetMiddlewares(middlewareMessages)
	router.AddRouteWithMiddlewares(http.MethodGet, DefaultUrl, callback, middlewares, DefaultName)
}
