package presenters

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type create struct {
	c echo.Context
}

func NewCreate(c echo.Context) business.CreatePresenter {
	return create{c: c}
}

func (c create) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return c.c.JSON(http.StatusCreated, tag)
}
