package logging

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

var logger = logging.GetLogger()

type create struct {
}

func NewCreate() business.CreateLogger {
	return create{}
}

func (c create) TagCreated(tag domain.Tag) {
	logger.Info("Tag has been created.")
	logger.Debugf("Tag:", tag)
}
