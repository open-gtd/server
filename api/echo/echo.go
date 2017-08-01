package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

func NewEchoRegisterer(g *echo.Echo) api.RestRegisterer {
	return &registerer{echo: g}
}

type registerer struct {
	echo *echo.Echo
}

func (r *registerer) GET(path string, handlerFunc api.HandlerFunc) {
	r.echo.GET(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) POST(path string, handlerFunc api.HandlerFunc) {
	r.echo.POST(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) PUT(path string, handlerFunc api.HandlerFunc) {
	r.echo.PUT(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) DELETE(path string, handlerFunc api.HandlerFunc) {
	r.echo.DELETE(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
