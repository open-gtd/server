package logging

import "github.com/open-gtd/server/logging"

var logger = logging.NullLogger{}

func RegisterLogger(l logging.Logger) {
	logger = logger
}
