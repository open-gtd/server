package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type get struct {
	c          echo.Context
	interactor business.Get
}

func NewGet(c echo.Context, i business.Get) business.GetController {
	return get{c: c, interactor: i}
}

func (g get) Run() error {
	name := g.c.QueryParam(NameQueryParam)
	return g.interactor.Run(domain.Name(name))
}
