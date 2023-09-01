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

	netUrl "net/url"
)

func Test_should_proceed_only_middlewares_until_false_was_returned(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	define.DefineRoutes(DefineRoutes.Mux)

	assertCallback()
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

func assertCallback() {
	urlObj, err := netUrl.Parse("/")
	tester.AssertNil(err)

	request := http.Request{
		Method: http.MethodGet,
		URL:    urlObj,
	}

	handler, _ := DefineRoutes.Mux.Handler(&request)

	writer := tester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	messages := writer.GetMessages()

	tester.AssertLen(messages, 2)

	tester.AssertSame(messages[0], DefineRoutes.Middleware1Message)
	tester.AssertSame(messages[1], DefineRoutes.Middleware2Message)
}
