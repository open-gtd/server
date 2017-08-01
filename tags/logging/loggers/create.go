package loggers

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type create struct {
	logger logging.Logger
}

func NewCreate(logger logging.Logger) business.CreateLogger {
	return create{logger: logger}
}

func (c create) TagCreated(tag domain.Tag) {
	c.logger.Info("Tag has been created.")
	c.logger.Debugf("Tag:", tag)
}
