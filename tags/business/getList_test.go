package business

import (
	"errors"
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetListRunShouldCallDaoGet(t *testing.T) {
	getListPresenterMock := &getListPresenterMock{}
	getListPresenterMock.On("Show", mock.Anything).Return(nil)

	getListDaoMock := &getListDaoMock{}
	getListDaoMock.On("Get").Return([]domain.Tag{}, nil)

	NewGetList(getListPresenterMock, getListDaoMock).Run()

	getListDaoMock.AssertExpectations(t)
}

func TestGetListRunShouldReturnErrorIfDaoGetReturnsError(t *testing.T) {
	const daoError = "dao error"

	getListDaoMock := &getListDaoMock{}
	getListDaoMock.On("Get", mock.Anything).Return([]domain.Tag{}, errors.New(daoError))

	err := NewGetList(nil, getListDaoMock).Run()

	assert.EqualError(t, err, daoError)
}

func TestGetListRunShouldReturnErrorIfPresenterShowReturnsError(t *testing.T) {
	const presenterError = "presenter error"

	getListDaoMock := &getListDaoMock{}
	getListDaoMock.On("Get", mock.Anything).Return([]domain.Tag{}, nil)

	getListPresenterMock := &getListPresenterMock{}
	getListPresenterMock.On("Show", mock.Anything).Return(errors.New(presenterError))

	err := NewGetList(getListPresenterMock, getListDaoMock).Run()

	assert.EqualError(t, err, presenterError)
}

func TestGetListRunShouldCallPresenterShowWithTagsFromDao(t *testing.T) {
	getListDaoMock := &getListDaoMock{}
	tag := domain.Tag{}
	tags := []domain.Tag{tag}
	getListDaoMock.On("Get", mock.Anything).Return(tags, nil)

	getListPresenterMock := &getListPresenterMock{}
	getListPresenterMock.On("Show", tags).Return(nil)

	NewGetList(getListPresenterMock, getListDaoMock).Run()

	getListPresenterMock.AssertExpectations(t)
}
