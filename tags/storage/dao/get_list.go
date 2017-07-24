package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

type getList struct {
	dao storage.Dao
}

func NewGetList(dao storage.Dao) business.GetListDao {
	return getList{dao: dao}
}

func (gl getList) Get() ([]domain.Tag, error) {

	tags, err := gl.dao.GetList()
	if err != nil {
		return nil, err
	}

	return ConvertAllToDomain(tags)
}
