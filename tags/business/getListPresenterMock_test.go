package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type getListPresenterMock struct {
	mock.Mock
}

func (g *getListPresenterMock) Show(t []domain.Tag) error {
	args := g.Called(t)
	return args.Error(0)
}



