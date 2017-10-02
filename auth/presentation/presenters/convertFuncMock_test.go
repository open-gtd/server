package presenters

import (
	"fmt"

	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
	"github.com/stretchr/testify/mock"
)

type convertFuncMock struct {
	mock.Mock
}

func (c *convertFuncMock) Convert(cert domain.Cert) presentation.Cert {
	args := c.Called(cert)

	result := args.Get(0)
	rCert, ok := result.(presentation.Cert)
	if !ok {
		panic(fmt.Sprintf("Return failed because object wasn't coddect type: {%v}", result))
	}
	return rCert
}
