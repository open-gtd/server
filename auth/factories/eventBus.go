package factories

import (
	authBus "github.com/open-gtd/server/auth/eventBus"
	"github.com/open-gtd/server/eventBus"
)

func GetBus() eventBus.Bus {
	return authBus.Get()
}
