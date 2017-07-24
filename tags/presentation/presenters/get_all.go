package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type getAll struct {
	c echo.Context
}

func NewGetAll(c echo.Context) business.GetAllPresenter {
	return getAll{c: c}
}

func (ga getAll) Show(t []domain.Tag) error {
	tags, err := converters.ConvertAllToPresentation(t)
	if err != nil {
		return err
	}
	return ga.c.JSON(http.StatusOK, tags)
}
