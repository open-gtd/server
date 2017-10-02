package api

import "github.com/stretchr/testify/mock"

type responseMock struct {
	mock.Mock
}

func (t responseMock) String(code int, s string) error {
	args := t.Called(code, s)
	return args.Error(0)
}
func (t responseMock) JSON(code int, i interface{}) error {
	args := t.Called(code, i)
	return args.Error(0)
}
func (t responseMock) NoContent(code int) error {
	args := t.Called(code)
	return args.Error(0)
}
