package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestAreaStrategyCreate(t *testing.T) {
	tagMock := &tagMock{}

	domainCreateArea = func(name domain.Name) domain.Tag {
		return tagMock
	}

	sut := areaStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}

func TestAreaStrategyConvert(t *testing.T) {
	tagMock := &tagMock{}

	sut := areaStrategy{}

	tagMock.On("ConvertToArea")

	tag, err := sut.Convert(tagMock)

	tagMock.AssertExpectations(t)
	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}
