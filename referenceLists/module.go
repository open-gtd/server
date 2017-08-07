package referenceLists

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/modules"
	apiReferenceLists "github.com/open-gtd/server/referenceLists/api"
	"github.com/open-gtd/server/sse"
	//sseReferenceLists "github.com/open-gtd/server/referenceLists/sse"
	"github.com/open-gtd/server/config"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) Register(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	reader config.Reader) {

	apiReferenceLists.RegisterHandlers(apiRegisterer)
	//eventBusReferenceLists.RegisterBusHandlers(busCollection)
	//sseReferenceLists.RegisterSse(sseRegisterer)
}
