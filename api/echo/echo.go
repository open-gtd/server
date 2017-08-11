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

func (r *registerer) GET(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.add(echo.GET, prefix, path, handlerFunc)
}

func (r *registerer) POST(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.add(echo.POST, prefix, path, handlerFunc)
}

func (r *registerer) PATCH(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.add(echo.PATCH, prefix, path, handlerFunc)
}
func (r *registerer) PUT(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.add(echo.PUT, prefix, path, handlerFunc)
}
func (r *registerer) DELETE(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.add(echo.DELETE, prefix, path, handlerFunc)
}

func (r *registerer) add(method string, prefix string, path string, handlerFunc api.HandlerFunc) {

	if group, ok := r.groups[prefix]; ok {
		group.Add(method, path, func(c echo.Context) error {
			return handlerFunc(c, c)
		})
		return
	}

	r.echo.Add(method, prefix+path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
