package tags

import (
	"github.com/open-gtd/server/tags/api"
	"github.com/open-gtd/server/tags/sse"
)

func Initialize() {
	api.Initialize()
	sse.Initialize()

}
