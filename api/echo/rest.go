package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

func NewRestRegisterer(g *echo.Group) api.RestRegisterer {
	return &restRegisterer{group: g}
}

type restRegisterer struct {
	group *echo.Group
}

func (r *restRegisterer) GET(path string, handlerFunc api.HandlerFunc) {
	r.group.GET(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *restRegisterer) POST(path string, handlerFunc api.HandlerFunc) {
	r.group.POST(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *restRegisterer) PUT(path string, handlerFunc api.HandlerFunc) {
	r.group.PUT(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *restRegisterer) DELETE(path string, handlerFunc api.HandlerFunc) {
	r.group.DELETE(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
