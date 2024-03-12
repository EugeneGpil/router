package tests

import (
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/router/app/ship/types"

	"github.com/EugeneGpil/response"
)

func GetHttpHandler(message []byte) types.HttpHandler {
	handler := func(response response.Response, request *request.Request) {
		response.Write(message)
	}

	return types.HttpHandler(handler)
}
