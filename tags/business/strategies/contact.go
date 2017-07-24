package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterCreateStrategy(domain.Contact, contactStrategy{})
	RegisterConvertStrategy(domain.Contact, contactStrategy{})
}

type contactStrategy struct {
}

func (cs contactStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateContact(name), nil
}

func (cs contactStrategy) Convert(tag domain.Tag) (domain.Tag, error) {
	tag.ConvertToContact()
	return tag, nil
}
