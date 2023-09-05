package DefineRoutes

import "net/http"

var Middleware1Message = []byte("Hello middleware1")
var Middleware2Message = []byte("Hello middleware2")
var Middleware3Message = []byte("Hello middleware3")
var CallbackMessage = []byte("Hello callback")

var Url = ""
var RouteName = "route.name"

var Mux = http.NewServeMux()
