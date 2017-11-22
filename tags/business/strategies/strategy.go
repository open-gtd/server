package strategies

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/tags/domain"
)

var createStrategy = map[domain.TypeEnum]CreateStrategy{}
var convertStrategy = map[domain.TypeEnum]ConvertStrategy{}

func RegisterCreateStrategy(typeEnum domain.TypeEnum, strategy CreateStrategy) {
	createStrategy[typeEnum] = strategy
}

func RegisterConvertStrategy(typeEnum domain.TypeEnum, strategy ConvertStrategy) {
	convertStrategy[typeEnum] = strategy
}

type CreateStrategy interface {
	Create(name domain.Name) (domain.Tag, error)
}

type ConvertStrategy interface {
	Convert(tag domain.Tag) (domain.Tag, error)
}

func GetCreateStrategy(typeEnum domain.TypeEnum) (CreateStrategy, error) {
	s, ok := createStrategy[typeEnum]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("Unknown create strategy for create '%v'", typeEnum))
	}

	return s, nil
}

func GetConvertStrategy(typeEnum domain.TypeEnum) (ConvertStrategy, error) {
	s, ok := convertStrategy[typeEnum]
	if ok != true {
		return nil, errors.New(fmt.Sprintf("Unknown convert strategy for create '%v'", typeEnum))
	}

	return s, nil
}
