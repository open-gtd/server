package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type createDaoMock struct {
	mock.Mock
}

func (cd *createDaoMock) Insert(t domain.Tag) error {
	args := cd.Called(t)
	return args.Error(0)
}
