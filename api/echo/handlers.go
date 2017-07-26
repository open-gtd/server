package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

func NewRegisterer(e *echo.Echo, prefix string, m ...echo.MiddlewareFunc) api.Registerer {
	return &Registerer{group: e.Group(prefix, m...)}
}

type Registerer struct {
	group *echo.Group
}

func (r *Registerer) GET(path string, handlerFunc api.HandlerFunc) {
	r.group.GET(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *Registerer) POST(path string, handlerFunc api.HandlerFunc) {
	r.group.POST(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *Registerer) PUT(path string, handlerFunc api.HandlerFunc) {
	r.group.PUT(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *Registerer) DELETE(path string, handlerFunc api.HandlerFunc) {
	r.group.DELETE(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
