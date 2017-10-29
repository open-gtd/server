package factories

import (
	"testing"

	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/business"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	rq := &requestMock{}
	rs := &responseMock{}

	c, cdf, _ := Login(rq, rs)

	var i business.LoginController
	assert.Implements(t, &i, c)

	var ecdf api.ControllerDestroyFunc
	assert.IsType(t, ecdf, cdf)
}
