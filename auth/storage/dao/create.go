package dao

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/domain"
)

type login struct {
}

func NewLogin() business.LoginDao {
	return login{}
}

func (c login) Authorize(t domain.Auth) error {
	//TODO: use OAuth2
	return nil
}
