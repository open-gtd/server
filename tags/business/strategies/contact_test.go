package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestContactStrategyCreate(t *testing.T) {
	sut := contactStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, domain.Contact, tag.GetType())
	assert.Nil(t, err)
}

func TestContactStrategyConvert(t *testing.T) {
	tag := domain.CreateContact(domain.Name("tagName"))

	sut := contactStrategy{}

	label, err := sut.Convert(tag)

	assert.Equal(t, domain.Contact, label.GetType())
	assert.Nil(t, err)
}
