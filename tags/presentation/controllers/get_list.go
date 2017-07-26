package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
)

type getList struct {
	request    api.Request
	interactor business.GetList
}

func NewGetList(r api.Request, i business.GetList) business.GetListController {
	return getList{request: r, interactor: i}
}

func (ga getList) Run() error {
	return ga.interactor.Run()
}
