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

func NewControllerHandler(controllerFactory ControllerFactoryFunc) api.Handler {
	return controllerHandler{controllerFactory}
}

type controllerHandler struct {
	controllerFactory ControllerFactoryFunc
}

var handleRequest = api.HandleRequest

func (ch controllerHandler) Handle(rq api.Request, rs api.Response) error {
	c, cdf, err := ch.createController(rq, rs)
	if err != nil {
		return err
	}

	dc := api.DestroyableController{
		C:   c,
		Cdf: cdf,
	}

	return handleRequest(dc, rq, rs)
}

func (ch controllerHandler) createController(rq api.Request, rs api.Response) (api.Controller, api.ControllerDestroyFunc, error) {
	innerController, destroy, nil := ch.controllerFactory(rq, rs)

	c := &controller{
		rq:         rq,
		rs:         rs,
		controller: innerController,
	}
	return c, destroy, nil
}
