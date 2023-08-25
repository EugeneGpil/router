package tests

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"
)

func GetMiddleware(message []byte) types.Middleware {
	middleware := func(writer http.ResponseWriter, request *http.Request) bool {
		writer.Write(message)

		return true
	}

	return types.Middleware(middleware)
}
