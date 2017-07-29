package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation/converters"
	"github.com/open-gtd/server/tags/eventBus/topics"
)

type create struct {
	response api.Response
	bus      eventBus.Bus
}

func NewCreate(rs api.Response, bus eventBus.Bus) business.CreatePresenter {
	return create{response: rs, bus: bus}
}

func (c create) Show(t domain.Tag) error {
	tag, err := converters.ConvertToPresentation(t)
	if err != nil {
		return err
	}

	c.bus.Publish( topics.Created, tag)
	return c.response.JSON(http.StatusCreated, tag)
}
