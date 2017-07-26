package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type get struct {
	response api.Response
}

func NewGet(rs api.Response) business.GetPresenter {
	return get{response: rs}
}

func (g get) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}
	return g.response.JSON(http.StatusOK, tag)
}

func (g get) ShowNotFound() error {
	return g.response.NoContent(http.StatusNotFound)
}
