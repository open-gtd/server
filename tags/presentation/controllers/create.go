package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/api/validation"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
	"github.com/open-gtd/server/contract/tags"
)

type create struct {
	request    api.Request
	interactor business.Create
}

type Tag tags.CreatedTag

func NewCreate(rq api.Request, i business.Create) business.CreateController {
	return create{request: rq, interactor: i}
}

func (c create) Run() error {
	tag := &Tag{}
	if err := c.request.Bind(tag); err != nil {
		return err
	}

	if err := validation.AllowedValue("tag.Type", string(tag.Type), tags.AllowedTypes); err != nil {
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
