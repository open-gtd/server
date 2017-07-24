package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterStrategy(domain.Area, createArea{})
}

type createArea struct {
}

func (cl createArea) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateArea(name), nil
}
