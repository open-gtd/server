package business

import "github.com/open-gtd/server/tags/domain"

type DeletePresenter interface {
	ShowSucced() error
}

type DeleteDao interface {
	Delete(name domain.Name) error
}

type DeleteController interface {
	Run() error
}

type Delete interface {
	Run() error
}

type delete struct {
	presenter DeletePresenter
	dao DeleteDao
}

func NewDelete(p DeletePresenter, d DeleteDao) Delete {
	return delete{presenter: p, dao: d}
}

func (d delete) Run() error {
	return d.presenter.ShowSucced()
}
