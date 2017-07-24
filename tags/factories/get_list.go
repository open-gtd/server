package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func NewGetList(c echo.Context) business.GetListController {
	return controllers.NewGetList(
		c,
		business.NewGetList(
			presenters.NewGetList(c),
			dao.NewGetList(),
		),
	)
}
