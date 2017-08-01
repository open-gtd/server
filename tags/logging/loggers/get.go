package loggers

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
)

type get struct {
	logger logging.Logger
}

func NewGet(logger logging.Logger) business.GetLogger {
	return get{logger: logger}
}
