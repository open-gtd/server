package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestContactStrategyCreate(t *testing.T) {
	tagMock := &tagMock{}

	domainCreateContact = func(name domain.Name) domain.Tag {
		return tagMock
	}

	sut := contactStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}

func TestContactStrategyConvert(t *testing.T) {
	tagMock := &tagMock{}

	sut := contactStrategy{}

	tagMock.On("ConvertToContact")

	tag, err := sut.Convert(tagMock)

	tagMock.AssertExpectations(t)
	assert.Equal(t, tagMock, tag)
	assert.Nil(t, err)
}
