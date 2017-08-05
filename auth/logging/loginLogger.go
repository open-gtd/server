package logging

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/logging"
)

type login struct {
}

func NewLogin() business.LoginLogger {
	return login{}
}

func (c login) Logged(auth domain.Auth) {
	logging.Infof("User '%v' has been logged.", auth.UserName)
	logging.Debugf("Auth:", auth)
}
