package converters

import (
	"github.com/open-gtd/server/contract/tags"
	"github.com/open-gtd/server/tags/domain"
)

func ConvertAllToTag(t []domain.Tag) ([]tags.Tag, error) {
	result := make([]tags.Tag, len(t))

	for i, tag := range t {
		pTag, err := ConvertToTag(tag)
		if err != nil {
			return make([]tags.Tag, 0), err
		}

		result[i] = pTag
	}

	return result, nil
}

func ConvertToTag(t domain.Tag) (tags.Tag, error) {

	typeDescriptor, err := ConvertTypeToPresentation(t.GetType())
	if err != nil {
		return tags.Tag{}, err
	}

	return tags.Tag{
		Name: string(t.GetName()),
		Type: typeDescriptor,
	}, nil
}
