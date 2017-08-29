package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestHandleRequest_ShouldRunControllerFromFactory(t *testing.T) {
	controller := TestController{}
	fac := TestFactory{}

	controller.On("Run").Once().Return(nil, nil)
	fac.On("Create").Return(controller, nil)

	HandleRequest(
		func(Request, Response) (Controller, ControllerDestroyFunc, error) {
			return controller, nil, nil
		},
		TestRequest{},
		TestResponse{},
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunction(t *testing.T) {
	controller := TestController{}
	fac := TestFactory{}

	controller.On("Run").Once().Return(nil, nil)
	fac.On("Create").Return(controller, nil)
	fac.On("Destroy").Once().Return(nil)

	HandleRequest(
		fac.Create,
		TestRequest{},
		TestResponse{},
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunctionEvenIfFactoryReturnsError(t *testing.T) {
	controller := TestController{}
	fac := TestFactory{}

	fac.On("Create").Return(nil, errors.New("Some error"))
	fac.On("Destroy").Once().Return(nil)

	HandleRequest(
		fac.Create,
		TestRequest{},
		TestResponse{},
	)

	controller.AssertExpectations(t)
}

type TestController struct {
	mock.Mock
}

func (t TestController) Run() error {
	args := t.Called()
	return args.Error(0)
}

type TestResponse struct {
	mock.Mock
}

func (t TestResponse) String(code int, s string) error {
	args := t.Called(code, s)
	return args.Error(0)
}
func (t TestResponse) JSON(code int, i interface{}) error {
	args := t.Called(code, i)
	return args.Error(0)
}
func (t TestResponse) NoContent(code int) error {
	args := t.Called(code)
	return args.Error(0)
}

type TestRequest struct {
	mock.Mock
}

func (t TestRequest) Param(name string) string {
	args := t.Called(name)
	return args.String(0)
}
func (t TestRequest) Bind(i interface{}) error {
	args := t.Called(i)
	return args.Error(0)
}

type TestFactory struct {
	mock.Mock
}

func (t TestFactory) Create(Request, Response) (Controller, ControllerDestroyFunc, error) {
	args := t.Called()
	return controller(args.Get(0)), t.Destroy, args.Error(1)
}

func controller(obj interface{}) Controller {
	var c Controller
	var ok bool
	if obj == nil {
		return nil
	}
	if c, ok = obj.(Controller); !ok {
		panic(fmt.Sprintf("assert: arguments: Controller failed because object wasn't correct type: %v", obj))
	}
	return c
}

func (t TestFactory) Destroy() error {
	args := t.Called()
	return args.Error(0)
}
