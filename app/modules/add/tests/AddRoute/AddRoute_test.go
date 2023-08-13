package AddRoute

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/test/types"
	"github.com/EugeneGpil/router/app/ship/vars/routes"
)

var helloMessage = []byte("Hello")
var url = ""
var tester *testing.T

func Test_should_add_route(t *testing.T) {
	tester = t

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

	if count := len(routes); count != 1 {
		tester.Fatalf(`routes count = %q, want match for %q`, count, 1)
	}

	route := routes[0]

	if route.Method != http.MethodGet {
		tester.Fatalf(`route.Method = %q, want match for %q`, route.Method, http.MethodGet)
	}

	if route.Url != url {
		tester.Fatalf(`route.Url = %q, want match for %q`, route.Url, url)
	}

	if route.Callback == nil {
		tester.Fatal(`route.Handler is nil`)
	}
}

func assertCallback() {
	testResponseWriter := types.ResponseWriter{}

	routes.GetAll()[0].Callback(&testResponseWriter, nil)

	responseMessage := testResponseWriter.GetMessages()[0]

	if bytes.Compare(responseMessage, helloMessage) != 0 {
		tester.Fatalf(`responseMessage = %q, want match for %q`, responseMessage, helloMessage)
	}
}
