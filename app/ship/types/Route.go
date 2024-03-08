package types

type Route struct {
	Method      string
	Url         string
	Callback    HttpHandler
	Name        string
	Middlewares []Middleware
}

type SimpleRoute struct {
	Method string
	Url    string
}
