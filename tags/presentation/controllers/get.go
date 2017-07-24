package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type get struct {
	c          echo.Context
	interactor business.Get
}

func NewGet(c echo.Context, i business.Get) business.GetController {
	return get{c: c, interactor: i}
}

func (g get) Run() error {
	return g.interactor.Run()
}
