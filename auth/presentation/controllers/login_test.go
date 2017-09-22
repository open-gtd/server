package controllers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/open-gtd/server/auth/business"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Run_ShouldCallRequestBindWithAuthInstance(t *testing.T) {
	rq := &testRequest{}
	l := &testLogin{}

	rq.On("Bind", mock.AnythingOfType("*presentation.LoginData")).Return(nil)
	l.On("Run", mock.Anything).Return(nil)

	sut := NewLogin(rq, l)

	sut.Run()
}

func TestLogin_Run_ShouldReturnError_IfRequestBindReturnsError(t *testing.T) {
	rq := &testRequest{}
	l := &testLogin{}

	const bindError = "bindError"
	rq.On("Bind", mock.Anything).Return(errors.New(bindError))
	l.On("Run", mock.Anything).Return(nil)

	sut := NewLogin(rq, l)

	err := sut.Run()

	assert.EqualError(t, err, bindError)
}

func TestLogin_Run_ShouldPassLoginDataFromRequestToConvert(t *testing.T) {
	rq := &testRequest{}
	l := &testLogin{}
	c := &testConvert{}
	convertFunc = c.Convert

	const userName = "userName"
	const securityCode = "sc"

	rq.On("Bind", mock.AnythingOfType("*presentation.LoginData")).
		Run(func(args mock.Arguments) {
			a := args.Get(0)
			ld, ok := a.(*presentation.LoginData)
			if !ok {
				t.Errorf("Bind Argument should be of type *presentation.LoginData. %v given.", a)
			}

			ld.UserName = userName
			ld.SecurityCode = securityCode
		}).
		Return(nil)

	c.On(
		"Convert",
		presentation.LoginData{
			UserName:     userName,
			SecurityCode: securityCode,
		},
	).Return(business.LoginData{})

	l.On("Run", mock.Anything).Return(nil)

	sut := NewLogin(rq, l)

	sut.Run()

	c.AssertExpectations(t)
}

func TestLogin_Run_ShouldPassLoginDataFromConvertToPresenterRun(t *testing.T) {
	rq := &testRequest{}
	l := &testLogin{}
	c := &testConvert{}
	convertFunc = c.Convert

	const userName = "userName"
	const securityCode = "sc"

	rq.On("Bind", mock.Anything).Return(nil)

	data := business.LoginData{
		Name:         userName,
		SecurityCode: securityCode,
	}
	c.On("Convert", mock.Anything).Return(data)

	l.On("Run", data).Return(nil)

	sut := NewLogin(rq, l)

	sut.Run()

	l.AssertExpectations(t)
}

func TestLogin_Run_ShouldReturnError_IfPresenterRunReturnsError(t *testing.T) {
	rq := &testRequest{}
	l := &testLogin{}
	c := &testConvert{}
	convertFunc = c.Convert

	const presenterError = "presenterError"

	rq.On("Bind", mock.Anything).Return(nil)

	c.On("Convert", mock.Anything).Return(business.LoginData{})

	l.On("Run", mock.Anything).Return(errors.New(presenterError))

	sut := NewLogin(rq, l)

	err := sut.Run()

	assert.EqualError(t, err, presenterError)
}

type testLogin struct {
	mock.Mock
}

func (l *testLogin) Run(ld business.LoginData) error {
	args := l.Called(ld)
	return args.Error(0)
}

type testRequest struct {
	mock.Mock
}

func (t testRequest) Param(name string) string {
	args := t.Called(name)
	return args.String(0)
}
func (t testRequest) Bind(i interface{}) error {
	args := t.Called(i)
	return args.Error(0)
}

type testConvert struct {
	mock.Mock
}

func (l *testConvert) Convert(ld presentation.LoginData) business.LoginData {
	args := l.Called(ld)
	return l.loginData(args.Get(0))
}
func (l *testConvert) loginData(i interface{}) business.LoginData {
	var c business.LoginData
	var ok bool
	if c, ok = i.(business.LoginData); !ok {
		panic(fmt.Sprintf("assert: arguments: business.LoginData failed because object wasn't correct type: %v", i))
	}
	return c
}
