package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

//NewEchoRegisterer returns API registerer for Echo Framework which allows to register endpoints at main level
func NewEchoRegisterer(e Router, groups map[string]Router) api.Registerer {
	return &registerer{echo: e, groups: groups}
}

type registerer struct {
	echo   Router
	groups map[string]Router
}

type Router interface {
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

func (r *registerer) GET(prefix string, path string, handler api.Handler) {
	r.add(echo.GET, prefix, path, handler)
}

func (r *registerer) POST(prefix string, path string, handler api.Handler) {
	r.add(echo.POST, prefix, path, handler)
}

func (r *registerer) PATCH(prefix string, path string, handler api.Handler) {
	r.add(echo.PATCH, prefix, path, handler)
}
func (r *registerer) PUT(prefix string, path string, handler api.Handler) {
	r.add(echo.PUT, prefix, path, handler)
}
func (r *registerer) DELETE(prefix string, path string, handler api.Handler) {
	r.add(echo.DELETE, prefix, path, handler)
}

func (r *registerer) add(method string, prefix string, path string, handler api.Handler) {

	if group, ok := r.groups[prefix]; ok {
		group.Add(method, path, func(c echo.Context) error {
			return handler.Handle(c, c)
		})
		return
	}

	r.echo.Add(method, prefix+path, func(c echo.Context) error {
		return handler.Handle(c, c)
	})
}
