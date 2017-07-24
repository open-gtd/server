package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type create struct {
	c          echo.Context
	interactor business.Create
}

func NewCreate(c echo.Context, i business.Create) business.CreateController {
	return create{c: c, interactor: i}
}

func (c create) Run() error {
	return c.interactor.Run()
}
