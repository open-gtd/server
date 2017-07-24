package business

import "github.com/open-gtd/server/tags/domain"

type CreatePresenter interface {
	Show(t domain.Tag) error
}

type CreateController interface {
	Run() error
}

type Create interface {
	Run() error
}

type create struct {
	presenter CreatePresenter
}

func NewCreate(p CreatePresenter) Create {
	return create{presenter: p}
}

func (c create) Run() error {
	return c.presenter.Show(domain.CreateLabel("xx"))
}
