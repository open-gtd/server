package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type create struct {
	response api.Response
}

func NewCreate(rs api.Response) business.CreatePresenter {
	return create{response: rs}
}

func (c create) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return c.response.JSON(http.StatusCreated, tag)
}
