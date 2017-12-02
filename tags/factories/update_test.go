package factories

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/open-gtd/server/tags/storage"
	"errors"
)

func TestUpdate(t *testing.T) {
	getDao = func() (storage.Dao, error) {
		return &daoMock{}, nil
	}

	rq := &requestMock{}
	rs := &responseMock{}

	c, cd, _ := Update(rq, rs)

	var i business.UpdateController
	assert.Implements(t, &i, c)

	var destroyer connectionDestroyer
	assert.IsType(t, destroyer, cd)
}

func TestUpdate_ShouldReturnError_IfGetDaoReturnsError(t *testing.T) {
	const getDaoError = "get dao error"

	getDao = func() (storage.Dao, error) {
		return nil, errors.New(getDaoError)
	}

	rq := &requestMock{}
	rs := &responseMock{}

	_, _, err := Update(rq, rs)
	assert.EqualError(t, err, getDaoError)
}
