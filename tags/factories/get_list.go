package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
	"github.com/open-gtd/server/tags/storage/mongo"
)

func NewGetList(c echo.Context) (business.GetListController, error) {
	conn, err := mongo.CreateDao("localhost")
	if err != nil {
		return nil, err
	}

	return controllers.NewGetList(
		c,
		business.NewGetList(
			presenters.NewGetList(c),
			dao.NewGetList(conn),
		),
	), nil
}
