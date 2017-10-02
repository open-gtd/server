package eBus

import "github.com/stretchr/testify/mock"

type eventBusMock struct {
	mock.Mock
}

func (b *eventBusMock) Subscribe(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *eventBusMock) SubscribeAsync(topic string, fn interface{}, transactional bool) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *eventBusMock) SubscribeOnce(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *eventBusMock) SubscribeOnceAsync(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *eventBusMock) Unsubscribe(topic string, handler interface{}) error {
	args := b.Called(topic, handler)
	return args.Error(0)
}

func (b *eventBusMock) HasCallback(topic string) bool {
	args := b.Called(topic)
	return args.Bool(0)
}

func (b *eventBusMock) WaitAsync() {
	b.Called()
}

func (b *eventBusMock) Publish(topic string, args ...interface{}) {
	b.Called(topic, args)
}
