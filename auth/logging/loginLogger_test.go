package logging

import (
	"testing"

	"github.com/open-gtd/server/auth/domain"
)

func TestLogin_Logged(t *testing.T) {
	l := &loggerMock{}

	auth := domain.Auth{
		UserName:     "userName",
		SecurityCode: "sec",
	}

	l.On("Infof", "User '%v' has been logged.", domain.Name("userName"))
	l.On("Debugf", "Auth:", auth)

	logger = l
	sut := NewLogin()
	sut.Logged(auth)

	l.AssertExpectations(t)
}
