package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	typedErrors "github.com/open-gtd/server/api/validation/errors"
)

const controlerValidationError = "Controller validation error"
const someError = "Some error"
const controllerError = "Controller error"

func TestHandleRequest_ShouldRunControllerFromFactory(t *testing.T) {
	controller := testController{}
	fac := TestFactory{}

	request := testRequest{}
	response := testResponse{}

	controller.On("Run").Return(nil)
	fac.On("Create", request, response).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	HandleRequest(
		fac.Create,
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunction(t *testing.T) {
	controller := testController{}
	fac := TestFactory{}

	controller.On("Run").Return(nil)
	fac.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	request := testRequest{}
	response := testResponse{}
	HandleRequest(
		fac.Create,
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunctionEvenIfFactoryReturnsError(t *testing.T) {
	controller := testController{}
	fac := TestFactory{}

	fac.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(someError))
	fac.On("Destroy").Return(nil)

	request := testRequest{}
	response := testResponse{}
	HandleRequest(
		fac.Create,
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldReturnErrorIfFactoryReturnsError(t *testing.T) {
	fac := TestFactory{}

	fac.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Some error"))
	fac.On("Destroy").Return(nil)

	request := testRequest{}
	response := testResponse{}
	err := HandleRequest(
		fac.Create,
		request,
		response,
	)

	assert.EqualError(t, err, someError)
}

func TestHandleRequest_ShouldReturnErrorIfControllerRunReturnsError(t *testing.T) {
	controller := testController{}
	fac := TestFactory{}

	controller.On("Run").Return(errors.New(controllerError))
	fac.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	request := testRequest{}
	response := testResponse{}
	err := HandleRequest(
		fac.Create,
		request,
		response,
	)

	assert.EqualError(t, err, controllerError)
}

func TestHandleRequest_ShouldProduceBadRequestResponseErrorIfControllerRunReturnsError(t *testing.T) {
	controller := testController{}
	fac := TestFactory{}

	controller.On("Run").Return(testValidationError{})
	fac.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	request := testRequest{}
	response := testResponse{}
	response.On("JSON", http.StatusBadRequest, mock.Anything).
		Run(func(args mock.Arguments) {
		assert.Equal(t, args.Get(1).(MessageResponse).Message, controlerValidationError)
	}).Return(nil)

	HandleRequest(
		fac.Create,
		request,
		response,
	)
}

type testController struct {
	mock.Mock
}

func (t testController) Run() error {
	args := t.Called()
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

type TestFactory struct {
	mock.Mock
}

func (t TestFactory) Create(rq Request, rs Response) (Controller, ControllerDestroyFunc, error) {
	args := t.Called(rq, rs)
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


type testValidationError struct {}

func (t testValidationError) Type() typedErrors.ErrorType {
	return typedErrors.Validation
}

func (t testValidationError) Error() string {
	return controlerValidationError
}
