package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type delete struct {
	c          echo.Context
	interactor business.Delete
}

func NewDelete(c echo.Context, i business.Delete) business.DeleteController {
	return delete{c: c, interactor: i}
}

func (d delete) Run() error {
	return d.interactor.Run()
}
