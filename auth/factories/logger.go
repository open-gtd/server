package factories

import (
	loggingAuth "github.com/open-gtd/server/auth/logging"
	"github.com/open-gtd/server/logging"
)

func GetLogger() logging.Logger {
	return loggingAuth.Get()
}
