package sse

type PushDataToSseFunc func(topic Topic, data interface{})
type ClientClosedNotificationFunc func()

type Prefix string
type Topic string

type Registerer interface {
	CreatePushDataFunc(prefix Prefix, closeNotify ClientClosedNotificationFunc) PushDataToSseFunc
}

type Envelope struct {
	Topic Topic
	Data  interface{}
}
