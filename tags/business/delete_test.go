package business

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/tags/domain"
	tagErrors "github.com/open-gtd/server/tags/business/errors"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestDeleteRunShouldCallDaoGetWithName(t *testing.T) {
	name := domain.Name("tagName")

	deletePresenterMock := &deletePresenterMock{}
	deletePresenterMock.On("ShowSucceed", mock.Anything).Return(nil)

	deleteDaoMock := &deleteDaoMock{}
	deleteDaoMock.On("Delete", name).Return(nil)

	NewDelete(deletePresenterMock, deleteDaoMock).Run(name)

	deleteDaoMock.AssertExpectations(t)
}

func TestDeleteRunShouldReturnErrorIfDaoGetReturnsError(t *testing.T) {
	const daoError = "dao error"

	deleteDaoMock := &deleteDaoMock{}
	deleteDaoMock.On("Delete", mock.Anything).Return(errors.New(daoError))

	err := NewDelete(nil, deleteDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, daoError)
}

func TestDeleteRunShouldCallPresenterShowNotFoundIfDaoGetReturnsNotFoundError(t *testing.T) {
	deleteDaoMock := &deleteDaoMock{}
	deleteDaoMock.On("Delete", mock.Anything).Return(tagErrors.NewNotFound())

	deletePresenterMock := &deletePresenterMock{}
	deletePresenterMock.On("ShowNotFound").Return(nil)

	NewDelete(deletePresenterMock, deleteDaoMock).Run(domain.Name("tagName"))

	deletePresenterMock.AssertExpectations(t)
}

func TestDeleteRunShouldCallPresenterShowSucceed(t *testing.T) {
	name := domain.Name("tagName")

	deleteDaoMock := &deleteDaoMock{}
	deleteDaoMock.On("Delete", mock.Anything).Return(nil)

	deletePresenterMock := &deletePresenterMock{}
	deletePresenterMock.On("ShowSucceed", name).Return(nil)

	NewDelete(deletePresenterMock, deleteDaoMock).Run(name)

	deletePresenterMock.AssertExpectations(t)
}

func TestDeleteRunShouldReturnErrorIfPresenterShowSucceedReturnsError(t *testing.T) {
	const presenterError = "presenter error"

	deleteDaoMock := &deleteDaoMock{}
	deleteDaoMock.On("Delete", mock.Anything).Return(nil)

	deletePresenterMock := &deletePresenterMock{}
	deletePresenterMock.On("ShowSucceed", mock.Anything).Return(errors.New(presenterError))

	err := NewDelete(deletePresenterMock, deleteDaoMock).Run(domain.Name("tagName"))

	assert.EqualError(t, err, presenterError)
}
