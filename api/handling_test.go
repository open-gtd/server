package api

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestHandleRequest_ShouldRunControllerFromFactory(t *testing.T) {
	controller := TestController{}

	controller.On("Run").Once()

	HandleRequest(
		func (Request, Response) (Controller, ControllerDestroyFunc, error) {
			return controller, nil, nil
		},
		TestRequest{},
		TestResponse{},
	)

	controller.AssertExpectations(t)
}

type TestController struct {
	mock.Mock

	e error
}

func (t TestController) Run() error {
	t.Called()
	return t.e
}

type TestResponse struct {
	mock.Mock

	e error
}

func (t TestResponse) String(code int, s string) error    { return t.e }
func (t TestResponse) JSON(code int, i interface{}) error { return t.e }
func (t TestResponse) NoContent(code int) error           { return t.e }

type TestRequest struct {
	mock.Mock

	name string
	e    error
}

func (t TestRequest) Param(name string) string { return t.name }
func (t TestRequest) Bind(i interface{}) error { return t.e }
