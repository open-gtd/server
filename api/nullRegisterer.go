package api

type NullRegisterer struct {
}

func (NullRegisterer) GET(prefix string, path string, handler Handler)    {}
func (NullRegisterer) PATCH(prefix string, path string, handler Handler)  {}
func (NullRegisterer) POST(prefix string, path string, handler Handler)   {}
func (NullRegisterer) PUT(prefix string, path string, handler Handler)    {}
func (NullRegisterer) DELETE(prefix string, path string, handler Handler) {}
