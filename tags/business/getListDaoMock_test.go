package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type getListDaoMock struct {
	mock.Mock
}

func (g *getListDaoMock) Get() ([]domain.Tag, error) {
	args := g.Called()
	return args.Get(0).([]domain.Tag), args.Error(1)
}

