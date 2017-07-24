package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type getList struct {
	c echo.Context
}

func NewGetList(c echo.Context) business.GetListPresenter {
	return getList{c: c}
}

func (gl getList) Show(t []domain.Tag) error {
	tags, err := converters.ConvertAllToPresentation(t)
	if err != nil {
		return err
	}
	return gl.c.JSON(http.StatusOK, tags)
}
