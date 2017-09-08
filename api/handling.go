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

//DestroyableController wrapper to handle controller destruction after usage
type DestroyableController struct {
	C   Controller
	Cdf ControllerDestroyFunc
}

func (d DestroyableController) Destroy() error {
	if d.Cdf != nil {
		return d.Cdf()
	}
	return nil
}

func (d DestroyableController) Run() error {
	return d.C.Run()
}

//HandleRequest handle API request using Controller produced by ControllerFactory
//and sets response status to BadRequest with error message, if controller returns ValidationError
func HandleRequest(controller DestroyableController, rq Request, rs Response) error {
	defer controller.Destroy()

	err := controller.Run()
	if errors.IsValidationError(err) {
		return rs.JSON(http.StatusBadRequest, Message(err))
	}

	return err
}
