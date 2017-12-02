package api

import "github.com/stretchr/testify/mock"

type destroyerMock struct {
	mock.Mock
}

func (t destroyerMock) Destroy() error {
	args := t.Called()
	return args.Error(0)
}
