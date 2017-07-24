package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type update struct {
	c echo.Context
}

func NewUpdate(c echo.Context) business.UpdatePresenter {
	return update{c: c}
}

func (u update) Show(t domain.Tag) error {
	return u.c.JSON(http.StatusOK, t)
}
