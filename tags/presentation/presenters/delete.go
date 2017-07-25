package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
)

type delete struct {
	c echo.Context
}

func NewDelete(c echo.Context) business.DeletePresenter {
	return delete{c: c}
}

func (c delete) ShowSucced() error {
	return c.c.NoContent(http.StatusOK)
}

func (c delete) ShowNotFound() error {
	return c.c.NoContent(http.StatusNotFound)
}
