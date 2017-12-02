package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/stretchr/testify/mock"
)

type controllerFactoryMock struct {
	mock.Mock
}

func (cfm *controllerFactoryMock) CreateController(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyer, error) {
	args := cfm.Called(rq, rs)
	return args.Get(0).(business.Controller), args.Get(1).(api.ControllerDestroyer), args.Error(2)
}
