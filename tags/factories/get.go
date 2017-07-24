package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
	"github.com/open-gtd/server/tags/storage/dao"
)

func NewGet(c echo.Context) business.GetController {
	return controllers.NewGet(
		c,
		business.NewGet(
			presenters.NewGet(c),
			dao.NewGet(),
		),
	)
}
