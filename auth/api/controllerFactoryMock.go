package api

import (
	"fmt"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/stretchr/testify/mock"
)

type controllerFactoryMock struct {
	mock.Mock
}

func (cfm *controllerFactoryMock) CreateController(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	args := cfm.Called(rq, rs)
	return cfm.controller(args.Get(0)), cfm.controllerDestroyFunc(args.Get(1)), args.Error(2)
}

func (cfm *controllerFactoryMock) controller(i interface{}) business.Controller {
	var c business.Controller
	var ok bool
	if i == nil {
		return nil
	}
	if c, ok = i.(business.Controller); !ok {
		panic(fmt.Sprintf("assert: arguments: Controller failed because object wasn't correct type: %v", i))
	}
	return c
}

func (cfm *controllerFactoryMock) controllerDestroyFunc(i interface{}) api.ControllerDestroyFunc {
	var c api.ControllerDestroyFunc
	var ok bool
	if i == nil {
		return nil
	}
	if c, ok = i.(api.ControllerDestroyFunc); !ok {
		panic(fmt.Sprintf("assert: arguments: ControllerDestroyFunc failed because object wasn't correct type: %v", i))
	}
	return c
}
