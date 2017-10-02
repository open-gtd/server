package presenters

import (
	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/mock"
)

type busMock struct {
	mock.Mock
}

func (b *busMock) Subscribe(topic eventBus.Topic, fn eventBus.BusHandler) error {
	args := b.Called(topic)
	return args.Error(0)
}
func (b *busMock) Unsubscribe(topic eventBus.Topic, handler eventBus.BusHandler) error {
	args := b.Called(topic)
	return args.Error(0)
}
func (b *busMock) Publish(topic eventBus.Topic, arg interface{}) {
	b.Called(topic, arg)
}
