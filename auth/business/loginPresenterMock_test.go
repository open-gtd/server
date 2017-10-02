package business

import (
	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/mock"
)

type loginPresenterMock struct {
	mock.Mock
}

func (lp *loginPresenterMock) ShowAuthFailed() error {
	args := lp.Called()
	return args.Error(0)
}

func (lp *loginPresenterMock) Show(cert domain.Cert) error {
	args := lp.Called(cert.Token)
	return args.Error(0)
}
