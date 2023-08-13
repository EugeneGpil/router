package AddRoute

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/test/types"
	"github.com/EugeneGpil/router/app/modules/test/vars/tester"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

var helloMessage = []byte("Hello")
var url = ""
var testingTester *testing.T

func Test_should_add_route(t *testing.T) {
	testingTester = t
	tester.SetTester(t)

	addRoute()

	assertCallback()
}

func addRoute() {
	callback := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(helloMessage)
	}

	add.AddRoute(http.MethodGet, url, callback)
}

func assertPrimitives() {
	routes := routes.GetAll()

	tester.AssertLen(routes, 1)

	route := routes[0]

	tester.AssertSame(route.Method, http.MethodGet)
	tester.AssertSame(route.Url, url)
	tester.AssertNotNil(route.Callback)
}

func assertCallback() {
	testResponseWriter := types.ResponseWriter{}

	routes.GetAll()[0].Callback(&testResponseWriter, nil)

	responseMessage := testResponseWriter.GetMessages()[0]

	tester.AssertSame(responseMessage, helloMessage)
}
