package types

import (
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/response"
)

type HttpHandler func(response.Response, *request.Request)
type Middleware func(response.Response, *request.Request) bool
