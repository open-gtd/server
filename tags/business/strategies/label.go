package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterCreateStrategy(domain.Label, labelStrategy{})
	RegisterConvertStrategy(domain.Label, labelStrategy{})
}

type labelStrategy struct {
}

var domainCreateLabel = domain.CreateLabel

func (ls labelStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domainCreateLabel(name), nil
}

func (ls labelStrategy) Convert(tag domain.Tag) (domain.Tag, error) {
	tag.ConvertToLabel()
	return tag, nil
}
