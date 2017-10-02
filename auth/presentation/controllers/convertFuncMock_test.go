package controllers

import (
	"fmt"

	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/stretchr/testify/mock"
)

type convertFuncMock struct {
	mock.Mock
}

func (l *convertFuncMock) Convert(ld presentation.LoginData) business.LoginData {
	args := l.Called(ld)
	return l.loginData(args.Get(0))
}
func (l *convertFuncMock) loginData(i interface{}) business.LoginData {
	var c business.LoginData
	var ok bool
	if c, ok = i.(business.LoginData); !ok {
		panic(fmt.Sprintf("assert: arguments: business.LoginData failed because object wasn't correct type: %v", i))
	}
	return c
}
