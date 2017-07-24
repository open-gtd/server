package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterStrategy(domain.Contact, createContact{})
}

type createContact struct {
}

func (cl createContact) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateContact(name), nil
}
