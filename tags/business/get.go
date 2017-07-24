package business

import "github.com/open-gtd/server/tags/domain"

type GetPresenter interface {
	Show(t domain.Tag) error
}

type GetController interface {
	Run() error
}

type Get interface {
	Run() error
}

type get struct {
	presenter GetPresenter
}

func NewGet(p GetPresenter) Get {
	return get{presenter: p}
}

func (ga get) Run() error {
	return ga.presenter.Show(domain.CreateLabel("xx"))
}
