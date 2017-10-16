package eBus

import "github.com/stretchr/testify/mock"

type busMock struct {
	mock.Mock
}

func (b *busMock) HasCallback(topic string) bool {
	args := b.Called(topic)
	return args.Bool(0)
}

func (b *busMock) WaitAsync() {
	b.Called()
}

func (b *busMock) Subscribe(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *busMock) SubscribeAsync(topic string, fn interface{}, transactional bool) error {
	args := b.Called(topic, fn, transactional)
	return args.Error(0)
}

func (b *busMock) SubscribeOnce(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *busMock) SubscribeOnceAsync(topic string, fn interface{}) error {
	args := b.Called(topic, fn)
	return args.Error(0)
}

func (b *busMock) Unsubscribe(topic string, handler interface{}) error {
	args := b.Called(topic, handler)
	return args.Error(0)
}

func (b *busMock) Publish(topic string, args ...interface{}) {
	b.Called(topic, args)
}
