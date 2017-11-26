package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type convertStrategyMock struct {
	mock.Mock
}

func (c *convertStrategyMock) Convert(tag domain.Tag) (domain.Tag, error) {
	args := c.Called(tag)
	return args.Get(0).(domain.Tag), args.Error(1)
}

