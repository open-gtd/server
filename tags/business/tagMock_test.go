package business

import "github.com/open-gtd/server/tags/domain"

type tagMock struct {
}

func (tagMock) Rename(name domain.Name) {
	panic("implement me")
}

func (tagMock) ConvertToLabel() {
	panic("implement me")
}

func (tagMock) ConvertToArea() {
	panic("implement me")
}

func (tagMock) ConvertToContact() {
	panic("implement me")
}

func (tagMock) GetType() domain.TypeEnum {
	panic("implement me")
}

func (tagMock) GetOriginalName() domain.Name {
	panic("implement me")
}

func (tagMock) GetName() domain.Name {
	panic("implement me")
}
