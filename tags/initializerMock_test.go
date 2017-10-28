package tags

import "github.com/stretchr/testify/mock"

type initializerMock struct {
	mock.Mock
}

func (i *initializerMock) apiInitialize() {
	i.Called()
}

func (i *initializerMock) sseInitialize() {
	i.Called()
}
