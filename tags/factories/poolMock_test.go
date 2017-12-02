package factories

import "github.com/stretchr/testify/mock"

type poolMock struct {
	mock.Mock
}

func (p *poolMock) BorrowObject() (interface{}, error) {
	args := p.Called()
	return args.Get(0), args.Error(1)
}

func (p *poolMock) ReturnObject(object interface{}) error {
	args := p.Called(object)
	return args.Error(0)
}

