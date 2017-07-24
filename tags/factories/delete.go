package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
	"github.com/open-gtd/server/tags/storage/mongo"
)

func NewDelete(c echo.Context) (business.DeleteController, error) {
	conn, err := mongo.CreateDao("localhost")
	if err != nil {
		return nil, err
	}

	return controllers.NewDelete(
		c,
		business.NewDelete(
			presenters.NewDelete(c),
			dao.NewDelete(conn),
		),
	), nil
}
