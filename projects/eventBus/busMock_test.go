package eventBus

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/eventBus"
)

type busMock struct {
	mock.Mock
}

func (b *busMock) Subscribe(topic eventBus.Topic, fn eventBus.BusHandler) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *busMock) Unsubscribe(topic eventBus.Topic, handler eventBus.BusHandler) error {
	args := b.Called(topic, handler)
	return args.Error(0)
}

func (b *busMock) Publish(topic eventBus.Topic, arg interface{}) {
	b.Called(topic, arg)
}
