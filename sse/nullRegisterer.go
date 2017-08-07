package sse

type NullRegisterer struct{}

func (NullRegisterer) CreatePushDataFunc(prefix Prefix, closeNotify ClientClosedNotificationFunc) PushDataToSseFunc {
	return func(topic Topic, data interface{}) {}
}
