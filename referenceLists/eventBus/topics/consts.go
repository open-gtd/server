package topics

import "github.com/open-gtd/server/eventBus"

const (
	Created  eventBus.Topic = "referenceListCreated"
	Modified eventBus.Topic = "referenceListModified"
	Deleted  eventBus.Topic = "referenceListDeleted"
)
