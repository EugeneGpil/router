package types

import "net/http"

type HttpHandler func(http.ResponseWriter, *http.Request)
type Middleware func(http.ResponseWriter, *http.Request) bool
