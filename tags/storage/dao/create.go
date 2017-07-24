package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type create struct {
}

func NewCreate() business.CreateDao {
	return create{}
}

func (c create) Insert(domain.Tag) error {
	return nil
}
