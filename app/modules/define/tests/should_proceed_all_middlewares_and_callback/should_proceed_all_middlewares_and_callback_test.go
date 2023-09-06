package should_proceed_all_middlewares_and_callback

import (
	"testing"

	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/router/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	defineTests "github.com/EugeneGpil/router/app/modules/define/tests"
)

func Test_should_proceed_all_middlewares_and_callback(t *testing.T) {
	tester.SetTester(t)

	tests.AddRouteWithMiddlewares(defineTests.CallbackMessage, [][]byte{
		defineTests.Middleware1Message,
		defineTests.Middleware2Message,
	})

	router.DefineRoutes(defineTests.Mux)

	defineTests.AssertWriterMessages([][]byte{
		defineTests.Middleware1Message,
		defineTests.Middleware2Message,
		defineTests.CallbackMessage,
	})
}
