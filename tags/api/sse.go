package api

import "github.com/open-gtd/server/api"

var pushDataToSseFunc api.PushDataToSseFunc

func RegisterSse(sr api.PushDataToSseFunc) {
	pushDataToSseFunc = sr
}

func GetPushDataToSseFunc() api.PushDataToSseFunc {
	return pushDataToSseFunc
}
