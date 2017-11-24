package business

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/mock"
)

type createPresenterMock struct {
	mock.Mock
}

func (cp *createPresenterMock) Show(t domain.Tag) error {
	args := cp.Called(t)
	return args.Error(0)
}

func (cp *createPresenterMock) ShowNotUnique() error {
	args := cp.Called()
	return args.Error(0)
}
