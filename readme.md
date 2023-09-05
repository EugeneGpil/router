# Simple router

Easily define routes for `net/http.ServeMux`

## Dummy example

```go
package main

import (
	"net/http"

	"github.com/EugeneGpil/router"
)

var method = http.MethodGet
var url = "hello/world"
var name = "hello.world"

var handler = func(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!\n"))
}

var mux = http.NewServeMux()

func main() {
	router.AddRoute(method, url, handler, name)

	router.DefineRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
```
