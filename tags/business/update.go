package business

import "github.com/open-gtd/server/tags/domain"

type UpdatePresenter interface {
	Show(t domain.Tag) error
}

type UpdateController interface {
	Run() error
}

type Update interface {
	Run() error
}

type update struct {
	presenter UpdatePresenter
}

func NewUpdate(p UpdatePresenter) Update {
	return update{presenter: p}
}

func (u update) Run() error {
	label := domain.CreateLabel("xx")
	label.ConvertToArea()
	return u.presenter.Show(label)
}
