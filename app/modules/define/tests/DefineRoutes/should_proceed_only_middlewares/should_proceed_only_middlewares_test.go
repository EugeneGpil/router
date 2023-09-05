package should_proceed_only_middlewares

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
	"github.com/EugeneGpil/router/app/modules/define/tests/DefineRoutes"
	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"
)

func Test_should_proceed_only_middlewares(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	define.DefineRoutes(DefineRoutes.Mux)

	DefineRoutes.AssertWriterMessages([][]byte{
		DefineRoutes.Middleware1Message,
		DefineRoutes.Middleware2Message,
		DefineRoutes.Middleware3Message,
	})
}

func addRoute() {
	callback := tests.GetHttpHandler(DefineRoutes.CallbackMessage)

	middlewares := getMiddlewares()

	add.AddRouteWitMiddlewares(
		http.MethodGet,
		DefineRoutes.Url,
		callback,
		DefineRoutes.RouteName,
		middlewares,
	)
}

func getMiddlewares() []types.Middleware {
	middleware1 := tests.GetMiddleware(DefineRoutes.Middleware1Message)
	middleware2 := tests.GetMiddleware(DefineRoutes.Middleware2Message)
	middleware3 := tests.GetMiddlewareFalse(DefineRoutes.Middleware3Message)

	return []types.Middleware{
		middleware1,
		middleware2,
		middleware3,
	}
}
