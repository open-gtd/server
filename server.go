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
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/eventBus/eBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/projects"
	"github.com/open-gtd/server/referenceLists"
	"github.com/open-gtd/server/tags"
	"github.com/open-gtd/server/tasks"
	"github.com/spf13/viper"
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

var (
	e *echo.Echo
)

func init() {
	e = echo.New()

	e.Logger.SetLevel(log.DEBUG)
	logging.SetLogger(e.Logger)

	eventBus.SetBus(eBus.NewBus())

	viper.SetConfigFile("server.conf")
	viper.AddConfigPath("/etc/open-gtd")
	viper.AddConfigPath("$HOME/.open-gtd")
	viper.AddConfigPath(".")

	config.SetReader(viper.GetViper())

	initializeAPI(e)

	auth.Initialize()
	projects.Initialize()
	referenceLists.Initialize()
	tags.Initialize()
	tasks.Initialize()
}

func main() {
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}

func initializeAPI(e *echo.Echo) {
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

	groups := map[string]*echo.Group{}

	groups["/api"] = e.Group("/api", jwt)
	groups["/events"] = e.Group("/events", jwt)

	api.SetRegisterer(apiEcho.NewEchoRegisterer(e, groups))
}
