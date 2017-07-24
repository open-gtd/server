package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
)

func NewUpdate(c echo.Context) business.UpdateController {
	return controllers.NewUpdate(
		c,
		business.NewUpdate(
			presenters.NewUpdate(c),
		),
	)
}
