package business

import (
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/domain"
)

type DeletePresenter interface {
	ShowSucceed() error
	ShowNotFound() error
}

type DeleteDao interface {
	Delete(name domain.Name) error
}

type DeleteController Controller

type DeleteLogger interface {}

type Delete interface {
	Run(name domain.Name) error
}

type deleteTag struct {
	presenter DeletePresenter
	dao       DeleteDao
	logger    DeleteLogger
}

func NewDelete(p DeletePresenter, d DeleteDao, l DeleteLogger) Delete {
	return deleteTag{presenter: p, dao: d, logger: l}
}

func (d deleteTag) Run(name domain.Name) error {
	if err := d.dao.Delete(name); err != nil {

		if _, ok := err.(errors.NotFoundError); ok {
			return d.presenter.ShowNotFound()
		}

		return err
	}

	return d.presenter.ShowSucceed()
}
