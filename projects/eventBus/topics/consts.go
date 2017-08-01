package topics

import "github.com/open-gtd/server/eventBus"

const (
	Created  eventBus.Topic = "projectCreated"
	Modified eventBus.Topic = "projectModified"
	Deleted  eventBus.Topic = "projectDeleted"
)
