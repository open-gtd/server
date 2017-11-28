package dao

import (
	"errors"
	"fmt"

	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/storage/dao/strategies"
)

func ConvertAllToDomain(tags []storage.Tag) ([]domain.Tag, error) {
	result := make([]domain.Tag, len(tags))
	for index, tag := range tags {
		dTag, err := ConvertToDomain(tag)
		if err != nil {
			return []domain.Tag{}, err
		}

		result[index] = dTag
	}

	return result, nil
}

func ConvertToDomain(tag storage.Tag) (domain.Tag, error) {
	s, err := strategies.GetCreateStrategy(tag.Type)
	if err != nil {
		return domain.Tag{}, err
	}

	return s.Create(domain.Name(tag.Name))
}

func ConvertToStorage(tag domain.Tag) (storage.Tag, error) {
	t, err := ConvertTypeToStorage(tag.GetType())
	if err != nil {
		return storage.Tag{}, err
	}

	return storage.Tag{
		Name: string(tag.GetName()),
		Type: t,
	}, nil
}

func ConvertTypeToStorage(t domain.TypeEnum) (uint8, error) {
	switch t {
	case domain.Label:
		return storage.Label, nil
	case domain.Area:
		return storage.Area, nil
	case domain.Contact:
		return storage.Contact, nil
	}

	return 0, errors.New(fmt.Sprintf("Unknown tag type for storage value '%v'", t))
}
