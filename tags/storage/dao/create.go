package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

type create struct {
	dao storage.Dao
}

func NewCreate(dao storage.Dao) business.CreateDao {
	return create{dao: dao}
}

func (c create) Insert(tag domain.Tag) error {
	t, err := ConvertToStorage(tag)
	if err != nil {
		return err
	}

	err = c.dao.Insert(t)
	if err.Error() == storage.NotUniqueError {
		return errors.NewNotUnique()
	}

	return err
}
