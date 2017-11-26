package business

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
	tagErrors "github.com/open-gtd/server/tags/business/errors"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestGetRunShouldCallDaoGetWithName(t *testing.T) {
	name := domain.Name("tagName")

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", mock.Anything).Return(nil)

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", name).Return(&tagMock{}, nil)

	NewGet(getPresenterMock, getDaoMock).Run(name)

	getDaoMock.AssertExpectations(t)
}

func TestGetRunShouldReturnErrorIfDaoGetReturnsError(t *testing.T) {
	const daoError = "dao error"

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(&tagMock{}, errors.New(daoError))

	err := NewGet(nil, getDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, daoError)
}

func TestGetRunShouldCallPresenterShowNotFoundIfDaoGetReturnsNotFoundError(t *testing.T) {
	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(&tagMock{}, tagErrors.NewNotFound())

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("ShowNotFound").Return(nil)

	NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	getPresenterMock.AssertExpectations(t)
}

func TestGetRunShouldReturnErrorIfPresenterShowReturnsError(t *testing.T) {
	const presenterError = "presenter error"

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(&tagMock{}, nil)

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", mock.Anything).Return(errors.New(presenterError))

	err := NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, presenterError)
}

func TestGetRunShouldCallPresenterShowWithTagFromDao(t *testing.T) {
	getDaoMock := &getDaoMock{}
	tagMock := &tagMock{}
	getDaoMock.On("Get", mock.Anything).Return(tagMock, nil)

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", tagMock).Return(nil)

	NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	getPresenterMock.AssertExpectations(t)
}