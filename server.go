package main

import (
	"net/http"

	"fmt"
	"os"
	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/open-gtd/server/api"
	apiEcho "github.com/open-gtd/server/api/echo"
	projectsApi "github.com/open-gtd/server/api/projects"
	referenceApi "github.com/open-gtd/server/api/reference"
	tasksApi "github.com/open-gtd/server/api/tasks"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/eventBus/eBus"
	projectsBus "github.com/open-gtd/server/eventBus/projects"
	referenceBus "github.com/open-gtd/server/eventBus/reference"
	tasksBus "github.com/open-gtd/server/eventBus/tasks"
	"github.com/open-gtd/server/sse"
	sseEcho "github.com/open-gtd/server/sse/echo"
	tagsApi "github.com/open-gtd/server/tags/api"
	tagsBus "github.com/open-gtd/server/tags/eventBus"
	tagsSse "github.com/open-gtd/server/tags/sse"
)

func LogErrorDetails() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				fmt.Fprintln(os.Stderr, reflect.TypeOf(err))
				fmt.Fprintln(os.Stderr, err)
			}

			return nil
		}
	}
}

func main() {
	e := echo.New()
	c := eBus.NewCollection()

	initializeBus(c)
	initializeApi(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func initializeApi(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.Gzip(),
		LogErrorDetails(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome in Open GTD!")
	})

	//jwt := middleware.JWT([]byte("secret"))
	apiGroup := e.Group("/api") //, jwt)
	initializeRestModules(apiEcho.NewRestRegisterer(apiGroup))
	sseGroup := e.Group("/events") //, jwt)
	initializeSseModules(sseEcho.NewSseRegisterer(sseGroup))
}

func initializeRestModules(registerer api.RestRegisterer) {
	projectsApi.RegisterHandlers(registerer)
	referenceApi.RegisterHandlers(registerer)
	tagsApi.RegisterHandlers(registerer)
	tasksApi.RegisterHandlers(registerer)
}

func initializeSseModules(sr sse.SseRegisterer) {
	//projectsBus.RegisterSse(sr)
	//referenceBus.RegisterSse(sr)
	tagsSse.RegisterSse(sr)
	//tasksBus.RegisterSse(sr)
}

func initializeBus(c eventBus.BusCollection) {
	projectsBus.RegisterBus(c)
	referenceBus.RegisterBus(c)
	tagsBus.RegisterBus(c)
	tasksBus.RegisterBus(c)

	projectsBus.RegisterBusHandlers(c)
}
