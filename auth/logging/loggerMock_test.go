package logging

import "github.com/stretchr/testify/mock"

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Print(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Printf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Debug(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Debugf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Info(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Infof(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Warn(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Warnf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Error(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Errorf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Fatal(i ...interface{}) {
	l.Called(i)
}

func (l *loggerMock) Fatalf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}

func (l *loggerMock) Panic(i ...interface{}) {
	l.Called(i...)
}

func (l *loggerMock) Panicf(format string, args ...interface{}) {
	arg := []interface{}{format}
	for _, a := range args {
		arg = append(arg, a)
	}
	l.Called(arg...)
}
