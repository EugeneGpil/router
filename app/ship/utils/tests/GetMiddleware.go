package tests

import (
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/router/app/ship/types"
)

func GetMiddleware(message []byte) types.Middleware {
	return getMiddleware(message, true)
}

func GetMiddlewareFalse(message []byte) types.Middleware {
	return getMiddleware(message, false)
}

func getMiddleware(message []byte, returnValue bool) types.Middleware {
	middleware := types.Middleware(func(response response.Response, request *request.Request) bool {
		response.Write(message)

		return returnValue
	})

	return types.Middleware(middleware)
}
