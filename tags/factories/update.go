package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
	"github.com/open-gtd/server/tags/logging/loggers"
)

func Update(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	conn, err := Dao()
	if err != nil {
		return nil, nil, err
	}

	return controllers.NewUpdate(
			rq,
			business.NewUpdate(
				presenters.NewUpdate(rs),
				dao.NewUpdate(conn),
				loggers.NewUpdate(GetLogger()),
			),
		), func() error {
			return returnDao(conn)
		}, nil
}
