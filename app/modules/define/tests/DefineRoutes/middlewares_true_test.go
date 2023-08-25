package DefineRoutes

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	netUrl "net/url"
)

var middleware1Message = []byte("Hello middleware1")
var middleware2Message = []byte("Hello middleware2")
var callbackMessage = []byte("Hello callback")

var url = ""
var mux = http.NewServeMux()

func Test_should_proceed_all_middlewares_and_callback(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	define.DefineRoutes(mux)

	assertCallback()
}

func addRoute() {
	callback := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(callbackMessage)
	}

	middlewares := tests.GetMiddlewares([][]byte{
		middleware1Message,
		middleware2Message,
	})
	add.AddRouteWitMiddlewares(http.MethodGet, url, callback, middlewares)
}

func assertCallback() {
	urlObj, err := netUrl.Parse("/")
	tester.AssertNil(err)

	request := http.Request{
		Method: http.MethodGet,
		URL:    urlObj,
	}

	handler, _ := mux.Handler(&request)

	writer := tester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	messages := writer.GetMessages()

	tester.AssertLen(messages, 3)

	tester.AssertSame(messages[0], middleware1Message)
	tester.AssertSame(messages[1], middleware2Message)
	tester.AssertSame(messages[2], callbackMessage)
}
