package dao

import (
	"github.com/open-gtd/server/tags/business"
	"github.com/open-gtd/server/tags/domain"
)

type getList struct {
}

func NewGetList() business.GetListDao {
	return getList{}
}

func (gl getList) Get() ([]domain.Tag, error) {
	return make([]domain.Tag, 0), nil
}
