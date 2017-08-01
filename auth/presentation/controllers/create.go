package controllers

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/open-gtd/server/auth/presentation/converters"
)

type login struct {
	request    api.Request
	interactor business.Login
}

func NewLogin(rq api.Request, i business.Login) business.LoginController {
	return login{request: rq, interactor: i}
}

func (c login) Run() error {
	auth := &presentation.LoginData{}
	if err := c.request.Bind(auth); err != nil {
		return err
	}

	dLoginData, err := converters.ConvertLoginDataToBusiness(*auth)
	if err != nil {
		return err
	}

	return c.interactor.Run(dLoginData)
}
