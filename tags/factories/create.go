package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/logging"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func Create(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	conn, err := Dao()
	if err != nil {
		return nil, nil, err
	}

	return controllers.NewCreate(
			rq,
			business.NewCreate(
				presenters.NewCreate(rs, GetBus()),
				dao.NewCreate(conn),
				logging.NewCreate(),
			),
		), func() error {
			return returnDao(conn)
		}, nil
}
