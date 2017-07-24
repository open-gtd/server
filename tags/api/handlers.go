package api

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/factories"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/tags", GetAll)
	g.GET("/tags/:name", Get)
	g.POST("/tags", Create)
	g.PUT("/tags/:name", Update)
	g.DELETE("/tags/:name", Delete)
}

func GetAll(c echo.Context) error {
	return factories.NewGetAll(c).Run()
}

func Get(c echo.Context) error {
	return factories.NewGet(c).Run()
}

func Create(c echo.Context) error {
	return factories.NewCreate(c).Run()
}

func Update(c echo.Context) error {
	return factories.NewUpdate(c).Run()
}

func Delete(c echo.Context) error {
	return factories.NewDelete(c).Run()
}
