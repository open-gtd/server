package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type updatePresenterMock struct {
	mock.Mock
}

func (u *updatePresenterMock) Show(t domain.Tag) error {
	args := u.Called(t)
	return args.Error(0)
}

func (u *updatePresenterMock) ShowNotFound() error {
	args := u.Called()
	return args.Error(0)
}

