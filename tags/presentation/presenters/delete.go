package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/contract/tags"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/eventBus/topics"
)

type deleteTag struct {
	response api.Response
	bus      eventBus.Bus
}

func NewDelete(rs api.Response, bus eventBus.Bus) business.DeletePresenter {
	return deleteTag{response: rs, bus: bus}
}

func (d deleteTag) ShowSucceed(name domain.Name) error {
	d.bus.Publish(topics.Deleted, tags.DeletedTag{Name: string(name)})

	return d.response.NoContent(http.StatusOK)
}

func (d deleteTag) ShowNotFound() error {
	return d.response.NoContent(http.StatusNotFound)
}
