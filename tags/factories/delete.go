package factories

import (
	"github.com/labstack/echo"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/presentation/controllers"
	"github.com/open-gtd/server/tags/presentation/presenters"
)

func NewDelete(c echo.Context) business.DeleteController {
	return controllers.NewDelete(
		c,
		business.NewDelete(
			presenters.NewDelete(c),
		),
	)
}
