package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type update struct {
	c          echo.Context
	interactor business.Update
}

func NewUpdate(c echo.Context, i business.Update) business.UpdateController {
	return update{c: c, interactor: i}
}

func (u update) Run() error {
	return u.interactor.Run()
}
