package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type create struct {
	c echo.Context
}

func NewCreate(c echo.Context) business.CreatePresenter {
	return create{c: c}
}

func (c create) Show(t domain.Tag) error {
	return c.c.JSON(http.StatusOK, t)
}
