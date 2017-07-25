package main

import (
	"net/http"

	"fmt"
	"os"
	"reflect"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/open-gtd/server/api/projects"
	"github.com/open-gtd/server/api/reference"
	"github.com/open-gtd/server/api/tasks"
	tags "github.com/open-gtd/server/tags/api"
)

func LogErrorDetails() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				fmt.Fprint(os.Stderr, reflect.TypeOf(err))
				fmt.Fprint(os.Stderr, err)
			}

			return nil
		}
	}
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(LogErrorDetails())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome in Open GTD!")
	})

	//jwt := middleware.JWT([]byte("secret"))
	//apiGroup := e.Group("/api", jwt)
	apiGroup := e.Group("/api")
	projects.RegisterHandlers(apiGroup)
	reference.RegisterHandlers(apiGroup)
	tags.RegisterHandlers(apiGroup)
	tasks.RegisterHandlers(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
