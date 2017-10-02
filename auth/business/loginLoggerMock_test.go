package business

import (
	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/mock"
)

type loginLoggerMock struct {
	mock.Mock
}

func (ll *loginLoggerMock) Logged(auth domain.Auth) {
	ll.Called(auth.UserName, auth.SecurityCode)
}
