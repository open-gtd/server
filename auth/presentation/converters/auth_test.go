package converters

import (
	"testing"
	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/auth/presentation"
)

func TestConvertToPresentation(t *testing.T) {
	const cert = "asd"

	p := ConvertToPresentation(domain.Cert{Token: cert})

	assert.Equal(t, presentation.Cert{Token: cert}, p)
}
