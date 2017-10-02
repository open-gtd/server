package api

import "github.com/stretchr/testify/mock"

type factoryMock struct {
	mock.Mock
}

func (t factoryMock) Destroy() error {
	args := t.Called()
	return args.Error(0)
}
