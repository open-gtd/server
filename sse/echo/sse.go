package echo

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/sse"
)

func NewSseRegisterer(g *echo.Group) sse.SseRegisterer {
	return &sseRegisterer{
		group:   g,
		channel: make(chan interface{}),
		exit:    make(chan bool),
	}
}

type sseRegisterer struct {
	group   *echo.Group
	channel chan interface{}
	exit    chan bool
}

func (sr sseRegisterer) CreatePushDataFunc(prefix sse.Prefix, closeNotify sse.ClientClosedNotificationFunc) sse.PushDataToSseFunc {
	sr.group.GET("/"+string(prefix), func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		notify := c.Response().CloseNotify()

		go func() {
			<-notify
			fmt.Println("notify")

			sr.exit <- true
			closeNotify()
		}()

		for {
			fmt.Println("new loop iteratinon")
			select {
			case <-sr.exit:
				fmt.Println("exit")
				return nil
			case v := <-sr.channel:
				if err := json.NewEncoder(c.Response()).Encode(v); err != nil {
					return err
				}
				c.Response().Flush()
			}
		}

		fmt.Println("exit")
		closeNotify()

		return nil
	})

	return func(topic sse.Topic, v interface{}) {
		sr.channel <- sse.Envelope{Topic: topic, Data: v}
	}
}
