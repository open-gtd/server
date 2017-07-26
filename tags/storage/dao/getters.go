package dao

import (
	"errors"

	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func GetTagByName(dao storage.Dao, name domain.Name) (domain.Tag, error) {
	tag, err := dao.Get(string(name))
	if err != nil {
		if err.Error() == storage.NotFoundError {
			return nil, errors.New(business.NotFoundError)
		}

		return nil, err
	}

	return ConvertToDomain(tag)
}
