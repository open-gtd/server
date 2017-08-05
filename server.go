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
	"github.com/open-gtd/server/config"
	_ "github.com/open-gtd/server/config"
	"github.com/open-gtd/server/config/viper"
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

func logErrorDetails(l echo.Logger) echo.MiddlewareFunc {
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
	e.Logger.SetLevel(log.DEBUG)
	logging.SetLogger(e.Logger)

	conf, err := viper.New().
		FileName("server.conf").
		AddConfigPath("/etc/open-gtd").
		AddConfigPath("$HOME/.open-gtd").
		AddConfigPath(".").
		Build()

	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	c := eBus.NewCollection()

	auth.Initialize(apiEcho.NewEchoRegisterer(e), c, e.Logger)
	apiRegisterer, sseRegisterer := initializeAPI(e)

	initializeModules(apiRegisterer, sseRegisterer, c, conf)

	e.Logger.Fatal(e.Start(":" + conf.GetString("port")))
}

func initializeModules(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	reader config.Reader) {

	mngr := modules.NewModuleManager(apiRegisterer, sseRegisterer, busCollection, reader)

	mngr.Register(tags.Module())
	mngr.Register(referenceLists.Module())
	mngr.Register(tasks.Module())
	mngr.Register(projects.Module())
}

func initializeAPI(e *echo.Echo) (api.Registerer, sse.Registerer) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.Gzip(),
		logErrorDetails(e.Logger),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome in Open GTD!")
	})

	jwt := middleware.JWT([]byte("secret"))
	apiGroup := e.Group("/api", jwt)
	sseGroup := e.Group("/events", jwt)

	return apiEcho.NewGroupRegisterer(apiGroup), sseEcho.NewSseRegisterer(sseGroup)
}
