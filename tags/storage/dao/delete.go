package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/business/errors"
)

type delete struct {
	dao storage.Dao
}

func NewDelete(dao storage.Dao) business.DeleteDao {
	return delete{dao: dao}
}

func (d delete) Delete(name domain.Name) error {
	err := d.dao.Delete(string(name))
	if err.Error() == storage.NotFoundError {
		return errors.NewNotFound()
	}

	return err
}
