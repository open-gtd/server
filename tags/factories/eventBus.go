package factories

import (
	"github.com/open-gtd/server/eventBus"
	tagsBus "github.com/open-gtd/server/tags/eventBus"
)

func GetBus() eventBus.Bus {
	return tagsBus.Get()
}
