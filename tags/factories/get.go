package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func NewGet(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	conn, err := CreateDao()
	if err != nil {
		return nil, nil, err
	}

	return controllers.NewGet(
		rq,
		business.NewGet(
			presenters.NewGet(rs),
			dao.NewGet(conn),
		),
	), func() error {
		return DestroyDao(conn)
	}, nil
}
