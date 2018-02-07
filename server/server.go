package server

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
	"github.com/open-gtd/server/sse"
	sseEcho "github.com/open-gtd/server/sse/echo"
	"github.com/open-gtd/server/tags"
	"github.com/spf13/viper"
)

var (
	e *echo.Echo
)

func Run() {
	e = echo.New()

	e.Logger.SetLevel(log.DEBUG)
	logging.SetLogger(e.Logger)

	eventBus.SetBus(eBus.NewBus())

	viper.SetConfigName("server.conf")
	viper.AddConfigPath("/etc/open-gtd")
	viper.AddConfigPath("$HOME/.open-gtd")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		e.Logger.Fatal("Configuration file not found!")
		return
	}

	e.Logger.Info("Used config file: " + viper.ConfigFileUsed())
	config.SetReader(viper.GetViper())

	initializeAPI(e)

	auth.Initialize()
	tags.Initialize()

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}

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

	groups := map[string]apiEcho.Router{}

	groups["/api"] = e.Group("/api", jwt)
	sseGroup := e.Group("/events", jwt)
	groups["/events"] = sseGroup

	api.SetRegisterer(apiEcho.NewEchoRegisterer(e, groups))
	sse.SetRegisterer(sseEcho.NewSseRegisterer(sseGroup))
}
