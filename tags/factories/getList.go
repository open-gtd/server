package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func GetList(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyer, error) {
	conn, err := getDao()
	if err != nil {
		return nil, nil, err
	}

	return controllers.NewGetList(
		rq,
		business.NewGetList(
			presenters.NewGetList(rs),
			dao.NewGetList(conn),
		),
	), connectionDestroyer{conn:conn}, nil
}
