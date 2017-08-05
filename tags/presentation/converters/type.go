package converters

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/contract/tags"
	"github.com/open-gtd/server/tags/domain"
)

func ConvertTypeToPresentation(t domain.TypeEnum) (string, error) {
	switch t {
	case domain.Label:
		return tags.Label, nil
	case domain.Area:
		return tags.Area, nil
	case domain.Contact:
		return tags.Contact, nil

	}
	return tags.EmptyType, errors.New(fmt.Sprintf("Cannot convert TypeEnum '%v' to TypeDescriptor.", t))
}

func ConvertTypeToDomain(t string) (domain.TypeEnum, error) {
	switch t {
	case tags.Label:
		return domain.Label, nil
	case tags.Area:
		return domain.Area, nil
	case tags.Contact:
		return domain.Contact, nil

	}
	return domain.EmptyType, errors.New(fmt.Sprintf("Cannot convert TypeDescriptor '%v' to TypeEnum.", t))
}
