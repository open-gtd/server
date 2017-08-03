package logging

import "github.com/open-gtd/server/logging"

var logger logging.Logger = logging.NullLogger{}

func RegisterLogger(l logging.Logger) {
	logger = l
}
