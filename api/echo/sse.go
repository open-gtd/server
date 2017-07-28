package echo

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/api"
)

type PushDataToSseFunc func(v interface{})

func NewSseRegisterer(e *echo.Echo, prefix string, m ...echo.MiddlewareFunc) api.RestRegisterer {
	return &restRegisterer{group: e.Group(prefix, m...)}
}

type sseRegisterer struct {
	group   *echo.Group
	channel chan interface{}
}

func (sr sseRegisterer) CreatePushDataFunc() PushDataToSseFunc {
	sr.group.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		for {
			v := <-sr.channel
			if err := json.NewEncoder(c.Response()).Encode(v); err != nil {
				return err
			}
			c.Response().Flush()
		}
		return nil
	})

	return func(v interface{}) {
		sr.channel <- v
	}
}
