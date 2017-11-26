package business

import (
	"github.com/open-gtd/server/tags/business/errors"
	"github.com/open-gtd/server/tags/business/strategies"
	"github.com/open-gtd/server/tags/domain"
)

type UpdatePresenter interface {
	Show(t domain.Tag) error
	ShowNotFound() error
}

type UpdateDao interface {
	Get(name domain.Name) (domain.Tag, error)
	Save(name domain.Name, tag domain.Tag) error
}

type UpdateController Controller

type UpdateLogger interface{}

type Update interface {
	Run(ud UpdateData) error
}

type UpdateData struct {
	OriginalName domain.Name
	Name         *domain.Name
	Type         *domain.TypeEnum
}

type update struct {
	presenter UpdatePresenter
	dao       UpdateDao
	logger    UpdateLogger
}

func NewUpdate(p UpdatePresenter, d UpdateDao) Update {
	return update{presenter: p, dao: d}
}

var getConvertStrategy = strategies.GetConvertStrategy

func (u update) Run(ud UpdateData) error {
	tag, err := u.dao.Get(ud.OriginalName)
	if err != nil {
		if _, ok := err.(errors.NotFoundError); ok {
			return u.presenter.ShowNotFound()
		}

		return err
	}

	if ud.Name != nil {
		tag.Rename(*ud.Name)
	}

	if ud.Type != nil {
		convertStrategy, err := getConvertStrategy(*ud.Type)
		if err != nil {
			return err
		}

		tag, err = convertStrategy.Convert(tag)
		if err != nil {
			return err
		}
	}

	err = u.dao.Save(ud.OriginalName, tag)
	if err != nil {
		return err
	}

	return u.presenter.Show(tag)
}
