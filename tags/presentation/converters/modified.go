package converters

import (
	"github.com/open-gtd/server/contract/tags"
	"github.com/open-gtd/server/tags/domain"
)

func ConvertAllToModified(t []domain.Tag) ([]tags.ModifiedTag, error) {
	result := make([]tags.ModifiedTag, len(t))

	for i, tag := range t {
		pTag, err := ConvertToModifiedTag(tag)
		if err != nil {
			return make([]tags.ModifiedTag, 0), err
		}

		result[i] = pTag
	}

	return result, nil
}

func ConvertToModifiedTag(t domain.Tag) (tags.ModifiedTag, error) {

	typeDescriptor, err := ConvertTypeToPresentation(t.GetType())
	if err != nil {
		return tags.ModifiedTag{}, err
	}

	return tags.ModifiedTag{
		OriginalName: string(t.GetOriginalName()),
		Name:         string(t.GetName()),
		Type:         typeDescriptor,
	}, nil
}
