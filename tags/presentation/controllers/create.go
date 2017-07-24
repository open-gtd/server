package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type create struct {
	c          echo.Context
	interactor business.Create
}

type Tag presentation.Tag

func NewCreate(c echo.Context, i business.Create) business.CreateController {
	return create{c: c, interactor: i}
}

func (c create) Run() error {
	tag := &Tag{}
	if err := c.c.Bind(tag); err != nil {
		return err
	}

	dType, err := converters.ConvertTypeToDomain(tag.Type)
	if err != nil {
		return err
	}

	return c.interactor.Run(business.CreateData{
		Name: domain.Name(tag.Name),
		Type: dType,
	})
}
