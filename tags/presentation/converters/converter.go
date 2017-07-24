package converters

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/presentation"
)

func ConvertAllToPresentation(t []domain.Tag) ([]presentation.Tag, error) {
	result := make([]presentation.Tag, len(t))

	for i, tag := range t {
		pTag, err := ConvertToPresentation(tag)
		if err != nil {
			return make([]presentation.Tag, 0), err
		}

		result[i] = pTag
	}

	return result, nil
}

func ConvertToPresentation(t domain.Tag) (presentation.Tag, error) {

	typeDescriptor, err := ConvertTypeToPresentation(t.GetType())
	if err != nil {
		return presentation.Tag{}, err
	}

	return presentation.Tag{
		Name: presentation.Name(t.GetName()),
		Type: typeDescriptor,
	}, nil
}

func ConvertTypeToPresentation(t domain.TypeEnum) (presentation.TypeDescriptor, error) {
	switch t {
	case domain.Label:
		return presentation.Label, nil
	case domain.Area:
		return presentation.Area, nil
	case domain.Contact:
		return presentation.Contact, nil

	}
	return presentation.EmptyType, errors.New(fmt.Sprintf("Cannot convert TypeEnum '%v' to TypeDescriptor.", t))
}

func ConvertTypeToDomain(t presentation.TypeDescriptor) (domain.TypeEnum, error) {
	switch t {
	case presentation.Label:
		return domain.Label, nil
	case presentation.Area:
		return domain.Area, nil
	case presentation.Contact:
		return domain.Contact, nil

	}
	return domain.EmptyType, errors.New(fmt.Sprintf("Cannot convert TypeDescriptor '%v' to TypeEnum.", t))
}
