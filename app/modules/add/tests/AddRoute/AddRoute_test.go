package AddRoute

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
	"github.com/EugeneGpil/tester"
)

var helloMessage = []byte("Hello")
var url = ""
var routeName = "route.name"

func Test_should_add_route(t *testing.T) {
	tester.SetTester(t)

	addRoute()

	assertPrimitives()

	assertCallback()
}

func addRoute() {
	callback := tests.GetHttpHandler(helloMessage)

	add.AddRoute(http.MethodGet, url, callback, routeName)
}

func assertPrimitives() {
	routes := routes.GetAll()

	tester.AssertLen(routes, 1)

	route := routes[routeName]

	tester.AssertSame(route.Method, http.MethodGet)
	tester.AssertSame(route.Url, url)
	tester.AssertNotNil(route.Callback)
	tester.AssertLen(route.Middlewares, 0)
}

func assertCallback() {
	testResponseWriter := tester.GetTestResponseWriter()

	routes.GetAll()[routeName].Callback(&testResponseWriter, nil)

	responseMessage := testResponseWriter.GetMessages()[0]

	tester.AssertSame(responseMessage, helloMessage)
}
