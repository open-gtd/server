package topics

import "github.com/open-gtd/server/eventBus"

const (
	Created  eventBus.Topic = "taskCreated"
	Modified eventBus.Topic = "taskModified"
	Deleted  eventBus.Topic = "taskDeleted"
)
