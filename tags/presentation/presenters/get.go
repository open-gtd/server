package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type get struct {
	c echo.Context
}

func NewGet(c echo.Context) business.GetPresenter {
	return get{c: c}
}

func (g get) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return g.c.JSON(http.StatusOK, tag)
}
