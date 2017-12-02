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

	c, cd, _ := Login(rq, rs)

	var i business.LoginController
	assert.Implements(t, &i, c)

	assert.NotNil(t, cd)

	var controllerDestroyer api.NullDestroyer
	assert.IsType(t, controllerDestroyer, cd)
}
