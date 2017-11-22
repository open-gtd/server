package strategies

import (
	"testing"

	"github.com/open-gtd/server/tags/domain"
	"github.com/stretchr/testify/assert"
)

func TestRegisterCreateStrategy(t *testing.T) {
	checkRegisterCreateStrategy(t, domain.Label, areaStrategy{})
	checkRegisterCreateStrategy(t, domain.Label, labelStrategy{})
	checkRegisterCreateStrategy(t, domain.Area, areaStrategy{})
	checkRegisterCreateStrategy(t, domain.Area, labelStrategy{})
	checkRegisterCreateStrategy(t, domain.Contact, areaStrategy{})
	checkRegisterCreateStrategy(t, domain.Contact, labelStrategy{})
}
func checkRegisterCreateStrategy(t *testing.T, enum domain.TypeEnum, strategy CreateStrategy) {
	RegisterCreateStrategy(enum, strategy)
	assert.Equal(t, strategy, createStrategy[enum])
}

func TestRegisterConvertStrategy(t *testing.T) {
	checkRegisterConvertStrategy(t, domain.Label, areaStrategy{})
	checkRegisterConvertStrategy(t, domain.Label, labelStrategy{})
	checkRegisterConvertStrategy(t, domain.Area, areaStrategy{})
	checkRegisterConvertStrategy(t, domain.Area, labelStrategy{})
	checkRegisterConvertStrategy(t, domain.Contact, areaStrategy{})
	checkRegisterConvertStrategy(t, domain.Contact, labelStrategy{})
}
func checkRegisterConvertStrategy(t *testing.T, enum domain.TypeEnum, strategy ConvertStrategy) {
	RegisterConvertStrategy(enum, strategy)
	assert.Equal(t, strategy, convertStrategy[enum])
}

func TestGetCreateStrategy(t *testing.T) {
	strategy := areaStrategy{}
	createStrategy[domain.Label] = strategy
	s, err := GetCreateStrategy(domain.Label)

	assert.Equal(t, strategy, s)
	assert.Nil(t, err)
}

func TestGetCreateStrategyError(t *testing.T) {
	createStrategy = map[domain.TypeEnum]CreateStrategy{}
	_, err := GetCreateStrategy(domain.Label)

	assert.Error(t, err, "Unknown create strategy for create 'domain.Label'")
}

func TestGetConvertStrategy(t *testing.T) {
	strategy := areaStrategy{}
	convertStrategy[domain.Label] = strategy
	s, err := GetConvertStrategy(domain.Label)

	assert.Equal(t, strategy, s)
	assert.Nil(t, err)
}

func TestGetConvertStrategyError(t *testing.T) {
	convertStrategy = map[domain.TypeEnum]ConvertStrategy{}
	_, err := GetConvertStrategy(domain.Label)

	assert.Error(t, err, "Unknown create strategy for create 'domain.Label'")
}
