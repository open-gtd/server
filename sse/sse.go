package sse

type PushDataToSseFunc func(topic Topic, data interface{})

type Prefix string
type Topic string

type SseRegisterer interface {
	CreatePushDataFunc(prefix Prefix) PushDataToSseFunc
}

type Envelope struct {
	Topic Topic
	Data  interface{}
}
