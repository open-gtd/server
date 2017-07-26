package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type deleteTag struct {
	request    api.Request
	interactor business.Delete
}

func NewDelete(rq api.Request, i business.Delete) business.DeleteController {
	return deleteTag{request: rq, interactor: i}
}

func (d deleteTag) Run() error {
	name := d.request.Param(NameQueryParam)
	return d.interactor.Run(domain.Name(name))
}
