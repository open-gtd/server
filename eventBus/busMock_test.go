package eventBus

import (
	"github.com/stretchr/testify/mock"
)

type busMock struct {
	mock.Mock
}

func (b *busMock) Subscribe(topic Topic, fn BusHandler) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *busMock) Unsubscribe(topic Topic, handler BusHandler) error {
	args := b.Called(topic, handler)
	return args.Error(0)
}

func (b *busMock) Publish(topic Topic, arg interface{}) {
	b.Called(topic, arg)
}
