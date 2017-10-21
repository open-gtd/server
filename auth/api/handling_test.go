package api

import (
	"errors"
	"testing"

	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/assert"
)

func TestController_Run_ShouldRunInnerController(t *testing.T) {
	tc := controllerMock{}
	rq := requestMock{}
	rs := responseMock{}

	c := controller{rq, rs, &tc}

	tc.On("Run").Return(nil)

	c.Run()

	tc.AssertExpectations(t)
}

func TestController_Run_ShouldReturnError_IfInnerControllerRunReturnsError(t *testing.T) {
	tc := controllerMock{}
	rq := requestMock{}
	rs := responseMock{}

	c := controller{rq, rs, &tc}

	const errorMsg = "SampleError"
	tc.On("Run").Return(errors.New(errorMsg))

	err := c.Run()

	assert.EqualError(t, err, errorMsg)
}

func TestControllerHandler_Handle_ShouldCallControllerRun(t *testing.T) {
	cfm := controllerFactoryMock{}

	tc := &controllerMock{}
	rq := requestMock{}
	rs := responseMock{}

	tc.On("Run").Return(nil)

	var controllerDestroyFunc api.ControllerDestroyFunc
	cfm.On("CreateController", rq, rs).Return(tc, controllerDestroyFunc, nil)

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.Nil(t, err)
}

func TestControllerHandler_Handle_ShouldReturnError_IfControllerRunReturnsError(t *testing.T) {
	const someError = "some error"

	cfm := controllerFactoryMock{}

	tc := &controllerMock{}
	rq := requestMock{}
	rs := responseMock{}

	tc.On("Run").Return(errors.New(someError))

	var controllerDestroyFunc api.ControllerDestroyFunc
	cfm.On("CreateController", rq, rs).Return(tc, controllerDestroyFunc, nil)

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.EqualError(t, err, someError)
}

func TestControllerHandler_Handle_ShouldReturnError_IfControllerFactiryReturnsError(t *testing.T) {
	const someError = "some error"

	cfm := controllerFactoryMock{}

	tc := &controllerMock{}
	rq := requestMock{}
	rs := responseMock{}

	//var controller business.Controller
	var controllerDestroyFunc api.ControllerDestroyFunc
	cfm.On("CreateController", rq, rs).Return(tc, controllerDestroyFunc, errors.New(someError))

	ch := NewControllerHandler(cfm.CreateController)

	err := ch.Handle(rq, rs)

	assert.EqualError(t, err, someError)
}
