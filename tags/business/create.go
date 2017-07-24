package business

import (
	"github.com/open-gtd/server/tags/business/strategies"
	"github.com/open-gtd/server/tags/domain"
)

type CreatePresenter interface {
	Show(t domain.Tag) error
}

type CreateDao interface {
	Insert(t domain.Tag) error
}

type CreateController interface {
	Run() error
}

type Create interface {
	Run(cd CreateData) error
}

type CreateData struct {
	Name domain.Name
	Type domain.TypeEnum
}

type create struct {
	presenter CreatePresenter
	dao       CreateDao
}

func NewCreate(p CreatePresenter, d CreateDao) Create {
	return create{presenter: p, dao: d}
}

func (c create) Run(cd CreateData) error {
	s, err := strategies.GetCreateStrategy(cd.Type)
	if err != nil {
		return err
	}

	tag, err := s.Create(cd.Name)
	if err != nil {
		return err
	}

	err = c.dao.Insert(tag)
	if err != nil {
		return err
	}

	return c.presenter.Show(tag)
}
