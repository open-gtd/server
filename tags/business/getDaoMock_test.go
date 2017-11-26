package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type getDaoMock struct {
	mock.Mock
}

func (g *getDaoMock) Get(name domain.Name) (domain.Tag, error) {
	args := g.Called(name)
	return args.Get(0).(domain.Tag), args.Error(1)
}
