package converters

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/open-gtd/server/auth/business"
)

func TestConvertLoginDataToBusiness(t *testing.T) {
	const userName = "ua"
	const securityCode = "sc"

	b := ConvertLoginDataToBusiness(presentation.LoginData{UserName: userName, SecurityCode: securityCode})

	assert.Equal(t, business.LoginData{Name: userName, SecurityCode: securityCode}, b)
}
