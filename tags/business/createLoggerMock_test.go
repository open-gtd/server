package business

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
)

type createLoggerMock struct {
	mock.Mock
}

func (cl *createLoggerMock) TagCreated(tag domain.Tag) {
	cl.Called(tag)
}
