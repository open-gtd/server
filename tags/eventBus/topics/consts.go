package topics

import "github.com/open-gtd/server/eventBus"

const (
	Created  eventBus.Topic = "tagCreated"
	Modified eventBus.Topic = "tagModified"
	Deleted  eventBus.Topic = "tagDeleted"
)
