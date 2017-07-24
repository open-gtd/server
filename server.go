package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/open-gtd/server/api/projects"
	"github.com/open-gtd/server/api/reference"
	"github.com/open-gtd/server/api/tasks"
	tags "github.com/open-gtd/server/tags/api"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

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
