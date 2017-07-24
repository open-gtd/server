package business

import "github.com/open-gtd/server/tags/domain"

type GetPresenter interface {
	Show(t domain.Tag) error
}

type GetDao interface {
	Get(name domain.Name) (domain.Tag, error)
}

type GetController interface {
	Run() error
}

type Get interface {
	Run(name domain.Name) error
}

type get struct {
	presenter GetPresenter
	dao GetDao
}

func NewGet(p GetPresenter, d GetDao) Get {
	return get{presenter: p, dao: d}
}

func (ga get) Run(name domain.Name) error {
	tag, err := ga.dao.Get(name)
	if err != nil {
		return err
	}

	return ga.presenter.Show(tag)
}
