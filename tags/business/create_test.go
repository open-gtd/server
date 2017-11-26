package business

import (
	"github.com/open-gtd/server/tags/domain"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/tags/business/strategies"

	"github.com/stretchr/testify/mock"
	"errors"
	tagsErrors "github.com/open-gtd/server/tags/business/errors"
)

func TestCreateRunShouldReturnError_IfGetCreateStrategyReturnsError(t *testing.T) {
	createData := CreateData{
		Type: domain.TypeEnum(255),
	}

	err := NewCreate(nil, nil, nil).Run(createData)

	assert.EqualError(t, err, "Unknown create strategy for create '255'")
}

func TestCreateRunShouldReturnError_IfStrategyCreateReturnsError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	const strategyError = "strategy create error"

	strategyMock.
		On("Create", mock.Anything).
		Return(
		&tagMock{},
		errors.New(strategyError),
	)

	createData := CreateData{Type: domain.Label}

	err := NewCreate(nil, nil, nil).Run(createData)

	assert.EqualError(t, err, strategyError)
}

func TestCreateRunShouldPassTagCreatedByStrategyToDaoInsert(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	tagMock := &tagMock{}
	strategyMock.
		On("Create", mock.Anything).
		Return(tagMock, nil)

	const daoError = "dao error"
	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", tagMock).
		Return(errors.New(daoError))

	createData := CreateData{Type: domain.Label}
	NewCreate(nil, createDaoMock, nil).Run(createData)

	createDaoMock.AssertExpectations(t)
}

func TestCreateRunShouldReturnError_IfDaoInsertReturnsError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	strategyMock.
		On("Create", mock.Anything).
		Return(&tagMock{}, nil)

	const daoError = "dao error"
	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(errors.New(daoError))

	createData := CreateData{Type: domain.Label}
	err := NewCreate(nil, createDaoMock, nil).Run(createData)

	assert.EqualError(t, err, daoError)
}

func TestCreateRunShouldCallPresenterNotUnique_IfDaoInsertReturnsNotUniqueError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	strategyMock.
		On("Create", mock.Anything).
		Return(&tagMock{}, nil)

	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(tagsErrors.NewNotUnique())

	createPresenterMock := &createPresenterMock{}
	createPresenterMock.On("ShowNotUnique").Return(nil)

	createData := CreateData{Type: domain.Label}
	NewCreate(createPresenterMock, createDaoMock, nil).Run(createData)

	createPresenterMock.AssertExpectations(t)
}

func TestCreateRunShouldReturnError_IfPresenterShowNotUniqueReturnsError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	strategyMock.
		On("Create", mock.Anything).
		Return(&tagMock{}, nil)

	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(tagsErrors.NewNotUnique())

	const presenterError = "presenterError"
	createPresenterMock := &createPresenterMock{}
	createPresenterMock.On("ShowNotUnique").Return(errors.New(presenterError))

	createData := CreateData{Type: domain.Label}
	err := NewCreate(createPresenterMock, createDaoMock, nil).Run(createData)

	assert.EqualError(t, err, presenterError)
}

func TestCreateRunShouldReturnError_IfPresenterShowReturnsError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	strategyMock.
		On("Create", mock.Anything).
		Return(&tagMock{}, nil)

	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(nil)

	const presenterError = "presenterError"
	createPresenterMock := &createPresenterMock{}
	createPresenterMock.On("Show", mock.Anything).Return(errors.New(presenterError))

	createLoggerMock := &createLoggerMock{}
	createLoggerMock.On("TagCreated", mock.Anything)

	createData := CreateData{Type: domain.Label}
	err := NewCreate(createPresenterMock, createDaoMock, createLoggerMock).Run(createData)

	assert.EqualError(t, err, presenterError)
}

func TestCreateRunShouldPassCreatedTagToPresenter(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	tagMock := &tagMock{}
	strategyMock.
		On("Create", mock.Anything).
		Return(tagMock, nil)

	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(nil)

	const presenterError = "presenterError"
	createPresenterMock := &createPresenterMock{}
	createPresenterMock.On("Show", tagMock).Return(nil)

	createLoggerMock := &createLoggerMock{}
	createLoggerMock.On("TagCreated", mock.Anything)

	createData := CreateData{Type: domain.Label}
	NewCreate(createPresenterMock, createDaoMock, createLoggerMock).Run(createData)

	createPresenterMock.AssertExpectations(t)
}

func TestCreateRunShouldLogCreatedTag(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	tagMock := &tagMock{}
	strategyMock.
		On("Create", mock.Anything).
		Return(tagMock, nil)

	createDaoMock := &createDaoMock{}
	createDaoMock.
		On("Insert", mock.Anything).
		Return(nil)

	const presenterError = "presenterError"
	createPresenterMock := &createPresenterMock{}
	createPresenterMock.On("Show", mock.Anything).Return(nil)

	createLoggerMock := &createLoggerMock{}
	createLoggerMock.On("TagCreated", tagMock)

	createData := CreateData{Type: domain.Label}
	NewCreate(createPresenterMock, createDaoMock, createLoggerMock).Run(createData)

	createLoggerMock.AssertExpectations(t)
}
