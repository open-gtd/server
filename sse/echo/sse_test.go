package echo

import (
	"testing"

	"net/http"

	"errors"

	"github.com/labstack/echo"
	"github.com/open-gtd/server/sse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePushDataFuncShouldRegisterGetRouteWithGivenPrefix(t *testing.T) {
	var route *echo.Route

	router := &routerMock{}
	router.On("Add", "GET", "/xxx", mock.Anything, mock.Anything).Return(route)

	sut := NewSseRegisterer(router)

	sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})
}

func TestCreatePushDataFuncShouldSetOkStatus(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", http.StatusOK)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go handlerFunc(context)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})

	//make sure go routine ends
	notifyClose <- true

	writerMock.AssertExpectations(t)
}

func TestCreatePushDataFuncShouldSetContentType(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	headers := make(http.Header)
	writerMock.On("Header").Return(headers)
	writerMock.On("WriteHeader", mock.Anything)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go handlerFunc(context)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})

	//make sure go routine ends
	notifyClose <- true

	assert.Equal(t, headers[echo.HeaderContentType], []string{echo.MIMEApplicationJSON})
}

func TestCreatePushDataFuncShouldWritePushedDataToOutput(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", mock.Anything)
	writerMock.On("Write", "{\"Topic\":\"topic\",\"Data\":11}\n").Return(1, nil)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)
	writerMock.On("Flush")

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	handled := make(chan bool)

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go func(handleNotify chan<- bool) {
				handlerFunc(context)
				handleNotify <- true
			}(handled)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	pushDataFunc := sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})

	pushDataFunc("topic", 11)
	//make sure go routine ends
	notifyClose <- true

	<-handled

	writerMock.AssertExpectations(t)
}

func TestCreatePushDataFuncShouldWriteAllPushedDataToOutput(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", mock.Anything)
	writerMock.On("Write", "{\"Topic\":\"topic\",\"Data\":11}\n").Return(1, nil)
	writerMock.On("Write", "{\"Topic\":\"topic\",\"Data\":12}\n").Return(1, nil)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)
	writerMock.On("Flush")

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	handled := make(chan bool)

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go func(handleNotify chan<- bool) {
				handlerFunc(context)
				handleNotify <- true
			}(handled)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	pushDataFunc := sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})

	pushDataFunc("topic", 11)
	pushDataFunc("topic", 12)

	//make sure go routine ends
	notifyClose <- true

	<-handled

	writerMock.AssertExpectations(t)
}

func TestCreatePushDataFuncShouldCallCloseNotifyOnConnectionClosed(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", mock.Anything)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go handlerFunc(context)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	closed := false
	sut.CreatePushDataFunc(sse.Prefix("xxx"), func() { closed = true })

	//make sure go routine ends
	notifyClose <- true

	assert.True(t, closed)
}

func TestCreatePushDataFuncShouldCallCloseNotifyOnWriterError(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", mock.Anything)

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go handlerFunc(context)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	closed := false
	sut.CreatePushDataFunc(sse.Prefix("xxx"), func() { closed = true })

	//make sure go routine ends
	notifyClose <- true

	assert.True(t, closed)
}

func TestCreatePushDataFuncShouldWriteCallCloseNotifyOnWriterError(t *testing.T) {
	var route *echo.Route
	writerMock := &responseWriterMock{}
	writerMock.On("Header").Return(make(http.Header))
	writerMock.On("WriteHeader", mock.Anything)
	writerMock.On("Write", mock.Anything).Return(0, errors.New("a"))

	notifyClose := make(chan bool)
	var readNotifyClose <-chan bool = notifyClose
	writerMock.On("CloseNotify").Return(readNotifyClose)
	writerMock.On("Flush")

	var context echo.Context = &contextMock{
		ResponseWriterMock: writerMock,
	}

	handled := make(chan bool)
	router := &routerMock{}
	router.
		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			handlerFunc := args.Get(2).(echo.HandlerFunc)
			go func(handleNotify chan<- bool) {
				handlerFunc(context)
				handleNotify <- true
			}(handled)
		}).
		Return(route)

	sut := NewSseRegisterer(router)

	closed := false
	pushDataFunc := sut.CreatePushDataFunc(sse.Prefix("xxx"), func() { closed = true })

	pushDataFunc("topic", 11)

	//make sure go routine ends
	notifyClose <- true

	<-handled

	assert.True(t, closed)
}

//func TestCreatePushDataFuncShouldWorkEvenIfThereIsNoReceiver(t *testing.T) {
//	var route *echo.Route
//	writerMock := &responseWriterMock{}
//	writerMock.On("Header").Return(make(http.Header))
//	writerMock.On("WriteHeader", mock.Anything)
//	writerMock.On("Write", mock.Anything).Return(0, errors.New("a"))
//
//	notifyClose := make(chan bool)
//	var readNotifyClose <-chan bool = notifyClose
//	writerMock.On("CloseNotify").Return(readNotifyClose)
//	writerMock.On("Flush")
//
//	router := &routerMock{}
//	router.
//		On("Add", "GET", "/xxx", mock.Anything, mock.Anything).
//		Return(route)
//
//	sut := NewSseRegisterer(router)
//
//	pushDataFunc := sut.CreatePushDataFunc(sse.Prefix("xxx"), func() {})
//
//	pushDataFunc("topic", 11)
//
//	//make sure go routine ends
//	notifyClose <- true
//}
