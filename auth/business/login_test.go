package business

import (
	"testing"

	"errors"

	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Run_ShouldAuthorizeUserUsingDao(t *testing.T) {
	lp := &testLoginPresenter{}
	ld := &testLoginDao{}
	ll := &testLoginLogger{}
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
	lp := &testLoginPresenter{}
	ld := &testLoginDao{}
	ll := &testLoginLogger{}
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
	lp := &testLoginPresenter{}
	ld := &testLoginDao{}
	ll := &testLoginLogger{}
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
	lp := &testLoginPresenter{}
	ld := &testLoginDao{}
	ll := &testLoginLogger{}
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
	lp := &testLoginPresenter{}
	ld := &testLoginDao{}
	ll := &testLoginLogger{}
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

type testLoginPresenter struct {
	mock.Mock
}

func (lp *testLoginPresenter) ShowAuthFailed() error {
	args := lp.Called()
	return args.Error(0)
}

func (lp *testLoginPresenter) Show(cert domain.Cert) error {
	args := lp.Called(cert.Token)
	return args.Error(0)
}

type testLoginDao struct {
	mock.Mock
}

func (ld *testLoginDao) Authorize(auth domain.Auth) error {
	args := ld.Called(auth.UserName, auth.SecurityCode)
	return args.Error(0)
}

type testLoginLogger struct {
	mock.Mock
}

func (ll *testLoginLogger) Logged(auth domain.Auth) {
	ll.Called(auth.UserName, auth.SecurityCode)
}
