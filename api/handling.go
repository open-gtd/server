package api

import (
	"net/http"

	"github.com/open-gtd/server/api/validation/errors"
)

//Controller represents ability to call business logic from request handling method
type Controller interface {
	Run() error
}

//ControllerDestroyFunc function to be called when controller is no longer needed during request
type ControllerDestroyFunc func() error

//ControllerFactoryFunc function to be called when controller is no longer needed during request
type ControllerFactoryFunc func(Request, Response) (Controller, ControllerDestroyFunc, error)

//HandleRequest handle API request using Controller produced by ControllerFactory
//and sets response status to BadRequest with error message, if controller returns ValidationError
func HandleRequest(controllerFactory ControllerFactoryFunc, rq Request, rs Response) error {
	controller, destroy, err := controllerFactory(rq, rs)
	if destroy != nil {
		defer destroy()
	}

	if err != nil {
		return err
	}

	err = controller.Run()
	if errors.IsValidationError(err) {
		return rs.JSON(http.StatusBadRequest, Message(err))
	}

	return err
}
