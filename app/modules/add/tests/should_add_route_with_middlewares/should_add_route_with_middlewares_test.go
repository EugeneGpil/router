package should_add_route_with_middlewares

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var url = ""
var handlerMessage = []byte("hello handler")
var middlewareMessage = []byte("hello middleware")

func Test_should_add_route_with_middlewares(t *testing.T) {
	tester.SetTester(t)

	tests.AddRouteWithMiddlewares(handlerMessage, [][]byte{
		middlewareMessage,
	})

	assertPrimitives()

	//TODO: поправить и вернуть эту проверку
	// assertCallback()
}

func assertPrimitives() {
	routes := routes.GetAll()

	tester.AssertLen(routes, 1)

	route := routes[tests.DefaultName]

	tester.AssertSame(route.Method, http.MethodGet)
	tester.AssertSame(route.Url, url)
	tester.AssertNotNil(route.Callback)
	tester.AssertLen(route.Middlewares, 1)
}

func assertCallback() {
	testResponseWriter := httpTester.GetTestResponseWriter()

	response := response.New(testResponseWriter)

	//TODO:
	//тут так не проверить миддлвары
	//т.к. сейчас запускается стразу Callback
	//чтобы норм проверить миддлвары
	//надо регистрировать роут в mux и запускать через mux
	routes.GetAll()[tests.DefaultName].Callback(response, nil)

	messages := testResponseWriter.GetMessages()

	tester.AssertLen(messages, 2)

	tester.AssertSame(messages[0], middlewareMessage)
	tester.AssertSame(messages[1], handlerMessage)
}
