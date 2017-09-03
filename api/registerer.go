package api

type Handler interface {
	Handle(Request, Response) error
}

//Registerer represents ability to register methods in http API
type Registerer interface {
	GET(prefix string, path string, handler Handler)
	PATCH(prefix string, path string, handler Handler)
	POST(prefix string, path string, handler Handler)
	PUT(prefix string, path string, handler Handler)
	DELETE(prefix string, path string, handler Handler)
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
