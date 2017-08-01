package loggers

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
)

type update struct {
	logger logging.Logger
}

func NewUpdate(logger logging.Logger) business.UpdateLogger {
	return update{logger: logger}
}
