package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
)

type ControllerFactoryFunc func(api.Request, api.Response) (business.Controller, api.ControllerDestroyFunc, error)

type controller struct {
	rq         api.Request
	rs         api.Response
	controller business.Controller
}

func (ac *controller) Run() error {
	return ac.controller.Run()
}

func createController(controllerFactory ControllerFactoryFunc) api.ControllerFactoryFunc {
	return func(rq api.Request, rs api.Response) (api.Controller, api.ControllerDestroyFunc, error) {
		innerController, destroy, nil := controllerFactory(rq, rs)

		c := &controller{
			rq:         rq,
			rs:         rs,
			controller: innerController,
		}
		return c, destroy, nil
	}
}

func handler(controllerFactory ControllerFactoryFunc) api.HandlerFunc {
	return func(rq api.Request, rs api.Response) error {
		controller := createController(controllerFactory)
		return api.HandleRequest(controller, rq, rs)
	}
}
