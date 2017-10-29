package dao

import (
	"testing"

	"errors"

	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate_Get_ShouldCallDaoGetWithParameters(t *testing.T) {
	const tagName = "some tag name"

	daoMock := &daoMock{}
	var tag storage.Tag
	daoMock.On("Get", tagName).Return(tag, nil)

	sut := NewUpdate(daoMock)
	sut.Get(domain.Name(tagName))

	daoMock.AssertExpectations(t)
}

func TestUpdate_Get_ShouldReturnError_IfDaoGetWillReturnError(t *testing.T) {
	const tagName = "some tag name"
	const someError = "some error"

	daoMock := &daoMock{}
	var tag storage.Tag
	daoMock.On("Get", mock.Anything).Return(tag, errors.New(someError))

	sut := NewUpdate(daoMock)
	_, err := sut.Get(domain.Name(tagName))

	assert.EqualError(t, err, someError)
}

func TestUpdate_Get_ShouldReturnStorageNotFoundError_IfDaoGetWillReturnNotFoundError(t *testing.T) {
	const tagName = "some tag name"
	const daoNotFoundError = "not found"
	const storageNotFoundError = "not found"

	daoMock := &daoMock{}
	var tag storage.Tag
	daoMock.On("Get", mock.Anything).Return(tag, errors.New(daoNotFoundError))

	sut := NewUpdate(daoMock)
	_, err := sut.Get(domain.Name(tagName))

	assert.EqualError(t, err, storageNotFoundError)
}

func TestUpdateSave(t *testing.T) {

}
