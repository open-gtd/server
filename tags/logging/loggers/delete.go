package loggers

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
)

type delete struct {
	logger logging.Logger
}

func NewDelete(logger logging.Logger) business.DeleteLogger {
	return delete{logger: logger}
}