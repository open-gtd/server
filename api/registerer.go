package api

//Registerer represents ability to register methods in http API
type Registerer interface {
	GET(prefix string, path string, handlerFunc HandlerFunc)
	PATCH(prefix string, path string, handlerFunc HandlerFunc)
	POST(prefix string, path string, handlerFunc HandlerFunc)
	PUT(prefix string, path string, handlerFunc HandlerFunc)
	DELETE(prefix string, path string, handlerFunc HandlerFunc)
}

//Response represents ability to send response to user
type Response interface {
	String(code int, s string) error
	JSON(code int, i interface{}) error
	NoContent(code int) error
}

//Request represents ability to read request data
type Request interface {
	Param(name string) string
	Bind(i interface{}) error
}

//HandlerFunc function to handle request
type HandlerFunc func(Request, Response) error
