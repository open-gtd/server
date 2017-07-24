package business

import (
	"github.com/open-gtd/server/tags/business/strategies"
	"github.com/open-gtd/server/tags/domain"
)

type UpdatePresenter interface {
	Show(t domain.Tag) error
}

type UpdateDao interface {
	Get(name domain.Name) (domain.Tag, error)
	Save(name domain.Name, tag domain.Tag) error
}

type UpdateController interface {
	Run() error
}

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
}

func NewUpdate(p UpdatePresenter, d UpdateDao) Update {
	return update{presenter: p, dao: d}
}

func (u update) Run(ud UpdateData) error {
	tag, err := u.dao.Get(ud.OriginalName)
	if err != nil {
		return err
	}

	if ud.Name != nil {
		tag.Rename(*ud.Name)
	}

	if ud.Type != nil {
		conv, err := strategies.GetConvertStrategy(*ud.Type)
		if err != nil {
			return err
		}

		tag, err = conv.Convert(tag)
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
