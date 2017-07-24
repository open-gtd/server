package api

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/tags/presentation/controllers"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/tags", GetList)
	g.GET("/tags/:"+controllers.NameQueryParam, Get)
	g.POST("/tags", Create)
	g.PUT("/tags/:"+controllers.NameQueryParam, Update)
	g.DELETE("/tags/:"+controllers.NameQueryParam, Delete)
}

func GetList(c echo.Context) error {
	return factories.NewGetList(c).Run()
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
