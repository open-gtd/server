package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type deleteDaoMock struct {
	mock.Mock
}

func (d *deleteDaoMock) Delete(name domain.Name) error {
	args := d.Called(name)
	return args.Error(0)
}

