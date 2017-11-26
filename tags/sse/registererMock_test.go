package sse

import (
	"github.com/stretchr/testify/mock"
	"github.com/open-gtd/server/sse"
)

type registererMock struct {
	mock.Mock
}

func (r *registererMock) CreatePushDataFunc(prefix sse.Prefix, closeNotify sse.ClientClosedNotificationFunc) sse.PushDataToSseFunc {
	args := r.Called(prefix, closeNotify)
	return args.Get(0).(sse.PushDataToSseFunc)
}
