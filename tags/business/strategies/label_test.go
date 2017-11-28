package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestLabelStrategyCreate(t *testing.T) {
	sut := labelStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, domain.Label, tag.GetType())
	assert.Nil(t, err)
}

func TestLabelStrategyConvert(t *testing.T) {
	tag := domain.CreateContact(domain.Name("tagName"))

	sut := labelStrategy{}

	label, err := sut.Convert(tag)

	assert.Equal(t, domain.Label, label.GetType())
	assert.Nil(t, err)
}
