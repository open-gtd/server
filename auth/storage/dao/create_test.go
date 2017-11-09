package dao

import (
	"testing"
	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	assert.Nil(t,
		NewLogin().
			Authorize(
				domain.Auth{
					SecurityCode:"xx",
					UserName:"yy",
				},
			),
	)
}
