package api

import (
	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

type handlerMock struct {
	mock.Mock
}

func (h *handlerMock) Handle(rq api.Request, rs api.Response) error {
	args := h.Called(rq, rs)
	return args.Error(0)
}
