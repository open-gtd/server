package echo

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type responseWriterMock struct {
	mock.Mock
}

func (rw *responseWriterMock) Header() http.Header {
	args := rw.Called()
	return args.Get(0).(http.Header)
}

func (rw *responseWriterMock) Write(data []byte) (int, error) {
	args := rw.Called(string(data[:]))
	return args.Int(0), args.Error(1)
}

func (rw *responseWriterMock) WriteHeader(i int) {
	rw.Called(i)
}

func (rw *responseWriterMock) CloseNotify() <-chan bool {
	args := rw.Called()
	return args.Get(0).(<-chan bool)
}

func (rw *responseWriterMock) Flush() {
	rw.Called()
}
