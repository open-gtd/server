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
	controller, err := factories.NewGetList(c)
	if err != nil {
		return err
	}

	return controller.Run()
}

func Get(c echo.Context) error {
	controller, err := factories.NewGet(c)
	if err != nil {
		return err
	}

	return controller.Run()
}

func Create(c echo.Context) error {
	controller, err := factories.NewCreate(c)
	if err != nil {
		return err
	}

	return controller.Run()
}

func Update(c echo.Context) error {
	controller, err := factories.NewUpdate(c)
	if err != nil {
		return err
	}

	return controller.Run()
}

func Delete(c echo.Context) error {
	controller, err := factories.NewDelete(c)
	if err != nil {
		return err
	}

	return controller.Run()
}
