package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/presentation/controllers"
	"github.com/open-gtd/server/auth/presentation/presenters"
	"github.com/open-gtd/server/auth/storage/dao"
	"github.com/open-gtd/server/auth/logging"
)

func Login(rq api.Request, rs api.Response) (business.Controller, api.ControllerDestroyFunc, error) {
	//conn, err := Dao()
	//if err != nil {
	//	return nil, nil, err
	//}

	return controllers.NewLogin(
			rq,
			business.NewLogin(
				presenters.NewLogin(rs, GetBus()),
				dao.NewLogin(), //conn),
				logging.NewLogin(),
			),
		), func() error {
			return nil //returnDao(conn)
		}, nil
}
