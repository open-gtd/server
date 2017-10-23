package echo

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/api"
)

type handlerFuncMock struct {
	mock.Mock
}

func (h *handlerFuncMock) Handle(request api.Request, response api.Response) error {
	args := h.Called(request, response)
	return args.Error(0)
}