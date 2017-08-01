package domain

import (
	"time"

	"github.com/open-gtd/server/common"
	"github.com/open-gtd/server/tags/presentation"
)

type (
	Name            string
	StateDescriptor common.StateDescriptor
	Tags            []presentation.Tag
	Notes           string
	Focus           bool
	DueDate         time.Time
	ScheduledDate   time.Time
	TimeRequired    uint8
	EnergyLevel     uint8
)

const (
	InboxState     StateDescriptor = StateDescriptor(common.InboxState)
	NextState      StateDescriptor = StateDescriptor(common.NextState)
	WaitingState   StateDescriptor = StateDescriptor(common.WaitingState)
	ScheduledState StateDescriptor = StateDescriptor(common.ScheduledState)
	SomedayState   StateDescriptor = StateDescriptor(common.SomedayState)
	LogbookState   StateDescriptor = StateDescriptor(common.LogbookState)
)

const (
	EnergyLevelNone = EnergyLevel(iota)
	EnergyLevelLow
	EnergyLevelMedium
	EnergyLevelHigh
)

const (
	TimeRequiredNone = TimeRequired(iota)
	Minutes5
	Minutes10
	Minutes15
	Minutes30
	Minutes45
	Hour
	Hours2
	Hours3
	Hours4
	Hours6
	Hours8
	EvenMore
)

type TaskDescriptor struct {
	Name Name `json: "name"`
}

type Task struct {
	TaskDescriptor
	State         StateDescriptor `json: "state"`
	Tags          Tags            `json: "tags"`
	Notes         Notes           `json: "notes"`
	DueDate       DueDate         `json: "due_date"`
	ScheduledDate ScheduledDate   `json: "schedule_date"`
	TimeRequired  TimeRequired    `json: "time_required"`
	EnergyLevel   EnergyLevel     `json: "energy_level"`
}
