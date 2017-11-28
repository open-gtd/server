package dao

import (
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func GetTagByName(dao storage.Dao, name domain.Name) (domain.Tag, error) {
	tag, err := dao.Get(string(name))
	if err != nil {
		if err.Error() == storage.NotFoundError {
			return domain.Tag{}, errors.NewNotFound()
		}

		return domain.Tag{}, err
	}

	return ConvertToDomain(tag)
}
