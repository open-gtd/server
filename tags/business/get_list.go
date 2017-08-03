package business

import "github.com/open-gtd/server/tags/domain"

type GetListPresenter interface {
	Show(t []domain.Tag) error
}

type GetListDao interface {
	Get() ([]domain.Tag, error)
}

type GetListController Controller

type GetListLogger interface{}

type GetList interface {
	Run() error
}

type getList struct {
	presenter GetListPresenter
	dao       GetListDao
	logger    GetListLogger
}

func NewGetList(p GetListPresenter, d GetListDao, l GetListLogger) GetList {
	return getList{presenter: p, dao: d, logger: l}
}

func (ga getList) Run() error {
	list, err := ga.dao.Get()
	if err != nil {
		return err
	}

	return ga.presenter.Show(list)
}
