package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/EugeneGpil/router/app/globals"
	"github.com/EugeneGpil/router/app/modules/add"
	"github.com/EugeneGpil/router/app/modules/test"
)

func Test_should_add_route(t *testing.T) {
	helloMessage := []byte("Hello")

	callback := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(helloMessage)
	}

	url := ""

	add.AddRoute(http.MethodGet, url, callback)

	route := globals.Routes[0]

	if route.Method != http.MethodGet {
		t.Fatalf(`route.Method = %q, want match for %q`, route.Method, http.MethodGet)
	}

	if route.Url != url {
		t.Fatalf(`route.Url = %q, want match for %q`, route.Url, url)
	}

	if route.Callback == nil {
		t.Fatal(`route.Handler is nil`)
	}

	testResponseWriter := test.ResponseWriter{}

	route.Callback(&testResponseWriter, nil)

	responseMessage := testResponseWriter.GetMessages()[0]

	if bytes.Compare(responseMessage, helloMessage) != 0 {
		t.Fatalf(`responseMessage = %q, want match for %q`, responseMessage, helloMessage)
	}
}
