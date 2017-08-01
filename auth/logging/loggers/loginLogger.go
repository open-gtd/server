package loggers

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/logging"
)

type login struct {
	logger logging.Logger
}

func NewLogin(logger logging.Logger) business.LoginLogger {
	return login{logger: logger}
}

func (c login) Logged(auth domain.Auth) {
	c.logger.Info("User '%v' has been logged.", auth.UserName)
	c.logger.Debugf("Auth:", auth)
}
