package should_return_simple_route_by_name_if_use_AddRoute_function

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"
)

func Test_should_return_simple_route_by_name_if_use_AddRoute_function(t *testing.T) {
	tester.SetTester(t)

	router.AddRoute(
		http.MethodGet,
		tests.DefaultUrl,
		tests.DefaultCallback,
		tests.DefaultName,
	)

	route := router.ByName(tests.DefaultName)

	tester.AssertSame(tests.DefaultUrl, route.Url)
	tester.AssertSame(http.MethodGet, route.Method)
}
