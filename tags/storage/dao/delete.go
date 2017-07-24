package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type delete struct {
}

func NewDelete() business.DeleteDao {
	return delete{}
}

func (d delete) Delete(name domain.Name) error {
	return nil
}
