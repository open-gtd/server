package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type update struct {
	c echo.Context
}

func NewUpdate(c echo.Context) business.UpdatePresenter {
	return update{c: c}
}

func (u update) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return u.c.JSON(http.StatusOK, tag)
}

func (u update) ShowNotFound() error {
	return u.c.NoContent(http.StatusNotFound)
}
