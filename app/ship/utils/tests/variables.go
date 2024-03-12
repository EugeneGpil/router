package tests

import (
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/router/app/ship/types"
)

var DefaultUrl = "/"

var DefaultName = "default.name"

var DefaultCallback = types.HttpHandler(func(response.Response, *request.Request) {})
