package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
	"github.com/open-gtd/server/tags/storage/mongo"
)

func NewCreate(rq api.Request, rs api.Response) (business.Controller, error) {
	conn, err := mongo.CreateDao("localhost")
	if err != nil {
		return nil, err
	}

	return controllers.NewCreate(
		rq,
		business.NewCreate(
			presenters.NewCreate(rs),
			dao.NewCreate(conn),
		),
	), nil
}
