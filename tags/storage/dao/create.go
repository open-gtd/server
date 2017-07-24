package dao

import (
	"github.com/open-gtd/server/tags/business"
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

	return c.dao.Insert(t)
}
