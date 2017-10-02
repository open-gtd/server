package controllers

import (
	"github.com/open-gtd/server/auth/business"
	"github.com/stretchr/testify/mock"
)

type loginMock struct {
	mock.Mock
}

func (l *loginMock) Run(ld business.LoginData) error {
	args := l.Called(ld)
	return args.Error(0)
}
