package business

type DeletePresenter interface {
	ShowSucced() error
}

type DeleteController interface {
	Run() error
}

type Delete interface {
	Run() error
}

type delete struct {
	presenter DeletePresenter
}

func NewDelete(p DeletePresenter) Delete {
	return delete{presenter: p}
}

func (d delete) Run() error {
	return d.presenter.ShowSucced()
}
