package api

import "github.com/stretchr/testify/mock"

type controllerMock struct {
	mock.Mock
}

func (t controllerMock) Run() error {
	args := t.Called()
	return args.Error(0)
}
