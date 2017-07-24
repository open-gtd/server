package reference

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/reference", GetAll)
	g.GET("/reference/:name", Get)
	g.POST("/reference", Create)
	g.PUT("/reference/:name", Update)
	g.DELETE("/reference/:name", Delete)
}

func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "ReferenceLists!")
}

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "ReferenceList!")
}

func Create(c echo.Context) error {
	return c.String(http.StatusOK, "ReferenceList created!")
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "ReferenceList updated!")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "ReferenceList deleted!")
}
