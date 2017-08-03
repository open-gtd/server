package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

type update struct {
	dao storage.Dao
}

func NewUpdate(dao storage.Dao) business.UpdateDao {
	return update{dao: dao}
}

func (u update) Get(name domain.Name) (domain.Tag, error) {
	return GetTagByName(u.dao, name)
}

func (u update) Save(name domain.Name, tag domain.Tag) error {
	t, err := ConvertToStorage(tag)
	if err != nil {
		if err.Error() == storage.NotFoundError {
			return errors.NewNotFound()
		}

		return err
	}

	return u.dao.Update(string(name), t)
}
