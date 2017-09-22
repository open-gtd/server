package presenters

import (
	"net/http"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/eventBus/topics"
	"github.com/open-gtd/server/auth/presentation/converters"
	"github.com/open-gtd/server/eventBus"
)

type login struct {
	response api.Response
	bus      eventBus.Bus
}

func NewLogin(rs api.Response, bus eventBus.Bus) business.LoginPresenter {
	return login{response: rs, bus: bus}
}

var convertFunc = converters.ConvertToPresentation

func (c login) Show(cert domain.Cert) error {
	auth := convertFunc(cert)

	c.bus.Publish(topics.LoggedIn, cert)
	return c.response.JSON(http.StatusOK, auth)
}

func (c login) ShowAuthFailed() error {
	return c.response.NoContent(http.StatusForbidden)
}
