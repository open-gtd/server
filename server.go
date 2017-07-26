package main

import (
	"net/http"

	"fmt"
	"os"
	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	api_echo "github.com/open-gtd/server/api/echo"
	"github.com/open-gtd/server/api/projects"
	"github.com/open-gtd/server/api/reference"
	"github.com/open-gtd/server/api/tasks"
	"github.com/open-gtd/server/tags/api"
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
	registerer := api_echo.NewRegisterer(e, "/api") //, jwt)
	projects.RegisterHandlers(registerer)
	reference.RegisterHandlers(registerer)
	api.RegisterHandlers(registerer)
	tasks.RegisterHandlers(registerer)

	e.Logger.Fatal(e.Start(":1323"))
}
