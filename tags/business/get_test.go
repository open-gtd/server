package business

import (
	"errors"
	"testing"

	tagErrors "github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRunShouldCallDaoGetWithName(t *testing.T) {
	name := domain.Name("tagName")

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", mock.Anything).Return(nil)

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", name).Return(domain.Tag{}, nil)

	NewGet(getPresenterMock, getDaoMock).Run(name)

	getDaoMock.AssertExpectations(t)
}

func TestGetRunShouldReturnErrorIfDaoGetReturnsError(t *testing.T) {
	const daoError = "dao error"

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(domain.Tag{}, errors.New(daoError))

	err := NewGet(nil, getDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, daoError)
}

func TestGetRunShouldCallPresenterShowNotFoundIfDaoGetReturnsNotFoundError(t *testing.T) {
	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(domain.Tag{}, tagErrors.NewNotFound())

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("ShowNotFound").Return(nil)

	NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	getPresenterMock.AssertExpectations(t)
}

func TestGetRunShouldReturnErrorIfPresenterShowReturnsError(t *testing.T) {
	const presenterError = "presenter error"

	getDaoMock := &getDaoMock{}
	getDaoMock.On("Get", mock.Anything).Return(domain.Tag{}, nil)

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", mock.Anything).Return(errors.New(presenterError))

	err := NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, presenterError)
}

func TestGetRunShouldCallPresenterShowWithTagFromDao(t *testing.T) {
	getDaoMock := &getDaoMock{}
	tag := domain.Tag{}
	getDaoMock.On("Get", mock.Anything).Return(tag, nil)

	getPresenterMock := &getPresenterMock{}
	getPresenterMock.On("Show", tag).Return(nil)

	NewGet(getPresenterMock, getDaoMock).Run(domain.Name("tagName"))

	getPresenterMock.AssertExpectations(t)
}
