package logging

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type create struct {
}

func NewCreate() business.CreateLogger {
	return create{}
}

func (c create) TagCreated(tag domain.Tag) {
	logging.Info("Tag has been created.")
	logging.Debugf("Tag:", tag)
}
