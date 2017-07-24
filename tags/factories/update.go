package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
	"github.com/open-gtd/server/tags/storage/mongo"
)

func NewUpdate(c echo.Context) (business.UpdateController, error) {
	conn, err := mongo.CreateDao("localhost")
	if err != nil {
		return nil, err
	}

	return controllers.NewUpdate(
		c,
		business.NewUpdate(
			presenters.NewUpdate(c),
			dao.NewUpdate(conn),
		),
	), nil
}
