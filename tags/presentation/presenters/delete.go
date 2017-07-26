package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
)

type deleteTag struct {
	response api.Response
}

func NewDelete(rs api.Response) business.DeletePresenter {
	return deleteTag{response: rs}
}

func (c deleteTag) ShowSucceed() error {
	return c.response.NoContent(http.StatusOK)
}

func (c deleteTag) ShowNotFound() error {
	return c.response.NoContent(http.StatusNotFound)
}
