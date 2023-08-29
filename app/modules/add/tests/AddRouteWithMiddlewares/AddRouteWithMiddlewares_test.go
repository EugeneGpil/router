package AddRouteWithMiddlewares

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

var url = ""
var handlerMessage = "hello handler"
var middlewareMessage = "hello middleware"

func Test_should_add_route_with_middlewares(t *testing.T) {
	tester.SetTester(t)
}
