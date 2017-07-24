package strategies

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func init() {
	RegisterCreateStrategy(storage.Contact, contactStrategy{})
}

type contactStrategy struct {
}

func (cs contactStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateContact(name), nil
}
