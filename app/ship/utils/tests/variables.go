package tests

import "net/http"

var DefaultUrl = "/"

var DefaultName = "default.name"

var DefaultCallback = func(http.ResponseWriter, *http.Request) {}
