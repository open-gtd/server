package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type update struct {
	response api.Response
}

func NewUpdate(rs api.Response) business.UpdatePresenter {
	return update{response: rs}
}

func (u update) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return u.response.JSON(http.StatusOK, tag)
}

func (u update) ShowNotFound() error {
	return u.response.NoContent(http.StatusNotFound)
}
