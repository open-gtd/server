package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type getList struct {
	c          echo.Context
	interactor business.GetList
}

func NewGetList(c echo.Context, i business.GetList) business.GetListController {
	return getList{c: c, interactor: i}
}

func (ga getList) Run() error {
	return ga.interactor.Run()
}
