package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/auth/domain"
)

type loginDaoMock struct {
	mock.Mock
}

func (ld *loginDaoMock) Authorize(auth domain.Auth) error {
	args := ld.Called(auth.UserName, auth.SecurityCode)
	return args.Error(0)
}
