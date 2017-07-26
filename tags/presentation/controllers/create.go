package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type create struct {
	request    api.Request
	interactor business.Create
}

type Tag presentation.Tag

func NewCreate(rq api.Request, i business.Create) business.CreateController {
	return create{request: rq, interactor: i}
}

func (c create) Run() error {
	tag := &Tag{}
	if err := c.request.Bind(tag); err != nil {
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
