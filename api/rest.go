package api

import (
	"net/http"

	"github.com/open-gtd/server/api/validation/errors"
)

//MessageResponse represents JSON message response body
type MessageResponse struct {
	Message string `json:"message"`
}

//Message returns new message response body, based on error message
func Message(err error) MessageResponse {
	return MessageResponse{Message: err.Error()}
}

//Response represents ability to send response to user
type Response interface {
	String(code int, s string) error
	JSON(code int, i interface{}) error
	NoContent(code int) error
}

//Request represents ability to read request data
type Request interface {
	Param(name string) string
	Bind(i interface{}) error
}

//HandlerFunc function to handle request
type HandlerFunc func(Request, Response) error

//Registerer represents ability to register methods in http API
type Registerer interface {
	GET(path string, handlerFunc HandlerFunc)
	POST(path string, handlerFunc HandlerFunc)
	PUT(path string, handlerFunc HandlerFunc)
	DELETE(path string, handlerFunc HandlerFunc)
}

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
