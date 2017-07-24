package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type update struct {
}

func NewUpdate() business.UpdateDao {
	return update{}
}

func (u update) Get(name domain.Name) (domain.Tag, error) {
	return domain.CreateArea(name), nil
}

func (u update) Save(tag domain.Tag) error {
	return nil
}
