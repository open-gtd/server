package business

import (
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
)

type GetPresenter interface {
	Show(t domain.Tag) error
	ShowNotFound() error
}

type GetDao interface {
	Get(name domain.Name) (domain.Tag, error)
}

type GetController Controller

type GetLogger interface {}

type Get interface {
	Run(name domain.Name) error
}

type get struct {
	presenter GetPresenter
	dao       GetDao
	logger    GetLogger
}

func NewGet(p GetPresenter, d GetDao, l GetLogger) Get {
	return get{presenter: p, dao: d, logger: l}
}

func (ga get) Run(name domain.Name) error {
	tag, err := ga.dao.Get(name)
	if err != nil {
		if _, ok := err.(errors.NotFoundError); ok {
			return ga.presenter.ShowNotFound()
		}

		return err
	}

	return ga.presenter.Show(tag)
}
