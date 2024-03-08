package types

import (
	"net/http"

	"github.com/EugeneGpil/response"
)

type HttpHandler func(response.Response, *http.Request)
type Middleware func(response.Response, *http.Request) bool
