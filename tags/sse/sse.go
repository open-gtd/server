package sse

import (
	"github.com/open-gtd/server/sse"
	"github.com/open-gtd/server/tags/eventBus/topics"
	"github.com/open-gtd/server/tags/factories"
)

var sseRegisterer sse.SseRegisterer
var pushDataFunc sse.PushDataToSseFunc

func RegisterSse(sr sse.SseRegisterer) {
	sseRegisterer = sr
	pushDataFunc = sseRegisterer.CreatePushDataFunc("tags")

	factories.GetBus().Subscribe(topics.Created, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Created), v)
	})

	factories.GetBus().Subscribe(topics.Modified, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Modified), v)
	})

	factories.GetBus().Subscribe(topics.Deleted, func(v interface{}) {
		pushDataFunc(sse.Topic(topics.Deleted), v)
	})
}
