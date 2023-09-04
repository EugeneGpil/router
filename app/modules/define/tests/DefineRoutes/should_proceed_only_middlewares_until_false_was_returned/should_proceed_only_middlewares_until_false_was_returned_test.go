package DefineRoutes

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

func Test_should_proceed_only_middlewares_until_false_was_returned(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	define.DefineRoutes(DefineRoutes.Mux)

	DefineRoutes.AssertWriterMessages([][]byte{
		DefineRoutes.Middleware1Message,
		DefineRoutes.Middleware2Message,
	})
}

func addRoute() {
	callback := tests.GetHttpHandler(DefineRoutes.CallbackMessage)

	middleware1 := tests.GetMiddleware(DefineRoutes.Middleware1Message)
	middleware2 := tests.GetMiddlewareFalse(DefineRoutes.Middleware2Message)
	middleware3 := tests.GetMiddleware(DefineRoutes.Middleware3Message)

	add.AddRouteWitMiddlewares(http.MethodGet, DefineRoutes.Url, callback, []types.Middleware{
		middleware1,
		middleware2,
		middleware3,
	})
}
