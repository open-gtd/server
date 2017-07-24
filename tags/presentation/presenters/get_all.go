package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type getAll struct {
	c echo.Context
}

func NewGetAll(c echo.Context) business.GetAllPresenter {
	return getAll{c: c}
}

func (ga getAll) Show(t []domain.Tag) error {
	return ga.c.JSON(http.StatusOK, t)
}
