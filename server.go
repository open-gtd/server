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
	_ "github.com/open-gtd/server/config"
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
	apiRegisterer, sseRegisterer := initializeApi(e)

	initializeModules(apiRegisterer, sseRegisterer, c, e.Logger)

	e.Logger.Fatal(e.Start(":1323"))
}

func initializeModules(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	logger logging.Logger) {

	mngr := modules.NewModuleManager(apiRegisterer, sseRegisterer, busCollection, logger)

	mngr.Register(tags.Module())
	mngr.Register(referenceLists.Module())
	mngr.Register(tasks.Module())
	mngr.Register(projects.Module())
}

func initializeApi(e *echo.Echo) (api.Registerer, sse.Registerer) {
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

	return apiEcho.NewGroupRestRegisterer(apiGroup), sseEcho.NewSseRegisterer(sseGroup)
}
