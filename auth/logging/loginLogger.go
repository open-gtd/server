package logging

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/logging"
)

var logger = logging.GetLogger()

type login struct {
}

func NewLogin() business.LoginLogger {
	return login{}
}

func (c login) Logged(auth domain.Auth) {
	logger.Infof("User '%v' has been logged.", auth.UserName)
	logger.Debugf("Auth:", auth)
}
