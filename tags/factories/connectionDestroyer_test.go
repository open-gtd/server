package factories

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestDestroy_ShouldCallPoolReturnObject(t *testing.T) {

	poolMock := &poolMock{}
	daoPool = poolMock

	dao := &daoMock{}
	poolMock.On("ReturnObject", dao).Return(nil)

	connectionDestroyer{conn: dao}.Destroy()
}

func TestDestroy_ShouldReturnError_IfPoolReturnObjectReturnsError(t *testing.T) {
	const poolError = "pool error"

	poolMock := &poolMock{}
	daoPool = poolMock

	dao := &daoMock{}
	poolMock.On("ReturnObject", mock.Anything).Return(errors.New(poolError))

	err := connectionDestroyer{conn: dao}.Destroy()

	assert.EqualError(t, err, poolError)
}
