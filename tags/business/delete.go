package business

import "github.com/open-gtd/server/tags/domain"

type DeletePresenter interface {
	ShowSucced() error
	ShowNotFound() error
}

type DeleteDao interface {
	Delete(name domain.Name) error
}

type DeleteController interface {
	Run() error
}

type Delete interface {
	Run(name domain.Name) error
}

type delete struct {
	presenter DeletePresenter
	dao       DeleteDao
}

func NewDelete(p DeletePresenter, d DeleteDao) Delete {
	return delete{presenter: p, dao: d}
}

func (d delete) Run(name domain.Name) error {
	if err := d.dao.Delete(name); err != nil {

		if err.Error() == notFoundError {
			return d.presenter.ShowNotFound()
		}

		return err
	}

	return d.presenter.ShowSucced()
}
