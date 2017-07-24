package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type getAll struct {
	c          echo.Context
	interactor business.GetAll
}

func NewGetAll(c echo.Context, i business.GetAll) business.GetAllController {
	return getAll{c: c, interactor: i}
}

func (ga getAll) Run() error {
	return ga.interactor.Run()
}
