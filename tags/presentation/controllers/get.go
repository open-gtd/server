package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type get struct {
	request    api.Request
	interactor business.Get
}

func NewGet(rq api.Request, i business.Get) business.GetController {
	return get{request: rq, interactor: i}
}

func (g get) Run() error {
	name := g.request.Param(NameQueryParam)
	return g.interactor.Run(domain.Name(name))
}
