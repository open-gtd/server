package api

import (
	"errors"
	"testing"

	"net/http"

	typedErrors "github.com/open-gtd/server/api/validation/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const controllerValidationError = "Controller validation error"
const controllerError = "controller error"

func TestHandleRequest_ShouldRunControllerFromFactory(t *testing.T) {
	controller := controllerMock{}
	destroyer := destroyerMock{}

	response := responseMock{}

	controller.On("Run").Return(nil)
	destroyer.On("Destroy").Return(nil)

	HandleRequest(
		DestroyableController{controller, destroyer},
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldCallDestroyFunction(t *testing.T) {
	controller := controllerMock{}
	destroyer := destroyerMock{}

	controller.On("Run").Return(nil)
	destroyer.On("Destroy").Return(nil)

	response := responseMock{}
	HandleRequest(
		DestroyableController{controller, destroyer},
		response,
	)

	controller.AssertExpectations(t)
}

func TestHandleRequest_ShouldReturnErrorIfControllerRunReturnsError(t *testing.T) {
	controller := controllerMock{}
	destroyer := destroyerMock{}

	controller.On("Run").Return(errors.New(controllerError))
	destroyer.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	destroyer.On("Destroy").Return(nil)

	response := responseMock{}
	err := HandleRequest(
		DestroyableController{controller, destroyer},
		response,
	)

	assert.EqualError(t, err, controllerError)
}

func TestHandleRequest_ShouldProduceBadRequestResponseErrorIfControllerRunReturnsError(t *testing.T) {
	controller := controllerMock{}
	destroyer := destroyerMock{}

	controller.On("Run").Return(testValidationError{})
	destroyer.On("Create", mock.Anything, mock.Anything).Return(controller, nil)
	destroyer.On("Destroy").Return(nil)

	response := responseMock{}
	response.On("JSON", http.StatusBadRequest, mock.Anything).
		Run(func(args mock.Arguments) {
		assert.Equal(t, args.Get(1).(MessageResponse).Message, controllerValidationError)
	}).Return(nil)

	HandleRequest(
		DestroyableController{controller, destroyer},
		response,
	)
}

type testValidationError struct{}

func (t testValidationError) Type() typedErrors.ErrorType {
	return typedErrors.Validation
}

func (t testValidationError) Error() string {
	return controllerValidationError
}
