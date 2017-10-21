package dao

import (
	"github.com/open-gtd/server/tags/storage"
	"github.com/stretchr/testify/mock"
)

type daoMock struct {
	mock.Mock
}

func (d *daoMock) Insert(tag storage.Tag) error {
	args := d.Called(tag)
	return args.Error(0)
}

func (d *daoMock) Update(name string, tag storage.Tag) error {
	args := d.Called(name, tag)
	return args.Error(0)
}

func (d *daoMock) Delete(name string) error {
	args := d.Called(name)
	return args.Error(0)
}

func (d *daoMock) GetList() ([]storage.Tag, error) {
	args := d.Called()
	return args.Get(0).([]storage.Tag), args.Error(1)
}

func (d *daoMock) Get(name string) (storage.Tag, error) {
	args := d.Called(name)
	return args.Get(0).(storage.Tag), args.Error(1)
}
