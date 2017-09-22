package converters

import (
	"github.com/open-gtd/server/auth/domain"
	"github.com/open-gtd/server/auth/presentation"
)

func ConvertToPresentation(cert domain.Cert) presentation.Cert {

	return presentation.Cert{
		Token: string(cert.Token),
	}
}
