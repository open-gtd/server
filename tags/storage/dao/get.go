package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type get struct {
}

func NewGet() business.GetDao {
	return get{}
}

func (g get) Get(name domain.Name) (domain.Tag, error) {
	return domain.CreateArea(name), nil
}
