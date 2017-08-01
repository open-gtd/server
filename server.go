package main

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/open-gtd/server/api"
	apiEcho "github.com/open-gtd/server/api/echo"
	"github.com/open-gtd/server/auth"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/eventBus/eBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/projects"
	"github.com/open-gtd/server/referenceLists"
	"github.com/open-gtd/server/sse"
	sseEcho "github.com/open-gtd/server/sse/echo"
	"github.com/open-gtd/server/tags"
	"github.com/open-gtd/server/tasks"
)

func LogErrorDetails(l echo.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				l.Error(reflect.TypeOf(err), err)
				return err
			}

			return nil
		}
	}
}

func main() {
	e := echo.New()
	c := eBus.NewCollection()

	e.Logger.SetLevel(log.DEBUG)

	auth.Initialize(apiEcho.NewEchoRegisterer(e), c, e.Logger)
	initializeModules(c, e)

	e.Logger.Fatal(e.Start(":1323"))
}

func initializeModules(c eventBus.BusCollection, e *echo.Echo) {
	registeredModules := getModules()

	initializeBus(registeredModules, c)
	initializeApi(registeredModules, e)
	initializeLogger(registeredModules, e.Logger)
}

func getModules() []modules.Module {
	result := []modules.Module{}

	result = append(result, tags.Module())
	result = append(result, referenceLists.Module())
	result = append(result, tasks.Module())
	result = append(result, projects.Module())

	return result
}

func initializeApi(modules []modules.Module, e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.Gzip(),
		LogErrorDetails(e.Logger),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome in Open GTD!")
	})

	jwt := middleware.JWT([]byte("secret"))
	apiGroup := e.Group("/api", jwt)
	sseGroup := e.Group("/events", jwt)

	initializeRestModules(modules, apiEcho.NewGroupRestRegisterer(apiGroup))
	initializeSseModules(modules, sseEcho.NewSseRegisterer(sseGroup))
}

func initializeRestModules(modules []modules.Module, registerer api.RestRegisterer) {
	for _, module := range modules {
		module.RegisterHandlers(registerer)
	}
}

func initializeSseModules(modules []modules.Module, sr sse.SseRegisterer) {
	for _, module := range modules {
		module.RegisterSse(sr)
	}
}

func initializeBus(modules []modules.Module, c eventBus.BusCollection) {
	for _, module := range modules {
		module.RegisterBus(c)
	}
}

func initializeLogger(modules []modules.Module, l logging.Logger) {
	for _, module := range modules {
		module.RegisterLogger(l)
	}
}
