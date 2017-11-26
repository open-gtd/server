package factories

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/business"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/open-gtd/server/tags/storage"
)

func TestUpdate(t *testing.T) {
	getDao = func() (storage.Dao, error) {
		return &daoMock{}, nil
	}

	rq := &requestMock{}
	rs := &responseMock{}

	c, cdf, _ := Update(rq, rs)

	var i business.UpdateController
	assert.Implements(t, &i, c)

	var ecdf api.ControllerDestroyFunc
	assert.IsType(t, ecdf, cdf)
}
