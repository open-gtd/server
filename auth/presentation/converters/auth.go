package converters

import (
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
)

func ConvertToPresentation(cert domain.Cert) (presentation.Cert, error) {

	return presentation.Cert{
		Token: presentation.Token(cert.Token),
	}, nil
}
