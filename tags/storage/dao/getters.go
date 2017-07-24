package dao

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

func GetTagByName(dao storage.Dao, name domain.Name) (domain.Tag, error) {
	tag, err := dao.Get(string(name))
	if err != nil {
		return nil, err
	}

	return ConvertToDomain(tag)
}
