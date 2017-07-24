package strategies

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func init() {
	RegisterCreateStrategy(storage.Area, areaStrategy{})
}

type areaStrategy struct {
}

func (as areaStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateArea(name), nil
}
