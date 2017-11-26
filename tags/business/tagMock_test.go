package business

import (
	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/mock"
)

type tagMock struct {
	mock.Mock
}

func (t *tagMock) Rename(name domain.Name) {
	t.Called(name)
}

func (t *tagMock) ConvertToLabel() {
	t.Called()
}

func (t *tagMock) ConvertToArea() {
	t.Called()
}

func (t *tagMock) ConvertToContact() {
	t.Called()
}

func (t *tagMock) GetType() domain.TypeEnum {
	args := t.Called()
	return args.Get(0).(domain.TypeEnum)
}

func (t *tagMock) GetOriginalName() domain.Name {
	panic("implement me")
}

func (t *tagMock) GetName() domain.Name {
	panic("implement me")
}
