package tests

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"

	"github.com/EugeneGpil/response"
)

func GetHttpHandler(message []byte) types.HttpHandler {
	handler := func(response response.Response, request *http.Request) {
		response.Write(message)
	}

	return types.HttpHandler(handler)
}
