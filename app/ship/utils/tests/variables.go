package tests

import (
	"net/http"

	"github.com/EugeneGpil/response"
)

var DefaultUrl = "/"

var DefaultName = "default.name"

var DefaultCallback = func(response.Response, *http.Request) {}
