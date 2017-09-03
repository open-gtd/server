package api

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestController_Run_ShouldRunInnerController(t *testing.T) {
	tc := testController{}
	rq := testRequest{}
	rs := testResponse{}

	c := controller{rq, rs,&tc }

	tc.On("Run").Return(nil)

	c.Run()

	tc.AssertExpectations(t)
}

func TestController_Run_ShouldReturnError_IfInnerControllerRunReturnsError(t *testing.T) {
	tc := testController{}
	rq := testRequest{}
	rs := testResponse{}

	c := controller{rq, rs,&tc }

	const errorMsg = "SampleError"
	tc.On("Run").Return(errors.New(errorMsg))

	err := c.Run()

	assert.EqualError(t, err, errorMsg)
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