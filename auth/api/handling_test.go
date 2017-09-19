package api

import (
	"errors"
	"testing"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"fmt"
)

func TestController_Run_ShouldRunInnerController(t *testing.T) {
	tc := testController{}
	rq := testRequest{}
	rs := testResponse{}

	c := controller{rq, rs, &tc}

	tc.On("Run").Return(nil)

	c.Run()

	tc.AssertExpectations(t)
}

func TestController_Run_ShouldReturnError_IfInnerControllerRunReturnsError(t *testing.T) {
	tc := testController{}
	rq := testRequest{}
	rs := testResponse{}

	c := controller{rq, rs, &tc}

	const errorMsg = "SampleError"
	tc.On("Run").Return(errors.New(errorMsg))

	err := c.Run()

	assert.EqualError(t, err, errorMsg)
}

func TestControllerHandler_Handle_ShouldCallControllerRun(t *testing.T) {
	cfm := controllerFactoryMock{}

	tc := &testController{}
	rq := testRequest{}
	rs := testResponse{}

	tc.On("Run").Return(nil)

	cfm.On("CreateController", rq, rs).Return(tc, nil, nil)

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.Nil(t, err)
}

func TestControllerHandler_Handle_ShouldReturnError_IfControllerRunReturnsError(t *testing.T) {
	const someError = "some error";

	cfm := controllerFactoryMock{}

	tc := &testController{}
	rq := testRequest{}
	rs := testResponse{}

	tc.On("Run").Return(errors.New(someError))

	cfm.On("CreateController", rq, rs).Return(tc, nil, nil)

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.EqualError(t, err, someError)
}

func TestControllerHandler_Handle_ShouldReturnError_IfControllerFactiryReturnsError(t *testing.T) {
	const someError = "some error";

	cfm := controllerFactoryMock{}

	rq := testRequest{}
	rs := testResponse{}

	cfm.On("CreateController", rq, rs).Return(nil, nil, errors.New(someError))

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.EqualError(t, err, someError)
}

type testController struct {
	mock.Mock
}

func (t *testController) Run() error {
	args := t.Called()
	return args.Error(0)
}

type testRequest struct {
	mock.Mock
}

func (t testRequest) Param(name string) string {
	args := t.Called(name)
	return args.String(0)
}
func (t testRequest) Bind(i interface{}) error {
	args := t.Called(i)
	return args.Error(0)
}

type testResponse struct {
	mock.Mock
}

func (t testResponse) String(code int, s string) error {
	args := t.Called(code, s)
	return args.Error(0)
}
func (t testResponse) JSON(code int, i interface{}) error {
	args := t.Called(code, i)
	return args.Error(0)
}
func (t testResponse) NoContent(code int) error {
	args := t.Called(code)
	return args.Error(0)
}

type handleRequestMock struct {
	mock.Mock
}

func (am *handleRequestMock) HandleRequest() {
	am.Called()
}

type controllerFactoryMock struct {
	mock.Mock
}

func (cfm *controllerFactoryMock) CreateController(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	args := cfm.Called(rq, rs)
	return cfm.controller(args.Get(0)), cfm.controllerDestroyFunc(args.Get(1)), args.Error(2)
}

func (cfm *controllerFactoryMock) controller(i interface{}) business.Controller {
	var c business.Controller
	var ok bool
	if i == nil {
		return nil
	}
	if c, ok = i.(business.Controller); !ok {
		panic(fmt.Sprintf("assert: arguments: Controller failed because object wasn't correct type: %v", i))
	}
	return c
}

func (cfm *controllerFactoryMock) controllerDestroyFunc(i interface{}) api.ControllerDestroyFunc {
	var c api.ControllerDestroyFunc
	var ok bool
	if i == nil {
		return nil
	}
	if c, ok = i.(api.ControllerDestroyFunc); !ok {
		panic(fmt.Sprintf("assert: arguments: ControllerDestroyFunc failed because object wasn't correct type: %v", i))
	}
	return c
}