package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestAreaStrategyCreate(t *testing.T) {
	sut := areaStrategy{}
	tag, err := sut.Create(domain.Name("tagName"))

	assert.Equal(t, domain.Area, tag.GetType())
	assert.Nil(t, err)
}

func TestAreaStrategyConvert(t *testing.T) {
	tag := domain.CreateLabel("Name")

	sut := areaStrategy{}
	tag, err := sut.Convert(tag)

	assert.Equal(t, domain.Area, tag.GetType())
	assert.Nil(t, err)
}
