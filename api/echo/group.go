package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

func NewGroupRestRegisterer(g *echo.Group) api.RestRegisterer {
	return &groupRegisterer{group: g}
}

type groupRegisterer struct {
	group *echo.Group
}

func (r *groupRegisterer) GET(path string, handlerFunc api.HandlerFunc) {
	r.group.GET(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *groupRegisterer) POST(path string, handlerFunc api.HandlerFunc) {
	r.group.POST(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *groupRegisterer) PUT(path string, handlerFunc api.HandlerFunc) {
	r.group.PUT(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *groupRegisterer) DELETE(path string, handlerFunc api.HandlerFunc) {
	r.group.DELETE(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
