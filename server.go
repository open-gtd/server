package main

import (
	"net/http"

	"fmt"
	"os"
	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	apiEcho "github.com/open-gtd/server/api/echo"
	projectsApi "github.com/open-gtd/server/api/projects"
	referenceApi "github.com/open-gtd/server/api/reference"
	tasksApi "github.com/open-gtd/server/api/tasks"
	"github.com/open-gtd/server/eventBus"
	projectsBus "github.com/open-gtd/server/eventBus/projects"
	referenceBus "github.com/open-gtd/server/eventBus/reference"
	tasksBus "github.com/open-gtd/server/eventBus/tasks"
	tagsApi "github.com/open-gtd/server/tags/api"
	tagsBus "github.com/open-gtd/server/tags/eventBus"
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
	c := eventBus.New()

	initializeRestApi(e)

	initializeBus(c)

	e.Logger.Fatal(e.Start(":1323"))
}

func initializeRestApi(e *echo.Echo) {
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
	registerer := apiEcho.NewRegisterer(e, "/api") //, jwt)
	projectsApi.RegisterHandlers(registerer)
	referenceApi.RegisterHandlers(registerer)
	tagsApi.RegisterHandlers(registerer)
	tasksApi.RegisterHandlers(registerer)
}

func initializeBus(c eventBus.BusCollection) {
	projectsBus.RegisterBus(c)
	referenceBus.RegisterBus(c)
	tagsBus.RegisterBus(c)
	tasksBus.RegisterBus(c)

	projectsBus.RegisterBusHandlers(c)
}
