package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type update struct {
	request    api.Request
	interactor business.Update
}

type tag struct {
	Name *presentation.Name           `json:"name"`
	Type *presentation.TypeDescriptor `json:"type"`
}

func NewUpdate(rq api.Request, i business.Update) business.UpdateController {
	return update{request: rq, interactor: i}
}

func (u update) Run() error {
	originalName := u.request.Param(NameQueryParam)
	ud := business.UpdateData{
		OriginalName: domain.Name(originalName),
	}

	t := tag{}
	u.request.Bind(&t)

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
