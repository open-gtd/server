package api

type PushDataToSseFunc func(v interface{})

type SseRegisterer interface {
	CreatePushDataFunc() PushDataToSseFunc
}
