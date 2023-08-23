package DefineRoutes

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/define"
	"github.com/EugeneGpil/tester"

	netUrl "net/url"
)

var helloMessage = []byte("Hello")
var url = ""
var mux = http.NewServeMux()

func Test_should_define_routes(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	define.DefineRoutes(mux)

	assertCallback()
}

func addRoute() {
	callback := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(helloMessage)
	}

	add.AddRoute(http.MethodGet, url, callback)
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

	tester.AssertLen(writer.GetMessages(), 1)

	responseMessage := writer.GetMessages()[0]

	tester.AssertSame(responseMessage, helloMessage)
}
