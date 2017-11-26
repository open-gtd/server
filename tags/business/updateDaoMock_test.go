package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type updateDaoMock struct {
	mock.Mock
}

func (u *updateDaoMock) Get(name domain.Name) (domain.Tag, error) {
	args := u.Called(name)
	return args.Get(0).(domain.Tag), args.Error(1)
}

func (u *updateDaoMock) Save(name domain.Name, tag domain.Tag) error {
	args := u.Called(name, tag)
	return args.Error(0)
}
