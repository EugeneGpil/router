package tests

import (
	"github.com/EugeneGpil/router/app/ship/types"
)

func GetMiddlewares(messages [][]byte) []types.Middleware {
	middlewares := make([]types.Middleware, len(messages))
	for index, message := range messages {
		middlewares[index] = GetMiddleware(message)
	}

	return middlewares
}
