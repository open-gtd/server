package logging

import (
	"testing"

	"github.com/open-gtd/server/auth/domain"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Logged(t *testing.T) {
	l := &testLogger{}

	auth := domain.Auth{
		UserName:     "userName",
		SecurityCode: "sec",
	}

	l.On("Infof", "User '%v' has been logged.", domain.Name("userName"))
	l.On("Debugf", "Auth:", auth)

	logger = l
	sut := NewLogin()
	sut.Logged(auth)

	l.AssertExpectations(t)
}

type testLogger struct {
	mock.Mock
}

func (l *testLogger) Print(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Printf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Debug(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Debugf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Info(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Infof(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Warn(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Warnf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Error(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Errorf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Fatal(i ...interface{}) {
	l.Called(i)
}
func (l *testLogger) Fatalf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
func (l *testLogger) Panic(i ...interface{}) {
	l.Called(i...)
}
func (l *testLogger) Panicf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
