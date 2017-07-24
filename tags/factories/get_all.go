package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
)

func NewGetAll(c echo.Context) business.GetAllController {
	return controllers.NewGetAll(
		c,
		business.NewGetAll(
			presenters.NewGetAll(c),
		),
	)
}
