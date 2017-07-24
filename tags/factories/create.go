package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
)

func NewCreate(c echo.Context) business.CreateController {
	return controllers.NewCreate(
		c,
		business.NewCreate(
			presenters.NewCreate(c),
		),
	)
}
