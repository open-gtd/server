package business

import (
	"testing"

	"errors"

	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Run_ShouldAuthorizeUserUsingDao(t *testing.T) {
	lp := &loginPresenterMock{}
	ld := &loginDaoMock{}
	ll := &loginLoggerMock{}
	login := NewLogin(lp, ld, ll)

	const name = "name"
	const code = "code"

	ld.On("Authorize", domain.Name(name), domain.SecurityCode(code)).
		Return(nil)
	ll.On("Logged", mock.Anything, mock.Anything).
		Return(nil)
	lp.On("Show", mock.Anything).
		Return(nil)

	login.Run(LoginData{Name: name, SecurityCode: code})

	ld.AssertExpectations(t)
}

func TestLogin_Run_ShouldPresentLoginFailed_IfDaoAuthorizeReturnsError(t *testing.T) {
	lp := &loginPresenterMock{}
	ld := &loginDaoMock{}
	ll := &loginLoggerMock{}
	login := NewLogin(lp, ld, ll)

	const someError = "someError"

	ld.On("Authorize", mock.Anything, mock.Anything).
		Return(errors.New(someError))
	ll.On("Logged", mock.Anything, mock.Anything).
		Return(nil)
	lp.On("ShowAuthFailed").
		Return(nil)

	const name = "name"
	const code = "code"
	login.Run(LoginData{Name: name, SecurityCode: code})

	lp.AssertExpectations(t)
}

func TestLogin_Run_ShouldReturnError_IfPresenterShowAuthFailedReturnsError(t *testing.T) {
	lp := &loginPresenterMock{}
	ld := &loginDaoMock{}
	ll := &loginLoggerMock{}
	login := NewLogin(lp, ld, ll)

	const someError = "someError"
	const someOtherError = "someOtherError"

	ld.On("Authorize", mock.Anything, mock.Anything).
		Return(errors.New(someOtherError))
	ll.On("Logged", mock.Anything, mock.Anything).
		Return(nil)
	lp.On("ShowAuthFailed", mock.Anything).
		Return(errors.New(someError))

	const name = "name"
	const code = "code"
	err := login.Run(LoginData{Name: name, SecurityCode: code})

	assert.EqualError(t, err, someError)
}

func TestLogin_Run_ShouldLogLogged(t *testing.T) {
	lp := &loginPresenterMock{}
	ld := &loginDaoMock{}
	ll := &loginLoggerMock{}
	login := NewLogin(lp, ld, ll)

	const name = "name"
	const code = "code"

	ld.On("Authorize", mock.Anything, mock.Anything).
		Return(nil)
	ll.On("Logged", domain.Name(name), domain.SecurityCode(code)).
		Return(nil)
	lp.On("Show", mock.Anything).
		Return(nil)

	login.Run(LoginData{Name: name, SecurityCode: code})

	ll.AssertExpectations(t)
}

func TestLogin_Run_ShouldPresentToken(t *testing.T) {
	lp := &loginPresenterMock{}
	ld := &loginDaoMock{}
	ll := &loginLoggerMock{}
	login := NewLogin(lp, ld, ll)

	const name = "name"
	const code = "code"

	ld.On("Authorize", mock.Anything, mock.Anything).
		Return(nil)
	ll.On("Logged", mock.Anything, mock.Anything).
		Return(nil)
	lp.On("Show", mock.AnythingOfType("domain.Token")).
		Return(nil)

	login.Run(LoginData{Name: name, SecurityCode: code})

	lp.AssertExpectations(t)
}
