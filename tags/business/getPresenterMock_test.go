package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type getPresenterMock struct {
	mock.Mock
}

func (g *getPresenterMock) Show(t domain.Tag) error {
	args := g.Called(t)
	return args.Error(0)
}

func (g *getPresenterMock) ShowNotFound() error {
	args := g.Called()
	return args.Error(0)
}



