package common

type StateDescriptor string

const (
	InboxState     StateDescriptor = "inbox"
	NextState      StateDescriptor = "next"
	WaitingState   StateDescriptor = "waiting"
	ScheduledState StateDescriptor = "scheduled"
	SomedayState   StateDescriptor = "someday"
	LogbookState   StateDescriptor = "logbook"
)
