package tasks

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/tasks", GetAll)
	g.GET("/tasks/:name", Get)
	g.POST("/tasks", Create)
	g.PUT("/tasks/:name", Update)
	g.DELETE("/tasks/:name", Delete)
}

func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Tasks!")
}

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "Task!")
}

func Create(c echo.Context) error {
	return c.String(http.StatusOK, "Task created!")
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "Task updated!")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Task deleted!")
}
