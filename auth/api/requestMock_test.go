package api

import "github.com/stretchr/testify/mock"

type requestMock struct {
	mock.Mock
}

func (t requestMock) Param(name string) string {
	args := t.Called(name)
	return args.String(0)
}

func (t requestMock) Bind(i interface{}) error {
	args := t.Called(i)
	return args.Error(0)
}
