package projects

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/projects", GetAll)
	g.GET("/projects/:name", Get)
	g.POST("/projects", Create)
	g.PUT("/projects/:name", Update)
	g.DELETE("/projects/:name", Delete)
}

func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Projects!")
}

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "Project!")
}

func Create(c echo.Context) error {
	return c.String(http.StatusOK, "Project created!")
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "Project updated!")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Project deleted!")
}
