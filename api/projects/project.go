package projects

import (
	"github.com/open-gtd/server/common"
	"github.com/open-gtd/server/tags/domain"
)

type (
	Name            string
	StateDescriptor common.StateDescriptor
	Tags            []domain.Tag
	Notes           string
	Focus           bool
)

const (
	InboxState     StateDescriptor = StateDescriptor(common.InboxState)
	NextState      StateDescriptor = StateDescriptor(common.NextState)
	ScheduledState StateDescriptor = StateDescriptor(common.ScheduledState)
	SomedayState   StateDescriptor = StateDescriptor(common.SomedayState)
	LogbookState   StateDescriptor = StateDescriptor(common.LogbookState)
)

type ProjectDescriptor struct {
	Name Name `json: "name"`
}

type Project struct {
	ProjectDescriptor
	State StateDescriptor `json: "state"`
	Tags  Tags            `json: "tags"`
	Notes Notes           `json: "notes"`
}
