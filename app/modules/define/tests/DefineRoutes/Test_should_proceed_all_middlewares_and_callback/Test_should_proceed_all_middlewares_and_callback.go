package Test_should_proceed_all_middlewares_and_callback

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/define"
	"github.com/EugeneGpil/router/app/modules/define/tests/DefineRoutes"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	netUrl "net/url"
)

func Test_should_proceed_all_middlewares_and_callback(t *testing.T) {
	tester.SetTester(t)

	tests.AddRouteWithMiddlewares(DefineRoutes.Url, DefineRoutes.CallbackMessage, [][]byte{
		DefineRoutes.Middleware1Message,
		DefineRoutes.Middleware2Message,
	})

	define.DefineRoutes(DefineRoutes.Mux)

	assertCallback()
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

	tester.AssertLen(messages, 3)

	tester.AssertSame(messages[0], DefineRoutes.Middleware1Message)
	tester.AssertSame(messages[1], DefineRoutes.Middleware2Message)
	tester.AssertSame(messages[2], DefineRoutes.CallbackMessage)
}
