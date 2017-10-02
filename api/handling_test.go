package api

import (
	"errors"
	"testing"

	"net/http"

	typedErrors "github.com/open-gtd/server/api/validation/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const controlerValidationError = "Controller validation error"
const controllerError = "controller error"

func TestHandleRequest_ShouldRunControllerFromFactory(t *testing.T) {
	controller := controllerMock{}
	fac := factoryMock{}

	request := requestMock{}
	response := responseMock{}

	controller.On("Run").Return(nil)
	fac.On("Destroy").Return(nil)

	HandleRequest(
		DestroyableController{controller, fac.Destroy},
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunction(t *testing.T) {
	controller := controllerMock{}
	fac := factoryMock{}

	controller.On("Run").Return(nil)
	fac.On("Destroy").Return(nil)

	request := requestMock{}
	response := responseMock{}
	HandleRequest(
		DestroyableController{controller, fac.Destroy},
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldNotFail_IfDestroyFunctionIsNil(t *testing.T) {
	controller := controllerMock{}

	controller.On("Run").Return(nil)

	request := requestMock{}
	response := responseMock{}
	HandleRequest(
		DestroyableController{controller, nil},
		request,
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldReturnErrorIfControllerRunReturnsError(t *testing.T) {
	controller := controllerMock{}
	fac := factoryMock{}

	controller.On("Run").Return(errors.New(controllerError))
	fac.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	request := requestMock{}
	response := responseMock{}
	err := HandleRequest(
		DestroyableController{controller, fac.Destroy},
		request,
		response,
	)

	assert.EqualError(t, err, controllerError)
}

func TestHandleRequest_ShouldProduceBadRequestResponseErrorIfControllerRunReturnsError(t *testing.T) {
	controller := controllerMock{}
	fac := factoryMock{}

	controller.On("Run").Return(testValidationError{})
	fac.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	fac.On("Destroy").Return(nil)

	request := requestMock{}
	response := responseMock{}
	response.On("JSON", http.StatusBadRequest, mock.Anything).
		Run(func(args mock.Arguments) {
			assert.Equal(t, args.Get(1).(MessageResponse).Message, controlerValidationError)
		}).Return(nil)

	HandleRequest(
		DestroyableController{controller, fac.Destroy},
		request,
		response,
	)
}

type testValidationError struct{}

func (t testValidationError) Type() typedErrors.ErrorType {
	return typedErrors.Validation
}

func (t testValidationError) Error() string {
	return controlerValidationError
}
