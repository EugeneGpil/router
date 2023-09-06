package should_proceed_all_middlewares_and_callback

import (
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/router/app/modules/define/tests/DefineRoutes"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"
)

func Test_should_proceed_all_middlewares_and_callback(t *testing.T) {
	tester.SetTester(t)

	tests.AddRouteWithMiddlewares(DefineRoutes.CallbackMessage, [][]byte{
		DefineRoutes.Middleware1Message,
		DefineRoutes.Middleware2Message,
	})

	router.DefineRoutes(DefineRoutes.Mux)

	DefineRoutes.AssertWriterMessages([][]byte{
		DefineRoutes.Middleware1Message,
		DefineRoutes.Middleware2Message,
		DefineRoutes.CallbackMessage,
	})
}
