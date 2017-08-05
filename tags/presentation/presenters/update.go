package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/eventBus/topics"
	"github.com/open-gtd/server/tags/presentation/converters"
)

type update struct {
	response api.Response
	bus      eventBus.Bus
}

func NewUpdate(rs api.Response, bus eventBus.Bus) business.UpdatePresenter {
	return update{response: rs, bus: bus}
}

func (u update) Show(t domain.Tag) error {
	tag, err := converters.ConvertToModifiedTag(t)
	if err != nil {
		return err
	}

	u.bus.Publish(topics.Modified, tag)
	return u.response.JSON(http.StatusOK, tag)
}

func (u update) ShowNotFound() error {
	return u.response.NoContent(http.StatusNotFound)
}
