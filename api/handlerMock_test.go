package api

import "github.com/stretchr/testify/mock"

type handlerMock struct {
	mock.Mock
}

func (h *handlerMock) Handle(rq Request, rs Response) error {
	args := h.Called(rq, rs)
	return args.Error(0)
}
