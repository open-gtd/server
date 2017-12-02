package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func Delete(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyer, error) {

	conn, err := getDao()
	if err != nil {
		return nil, nil, err
	}

	return controllers.NewDelete(
		rq,
		business.NewDelete(
			presenters.NewDelete(rs, GetBus()),
			dao.NewDelete(conn),
		),
	), connectionDestroyer{conn:conn}, nil
}
