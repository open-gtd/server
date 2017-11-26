package business

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"errors"
	"github.com/open-gtd/server/tags/domain"
	tagErrors "github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/business/strategies"
)

func TestUpdateRunShouldCallDaoGetWithUpdateDataOriginalName(t *testing.T) {
	const originalName = domain.Name("original name")

	updatePresenterMock := &updatePresenterMock{}

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", originalName).Return(&tagMock{}, errors.New(""))

	data := UpdateData{
		OriginalName: originalName,
	}
	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	updateDaoMock.AssertExpectations(t)
}

func TestUpdateRunShouldReturnErrorIfDaoGetReturnsError(t *testing.T) {
	const daoError = "dao error"

	updatePresenterMock := &updatePresenterMock{}

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(&tagMock{}, errors.New(daoError))

	data := UpdateData{}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, daoError)
}

func TestUpdateRunShouldCallPresenterShowNotFoundIfDaoGetReturnsNotFoundError(t *testing.T) {
	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("ShowNotFound").Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(&tagMock{}, tagErrors.NewNotFound())

	data := UpdateData{}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	updatePresenterMock.AssertExpectations(t)
}

func TestUpdateRunShouldReturnErrorIfPresenterShowNotFoundReturnsError(t *testing.T) {
	const presenterError = "presenter error"

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("ShowNotFound").Return(errors.New(presenterError))

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(&tagMock{}, tagErrors.NewNotFound())

	data := UpdateData{}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, presenterError)
}

func TestUpdateRunShouldCallRenameWithNewNameOnTagIfUpdateDataContainsName(t *testing.T) {
	newName := domain.Name("new tag name")

	tagMock := &tagMock{}
	tagMock.On("Rename", newName)

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		Name: &newName,
	}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	tagMock.AssertExpectations(t)
}

func TestUpdateRunShouldGetStrategyToConvertTagIntoNewTagType(t *testing.T) {
	newType := domain.Label
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(tagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		Type: &newType,
	}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.Equal(t, requestedStrategyType, newType)
}

func TestUpdateRunShouldCallConvertWithTagOnStrategy(t *testing.T) {
	newType := domain.Contact
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", tagMock).Return(tagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		Type: &newType,
	}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	convertStrategyMock.AssertExpectations(t)
}

func TestUpdateRunShouldReturnErrorIfGetConvertStrategyReturnsError(t *testing.T) {
	const getConvertStrategyError = "get convert strategy error"
	newType := domain.Contact
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return nil, errors.New(getConvertStrategyError)
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)

	data := UpdateData{
		Type: &newType,
	}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, getConvertStrategyError)
}

func TestUpdateRunShouldReturnErrorIfStrategyConvertReturnsError(t *testing.T) {
	const convertError = "convert error"

	newType := domain.Contact
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(tagMock, errors.New(convertError))

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		Type: &newType,
	}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, convertError)
}

func TestUpdateRunShouldSaveUpdatedTagUsingOriginalName(t *testing.T) {
	const originalName = domain.Name("original name")
	newType := domain.Contact
	updatedTagMock := &tagMock{}
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(updatedTagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", originalName, updatedTagMock).Return(nil)

	data := UpdateData{
		OriginalName: originalName,
		Type:         &newType,
	}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	updateDaoMock.AssertExpectations(t)
}

func TestUpdateRunShouldReturnErrorIfSaveUpdatedTagReturnsError(t *testing.T) {
	const saveError = "save error"

	newType := domain.Contact
	const originalName = domain.Name("original name")
	updatedTagMock := &tagMock{}
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(updatedTagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(errors.New(saveError))

	data := UpdateData{
		OriginalName: originalName,
		Type:         &newType,
	}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, saveError)
}

func TestUpdateRunShouldCallPresenterShowTag(t *testing.T) {
	newType := domain.Contact
	const originalName = domain.Name("original name")
	updatedTagMock := &tagMock{}
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(updatedTagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", updatedTagMock).Return(nil)

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		OriginalName: originalName,
		Type:         &newType,
	}

	NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	updatePresenterMock.AssertExpectations(t)
}

func TestUpdateRunShouldReturnErrorIfPresenterShowTagReturnsError(t *testing.T) {
	const presenterError = "presenterError"

	newType := domain.Contact
	const originalName = domain.Name("original name")
	updatedTagMock := &tagMock{}
	tagMock := &tagMock{}

	var requestedStrategyType domain.TypeEnum

	convertStrategyMock := &convertStrategyMock{}
	convertStrategyMock.On("Convert", mock.Anything).Return(updatedTagMock, nil)

	getConvertStrategy = func(typeEnum domain.TypeEnum) (strategies.ConvertStrategy, error) {
		requestedStrategyType = typeEnum
		return convertStrategyMock, nil
	}

	updatePresenterMock := &updatePresenterMock{}
	updatePresenterMock.On("Show", mock.Anything).Return(errors.New(presenterError))

	updateDaoMock := &updateDaoMock{}
	updateDaoMock.On("Get", mock.Anything).Return(tagMock, nil)
	updateDaoMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	data := UpdateData{
		OriginalName: originalName,
		Type:         &newType,
	}

	err := NewUpdate(updatePresenterMock, updateDaoMock).Run(data)

	assert.EqualError(t, err, presenterError)
}
