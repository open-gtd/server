package eBus

import (
	"testing"

	"errors"

	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBus_Subscribe_ShouldCallBusSubscribeWithTopicAndHandler(t *testing.T) {
	busMock := &eventBusMock{}

	sut := &bus{eBus: busMock}

	busMock.On("Subscribe", "topic", mock.AnythingOfType("eventBus.BusHandler")).Return(nil)

	sut.Subscribe(eventBus.Topic("topic"), func(arg interface{}) {})

	busMock.AssertExpectations(t)
}

func TestBus_Subscribe_ShouldReturnError_IfBusSubscribeReturnsError(t *testing.T) {
	busMock := &eventBusMock{}

	sut := &bus{eBus: busMock}

	const subscribeError = "subscribeError"
	busMock.On("Subscribe", mock.Anything, mock.Anything).Return(errors.New(subscribeError))

	err := sut.Subscribe(eventBus.Topic("topic"), func(arg interface{}) {})

	assert.EqualError(t, err, subscribeError)
}

func TestBus_Publish_ShouldCallBusPublish(t *testing.T) {
	busMock := &eventBusMock{}

	sut := &bus{eBus: busMock}

	busMock.On("Publish", "topic", []interface{}{"arg"}).Return(nil)

	sut.Publish(eventBus.Topic("topic"), "arg")

	busMock.AssertExpectations(t)
}

func TestBus_Unsubscribe_ShouldCallBusUnsubscribeWithTopicAndHandler(t *testing.T) {
	busMock := &eventBusMock{}

	sut := &bus{eBus: busMock}

	busMock.On("Unsubscribe", "topic", mock.AnythingOfType("eventBus.BusHandler")).Return(nil)

	sut.Unsubscribe(eventBus.Topic("topic"), func(arg interface{}) {})

	busMock.AssertExpectations(t)
}

func TestBus_Unsubscribe_ShouldReturnError_IfBusUnsubscribeReturnsError(t *testing.T) {
	busMock := &eventBusMock{}

	sut := &bus{eBus: busMock}

	const unsubscribeError = "unsubscribeError"
	busMock.On("Subscribe", mock.Anything, mock.Anything).Return(errors.New(unsubscribeError))

	err := sut.Subscribe(eventBus.Topic("topic"), func(arg interface{}) {})

	assert.EqualError(t, err, unsubscribeError)
}

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
