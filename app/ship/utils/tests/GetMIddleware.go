package tests

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"
)

func GetMiddleware(message []byte) types.Middleware {
	return getMiddleware(message, true)
}

func GetMiddlewareFalse(message []byte) types.Middleware {
	return getMiddleware(message, false)
}

func getMiddleware(message []byte, returnValue bool) types.Middleware {
	middleware := func(writer http.ResponseWriter, request *http.Request) bool {
		writer.Write(message)

		return returnValue
	}

	return types.Middleware(middleware)
}
