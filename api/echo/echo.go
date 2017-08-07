package echo

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

//NewEchoRegisterer returns API registerer for Echo Framework which allows to register endpoints at main level
func NewEchoRegisterer(e *echo.Echo, groups map[string]*echo.Group) api.Registerer {
	return &registerer{echo: e, groups: groups}
}

type registerer struct {
	echo   *echo.Echo
	groups map[string]*echo.Group
}

func (r *registerer) GET(prefix string, path string, handlerFunc api.HandlerFunc) {

	if group, ok := r.groups[prefix]; ok {
		group.GET(path, func(c echo.Context) error {
			return handlerFunc(c, c)
		})
	}

	r.echo.GET(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) POST(prefix string, path string, handlerFunc api.HandlerFunc) {
	if group, ok := r.groups[prefix]; ok {
		group.POST(path, func(c echo.Context) error {
			return handlerFunc(c, c)
		})
	}

	r.echo.POST(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) PUT(prefix string, path string, handlerFunc api.HandlerFunc) {
	if group, ok := r.groups[prefix]; ok {
		group.PUT(path, func(c echo.Context) error {
			return handlerFunc(c, c)
		})
	}

	r.echo.PUT(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
func (r *registerer) DELETE(prefix string, path string, handlerFunc api.HandlerFunc) {
	if group, ok := r.groups[prefix]; ok {
		group.DELETE(path, func(c echo.Context) error {
			return handlerFunc(c, c)
		})
	}

	r.echo.DELETE(path, func(c echo.Context) error {
		return handlerFunc(c, c)
	})
}
