package should_return_simple_route_by_name_if_use_AddRouteWithMiddlewares_function

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/router/app/ship/types"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"
)

func Test_should_return_simple_route_by_name_if_use_AddRouteWithMiddlewares_function(t *testing.T) {
	tester.SetTester(t)

	router.AddRouteWithMiddlewares(
		http.MethodGet,
		tests.DefaultUrl,
		tests.DefaultCallback,
		tests.DefaultName,
		[]types.Middleware{},
	)

	route := router.ByName(tests.DefaultName)

	tester.AssertSame(tests.DefaultUrl, route.Url)
	tester.AssertSame(http.MethodGet, route.Method)
}
