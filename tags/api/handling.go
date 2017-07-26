package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
)

type ControllerFactoryFunc func(api.Request, api.Response) (business.Controller, error)

type controller struct {
	rq                api.Request
	rs                api.Response
	controllerFactory ControllerFactoryFunc
}

func (ac controller) Run() error {
	c, err := ac.controllerFactory(ac.rq, ac.rs)
	if err != nil {
		return err
	}

	return c.Run()
}

func createController(controllerFactory ControllerFactoryFunc) api.ControllerFactoryFunc {
	return func(rq api.Request, rs api.Response) (api.Controller, error) {
		return controller{rq: rq, rs: rs, controllerFactory: controllerFactory}, nil
	}
}

func handler(controllerFactory ControllerFactoryFunc) api.HandlerFunc {
	return func(rq api.Request, rs api.Response) error {
		return api.HandleRequest(createController(controllerFactory), rq, rs)
	}
}
