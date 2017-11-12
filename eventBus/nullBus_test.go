package eventBus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubscribe(t *testing.T) {
	assert.Nil(t, NullBus{}.Subscribe(Topic("topic"), func(arg interface{}) {}))
}

func TestUnsubscribe(t *testing.T) {
	assert.Nil(t, NullBus{}.Unsubscribe(Topic("topic"), func(arg interface{}) {}))
}

func TestPublish(t *testing.T) {
	NullBus{}.Publish(Topic("t"), "arg")
}
