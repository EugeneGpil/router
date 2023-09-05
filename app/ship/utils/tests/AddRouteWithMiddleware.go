package tests

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/add"
)

func AddRouteWithMiddlewares(handlerMessage []byte, middlewareMessages [][]byte) {
	callback := GetHttpHandler(handlerMessage)

	middlewares := GetMiddlewares(middlewareMessages)
	add.AddRouteWitMiddlewares(http.MethodGet, DefaultUrl, callback, DefaultName, middlewares)
}
