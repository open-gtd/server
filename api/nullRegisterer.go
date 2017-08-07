package api

type NullRegisterer struct {
}

func (NullRegisterer) GET(prefix string, path string, handlerFunc HandlerFunc)    {}
func (NullRegisterer) POST(prefix string, path string, handlerFunc HandlerFunc)   {}
func (NullRegisterer) PUT(prefix string, path string, handlerFunc HandlerFunc)    {}
func (NullRegisterer) DELETE(prefix string, path string, handlerFunc HandlerFunc) {}
