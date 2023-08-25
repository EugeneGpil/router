package types

import "net/http"

type Middleware func(http.ResponseWriter, *http.Request) bool
