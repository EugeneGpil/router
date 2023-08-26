package tests

import (
	"net/http"

	"github.com/EugeneGpil/router/app/ship/types"
)

func GetHttpHandler(message []byte) types.HttpHandler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(message)
	}

	return types.HttpHandler(handler)
}
