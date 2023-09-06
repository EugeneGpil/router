package should_proceed_only_middlewares_until_false_was_returned

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	defineTests "github.com/EugeneGpil/router/app/modules/define/tests"
)

func Test_should_proceed_only_middlewares_until_false_was_returned(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	router.DefineRoutes(defineTests.Mux)

	defineTests.AssertWriterMessages([][]byte{
		defineTests.Middleware1Message,
		defineTests.Middleware2Message,
	})
}

func addRoute() {
	callback := tests.GetHttpHandler(defineTests.CallbackMessage)

	middlewares := getMiddlewares()

	add.AddRouteWithMiddlewares(
		http.MethodGet,
		defineTests.Url,
		callback,
		defineTests.RouteName,
		middlewares,
	)
}

func getMiddlewares() []types.Middleware {
	middleware1 := tests.GetMiddleware(defineTests.Middleware1Message)
	middleware2 := tests.GetMiddlewareFalse(defineTests.Middleware2Message)
	middleware3 := tests.GetMiddleware(defineTests.Middleware3Message)

	return []types.Middleware{
		middleware1,
		middleware2,
		middleware3,
	}
}
