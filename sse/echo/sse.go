package echo

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/sse"
)

type Router interface {
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
}

func NewSseRegisterer(r Router) sse.Registerer {
	return &sseRegisterer{
		router:  r,
		channel: make(chan interface{}),
		exit:    make(chan bool),
	}
}

type sseRegisterer struct {
	router  Router
	channel chan interface{}
	exit    chan bool
}

func (sr sseRegisterer) CreatePushDataFunc(prefix sse.Prefix, closeNotify sse.ClientClosedNotificationFunc) sse.PushDataToSseFunc {
	sr.router.Add(echo.GET, "/"+string(prefix), func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		notify := c.Response().CloseNotify()

		go func() {
			<-notify
			sr.exit <- true
			closeNotify()
		}()

		for {
			select {
			case <-sr.exit:
				closeNotify()
				return nil
			case v := <-sr.channel:
				if err := json.NewEncoder(c.Response()).Encode(v); err != nil {
					closeNotify()
					return err
				}
				c.Response().Flush()
			}
		}
	})

	return func(topic sse.Topic, v interface{}) {
		sr.channel <- sse.Envelope{Topic: topic, Data: v}
	}
}
