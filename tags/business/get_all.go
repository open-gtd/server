package business

import "github.com/open-gtd/server/tags/domain"

type GetAllPresenter interface {
	Show(t []domain.Tag) error
}

type GetAllController interface {
	Run() error
}

type GetAll interface {
	Run() error
}

type getAll struct {
	presenter GetAllPresenter
}

func NewGetAll(p GetAllPresenter) GetAll {
	return getAll{presenter: p}
}

func (ga getAll) Run() error {
	return ga.presenter.Show(make([]domain.Tag, 0))
}
