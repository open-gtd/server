package sse

import (
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/sse"
	"github.com/open-gtd/server/tags/eventBus/topics"
)

var registerer sse.Registerer
var pushDataFunc sse.PushDataToSseFunc

type registration struct {
	topic   eventBus.Topic
	handler func(v interface{})
}

func RegisterSse(sr sse.Registerer) {
	registerer = sr

	handlers := createRegistrations()

	pushDataFunc = registerer.CreatePushDataFunc("tags", func() {
		unregisterTopics(handlers)
	})

	registerTopics(handlers)
}

func createRegistrations() []registration {
	result := []registration{}

	result = append(result, registration{topics.Created, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Created), v)
	}})

	result = append(result, registration{topics.Modified, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Modified), v)
	}})

	result = append(result, registration{topics.Deleted, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Deleted), v)
	}})

	return result
}

func registerTopics(registrations []registration) {
	for _, r := range registrations {
		factories.GetBus().Subscribe(r.topic, r.handler)
	}
}

func unregisterTopics(registrations []registration) {
	for _, r := range registrations {
		factories.GetBus().Unsubscribe(r.topic, r.handler)
	}
}
