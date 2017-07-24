package strategies

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func init() {
	RegisterCreateStrategy(storage.Label, labelStrategy{})
}

type labelStrategy struct {
}

func (ls labelStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateLabel(name), nil
}
