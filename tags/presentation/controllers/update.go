package controllers

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type update struct {
	c          echo.Context
	interactor business.Update
}

type tag struct {
	Name *presentation.Name           `json: "name"`
	Type *presentation.TypeDescriptor `json: "type"`
}

func NewUpdate(c echo.Context, i business.Update) business.UpdateController {
	return update{c: c, interactor: i}
}

func (u update) Run() error {
	originalName := u.c.QueryParam(NameQueryParam)
	ud := business.UpdateData{
		OriginalName: domain.Name(originalName),
	}

	t := tag{}
	u.c.Bind(&t)

	if t.Name != nil {
		name := domain.Name(*t.Name)
		ud.Name = &name
	}

	if t.Type != nil {
		typeEnum, err := converters.ConvertTypeToDomain(*t.Type)
		if err != nil {
			return err
		}
		ud.Type = &typeEnum
	}

	return u.interactor.Run(ud)
}
