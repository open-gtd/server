package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type getList struct {
	response api.Response
}

func NewGetList(rs api.Response) business.GetListPresenter {
	return getList{response: rs}
}

func (gl getList) Show(t []domain.Tag) error {
	tags, err := converters.ConvertAllToPresentation(t)
	if err != nil {
		return err
	}
	return gl.response.JSON(http.StatusOK, tags)
}
