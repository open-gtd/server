package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterCreateStrategy(domain.Area, areaStrategy{})
	RegisterConvertStrategy(domain.Area, areaStrategy{})
}

type areaStrategy struct {
}

func (as areaStrategy) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateArea(name), nil
}

func (as areaStrategy) Convert(tag domain.Tag) (domain.Tag, error) {
	tag.ConvertToArea()
	return tag, nil
}
