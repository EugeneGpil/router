# Simple router

Easily define routes for `net/http.ServeMux`

## Dummy example

```go
package main

import (
	"net/http"

	"github.com/EugeneGpil/router"
)

func main() {
	router.AddRoute(http.MethodGet, "hello/world", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	})

	mux := http.NewServeMux()

	router.DefineRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
```
