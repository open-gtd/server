package strategies

import (
	"github.com/open-gtd/server/tags/domain"
)

func init() {
	RegisterStrategy(domain.Label, createLabel{})
}

type createLabel struct {
}

func (cl createLabel) Create(name domain.Name) (domain.Tag, error) {
	return domain.CreateLabel(name), nil
}
