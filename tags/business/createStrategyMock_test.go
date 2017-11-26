package business

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/mock"
)

type createStrategyMock struct {
	mock.Mock
}

func (cs *createStrategyMock) Create(name domain.Name) (domain.Tag, error) {
	args := cs.Called(name)
	arg := args.Get(0)
	return arg.(domain.Tag), args.Error(1)
}
