package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestLabelStrategyCreate(t *testing.T) {
	tagMock := &tagMock{}

	domainCreateLabel = func(name domain.Name) domain.Tag {
		return tagMock
	}

	sut := labelStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}

func TestLabelStrategyConvert(t *testing.T) {
	tagMock := &tagMock{}

	sut := labelStrategy{}

	tagMock.On("ConvertToLabel")

	tag, err := sut.Convert(tagMock)

	tagMock.AssertExpectations(t)
	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}
