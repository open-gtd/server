package dao

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/business/errors"
)

func GetTagByName(dao storage.Dao, name domain.Name) (domain.Tag, error) {
	tag, err := dao.Get(string(name))
	if err != nil {
		if err.Error() == storage.NotFoundError {
			return nil, errors.NewNotFound()
		}

		return nil, err
	}

	return ConvertToDomain(tag)
}