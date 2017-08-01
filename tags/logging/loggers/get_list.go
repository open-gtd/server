package loggers

import (
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/tags/business"
)

type getList struct {
	logger logging.Logger
}

func NewGetList(logger logging.Logger) business.GetListLogger {
	return getList{logger: logger}
}
