package business

import (
	"github.com/open-gtd/server/tags/domain"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/tags/business/strategies"

	"github.com/stretchr/testify/mock"
	"errors"
)

func TestRun_ShouldReturnError_IfGetCreateStrategyReturnsError(t *testing.T) {
	createData := CreateData{
		Type: domain.TypeEnum(255),
	}

	err := NewCreate(nil,nil,nil).Run(createData)

	assert.EqualError(t, err, "Unknown create strategy for create '255'")
}

func TestRun_ShouldReturnError_IfStrategyCreateReturnsError(t *testing.T) {

	strategyMock := &createStrategyMock{}
	strategies.RegisterCreateStrategy(domain.Label, strategyMock)

	const strategyError = "strategy create error"

	strategyMock.
		On("Create", mock.Anything).
		Return(
			tagMock{},
			errors.New(strategyError),
		)

	createData := CreateData{Type: domain.Label}

	err := NewCreate(nil,nil,nil).Run(createData)

	assert.EqualError(t, err, strategyError)
}
