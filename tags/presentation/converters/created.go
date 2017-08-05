package converters

import (
	"github.com/open-gtd/server/contract/tags"
	"github.com/open-gtd/server/tags/domain"
)

func ConvertAllToCreated(t []domain.Tag) ([]tags.CreatedTag, error) {
	result := make([]tags.CreatedTag, len(t))

	for i, tag := range t {
		pTag, err := ConvertToCreatedTag(tag)
		if err != nil {
			return make([]tags.CreatedTag, 0), err
		}

		result[i] = pTag
	}

	return result, nil
}

func ConvertToCreatedTag(t domain.Tag) (tags.CreatedTag, error) {

	typeDescriptor, err := ConvertTypeToPresentation(t.GetType())
	if err != nil {
		return tags.CreatedTag{}, err
	}

	return tags.CreatedTag{
		Name: string(t.GetName()),
		Type: typeDescriptor,
	}, nil
}
