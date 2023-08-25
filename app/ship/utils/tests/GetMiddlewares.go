package tests

import "net/http"

func GetMiddlewares(messages [][]byte) []func(writer http.ResponseWriter, request *http.Request) bool {
	middlewares := make([]func(writer http.ResponseWriter, request *http.Request) bool, len(messages))
	for index, message := range messages {
		middlewares[index] = GetMiddleware(message)
	}

	return middlewares
}
