package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
	"github.com/open-gtd/server/tags/storage"
)

type get struct {
	dao storage.Dao
}

func NewGet(dao storage.Dao) business.GetDao {
	return get{dao: dao}
}

func (g get) Get(name domain.Name) (domain.Tag, error) {
	return GetTagByName(g.dao, name)
}
