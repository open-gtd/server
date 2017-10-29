package sse

import (
	"github.com/stretchr/testify/mock"
)

type registererMock struct {
	mock.Mock
}

func (r *registererMock) CreatePushDataFunc(prefix Prefix, closeNotify ClientClosedNotificationFunc) PushDataToSseFunc {
	args := r.Called(prefix, closeNotify)
	return args.Get(0).(PushDataToSseFunc)
}
