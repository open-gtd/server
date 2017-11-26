package business

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/tags/domain"
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
	tagMock := &tagMock{}
	tags := []domain.Tag{tagMock}
	getListDaoMock.On("Get", mock.Anything).Return(tags, nil)

	getListPresenterMock := &getListPresenterMock{}
	getListPresenterMock.On("Show", tags).Return(nil)

	NewGetList(getListPresenterMock, getListDaoMock).Run()

	getListPresenterMock.AssertExpectations(t)
}
