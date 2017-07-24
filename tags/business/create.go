package business

import (
	"github.com/open-gtd/server/tags/business/strategies"
	"github.com/open-gtd/server/tags/domain"
)

type CreatePresenter interface {
	Show(t domain.Tag) error
}

type CreateController interface {
	Run() error
}

type Create interface {
	Run(cd CreateData) error
}

type CreateData struct {
	Name domain.Name
	Type domain.TypeEnum
}

type create struct {
	presenter CreatePresenter
}

func NewCreate(p CreatePresenter) Create {
	return create{presenter: p}
}

func (c create) Run(cd CreateData) error {
	s, err := strategies.GetStrategy(cd.Type)
	if err != nil {
		return err
	}

	tag, err := s.Create(cd.Name)
	if err != nil {
		return err
	}

	return c.presenter.Show(tag)
}
