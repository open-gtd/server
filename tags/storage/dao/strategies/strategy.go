package strategies

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/tags/domain"
)

var createStrategy map[uint8]CreateStrategy

func RegisterCreateStrategy(t uint8, strategy CreateStrategy) {
	if createStrategy == nil {
		createStrategy = make(map[uint8]CreateStrategy)
	}

	createStrategy[t] = strategy
}

type CreateStrategy interface {
	Create(name domain.Name) (domain.Tag, error)
}

func GetCreateStrategy(t uint8) (CreateStrategy, error) {
	s, ok := createStrategy[t]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("Unknown create strategy for create '%v'", t))
	}

	return s, nil
}
