package factories

import (
	"github.com/open-gtd/server/logging"
	loggingTags "github.com/open-gtd/server/tags/logging"
)

func GetLogger() logging.Logger {
	return loggingTags.Get()
}
