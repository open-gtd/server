package factories

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/open-gtd/server/tags/storage"
	"errors"
)

func TestCreate(t *testing.T) {
	getDao = func() (storage.Dao, error) {
		return &daoMock{}, nil
	}

	rq := &requestMock{}
	rs := &responseMock{}

	c, cd, _ := Create(rq, rs)

	var i business.CreateController
	assert.Implements(t, &i, c)

	var destroyer connectionDestroyer
	assert.IsType(t, destroyer, cd)
}

func TestCreate_ShouldReturnError_IfGetDaoReturnsError(t *testing.T) {
	const getDaoError = "get dao error"

	getDao = func() (storage.Dao, error) {
		return nil, errors.New(getDaoError)
	}

	rq := &requestMock{}
	rs := &responseMock{}

	_, _, err := Create(rq, rs)
	assert.EqualError(t, err, getDaoError)
}
