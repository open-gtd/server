package strategies

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/tags/domain"
)

var createStrategy map[domain.TypeEnum]CreateStrategy

func RegisterStrategy(typeEnum domain.TypeEnum, strategy CreateStrategy) {
	if createStrategy == nil {
		createStrategy = make(map[domain.TypeEnum]CreateStrategy)
	}

	createStrategy[typeEnum] = strategy
}

type CreateStrategy interface {
	Create(name domain.Name) (domain.Tag, error)
}

func GetStrategy(typeEnum domain.TypeEnum) (CreateStrategy, error) {
	s, ok := createStrategy[typeEnum]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("Unknown strategy for create '%v'", typeEnum))
	}

	return s, nil
}
