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

func (ls labelStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateLabel(name), nil
}

func (ls labelStrategy) Convert(tag domain.Tag) (domain.Tag, error) {
	tag.ConvertToLabel()
	return tag, nil
}
