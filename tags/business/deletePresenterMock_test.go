package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type deletePresenterMock struct {
	mock.Mock
}

func (d *deletePresenterMock) ShowSucceed(name domain.Name) error {
	args := d.Called(name)
	return args.Error(0)
}

func (d *deletePresenterMock) ShowNotFound() error {
	args := d.Called()
	return args.Error(0)
}
