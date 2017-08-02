package api

import (
	"net/http"

	"github.com/open-gtd/server/api/validation/errors"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func Message(err error) MessageResponse {
	return MessageResponse{Message: err.Error()}
}

type Response interface {
	String(code int, s string) error
	JSON(code int, i interface{}) error
	NoContent(code int) error
}

type Request interface {
	Param(name string) string
	Bind(i interface{}) error
}

type HandlerFunc func(Request, Response) error

type RestRegisterer interface {
	GET(path string, handlerFunc HandlerFunc)
	POST(path string, handlerFunc HandlerFunc)
	PUT(path string, handlerFunc HandlerFunc)
	DELETE(path string, handlerFunc HandlerFunc)
}

func HandleRequest(controllerFactory ControllerFactoryFunc, rq Request, rs Response) error {
	controller, destroy, err := controllerFactory(rq, rs)
	if destroy != nil {
		defer destroy()
	}

	if err != nil {
		return err
	}

	err = controller.Run()
	if errors.IsNotAllowedValueError(err) {
		return rs.JSON(http.StatusBadRequest, Message(err))
	}

	return err
}

type ControllerDestroyFunc func() error
type ControllerFactoryFunc func(Request, Response) (Controller, ControllerDestroyFunc, error)

type Controller interface {
	Run() error
}
