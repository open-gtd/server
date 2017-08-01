package business

import (
	"github.com/open-gtd/server/tags/business/strategies"
	"github.com/open-gtd/server/tags/domain"
)

type CreatePresenter interface {
	Show(t domain.Tag) error
}

type CreateDao interface {
	Insert(t domain.Tag) error
}

type CreateController Controller

type CreateLogger interface{
	TagCreated(tag domain.Tag)
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
	dao       CreateDao
	logger    CreateLogger
}

func NewCreate(p CreatePresenter, d CreateDao, l CreateLogger) Create {
	return create{presenter: p, dao: d, logger: l}
}

func (c create) Run(cd CreateData) error {

	s, err := strategies.GetCreateStrategy(cd.Type)
	if err != nil {
		return err
	}

	tag, err := s.Create(cd.Name)
	if err != nil {
		return err
	}

	err = c.dao.Insert(tag)
	if err != nil {
		return err
	}

	c.logger.TagCreated(tag)

	return c.presenter.Show(tag)
}
