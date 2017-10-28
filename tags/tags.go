package tags

import (
	"github.com/open-gtd/server/tags/api"
	"github.com/open-gtd/server/tags/sse"
)

var apiInitializer = api.Initialize
var sseInitializer = sse.Initialize

func Initialize() {
	apiInitializer()
	sseInitializer()
}
